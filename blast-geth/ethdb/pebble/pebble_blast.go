// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package pebble implements the key-value database layer based on pebble.
package pebble

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

const (
	// numLevels is the level number of pebble sst files
	numLevels = 7
)

func (d *Database) onCompactionBeginBlast(info pebble.CompactionInfo) {
	if d.activeComp == 0 {
		d.compStartTime = time.Now()
	}
	l0 := info.Input[0]
	if l0.Level == 0 {
		d.level0Comp.Add(1)
	} else {
		d.nonLevel0Comp.Add(1)
	}
	d.activeComp++
}

var (
	compTotalCount atomic.Uint64
	compTotalTime  atomic.Value
)

func init() {
	compTotalTime.Store(time.Duration(0))
}

func AverageCompactionTime() (time.Duration, uint64) {
	soFar := compTotalCount.Load()
	if soFar == 0 {
		return 0, 0
	}
	dur := time.Duration(soFar)
	return compTotalTime.Load().(time.Duration) / dur, soFar
}

func (d *Database) onCompactionEndBlast(info pebble.CompactionInfo) {
	if d.activeComp == 1 {
		d.compTime.Add(int64(time.Since(d.compStartTime)))
	} else if d.activeComp == 0 {
		panic("should not happen")
	}
	d.activeComp--
	compTotalCount.Add(1)
	updated := compTotalTime.Load().(time.Duration) + info.Duration
	compTotalTime.Store(updated)
}

func NewBlastPebble(file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, formatVersion pebble.FormatMajorVersion) (*Database, error) {
	// Ensure we have some minimal caching and file guarantees
	if cache < minCache {
		cache = minCache
	}
	if handles < minHandles {
		handles = minHandles
	}
	logger := log.New("database", file)
	logger.Info("Allocated cache and file handles", "cache", common.StorageSize(cache*1024*1024), "handles", handles)

	// The max memtable size is limited by the uint32 offsets stored in
	// internal/arenaskl.node, DeferredBatchOp, and flushableBatchEntry.
	//
	// - MaxUint32 on 64-bit platforms;
	// - MaxInt on 32-bit platforms.
	//
	// It is used when slices are limited to Uint32 on 64-bit platforms (the
	// length limit for slices is naturally MaxInt on 32-bit platforms).
	//
	// Taken from https://github.com/cockroachdb/pebble/blob/master/internal/constants/constants.go
	maxMemTableSize := (1<<31)<<(^uint(0)>>63) - 1

	// Two memory tables is configured which is identical to leveldb,
	// including a frozen memory table and another live one.
	memTableLimit := 2
	memTableSize := cache * 1024 * 1024 / 2 / memTableLimit
	if memTableSize > maxMemTableSize {
		memTableSize = maxMemTableSize
	}
	db := &Database{
		fn:           file,
		log:          logger,
		quitChan:     make(chan chan error),
		writeOptions: &pebble.WriteOptions{Sync: !ephemeral},
	}
	opt := &pebble.Options{
		// Pebble has a single combined cache area and the write
		// buffers are taken from this too. Assign all available
		// memory allowance for cache.
		Cache:        pebble.NewCache(int64(cache * 1024 * 1024)),
		MaxOpenFiles: handles,

		// The size of memory table(as well as the write buffer).
		// Note, there may have more than two memory tables in the system.
		MemTableSize: uint64(memTableSize),

		// MemTableStopWritesThreshold places a hard limit on the size
		// of the existent MemTables(including the frozen one).
		// Note, this must be the number of tables not the size of all memtables
		// according to https://github.com/cockroachdb/pebble/blob/master/options.go#L738-L742
		// and to https://github.com/cockroachdb/pebble/blob/master/db.go#L1892-L1903.
		MemTableStopWritesThreshold: memTableLimit,

		// The default compaction concurrency(1 thread),
		// Here use all available CPUs for faster compaction.
		MaxConcurrentCompactions: func() int { return runtime.NumCPU() },

		// Per-level options. Options for at least one level must be specified. The
		// options for the last level are used for all subsequent levels.
		// Levels: []pebble.LevelOptions{
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// 	{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		// },
		ReadOnly: readonly,
		EventListener: &pebble.EventListener{
			CompactionBegin: db.onCompactionBeginBlast,
			CompactionEnd:   db.onCompactionEndBlast,
			WriteStallBegin: db.onWriteStallBegin,
			WriteStallEnd:   db.onWriteStallEnd,
		},
		Levels:             make([]pebble.LevelOptions, numLevels),
		Logger:             panicLogger{}, // TODO(karalabe): Delete when this is upstreamed in Pebble
		FormatMajorVersion: formatVersion,
	}

	for i := 0; i < len(opt.Levels); i++ {
		l := &opt.Levels[i]
		l.BlockSize = 32 << 10       // 32 KB
		l.IndexBlockSize = 256 << 10 // 256 KB
		l.FilterPolicy = bloom.FilterPolicy(10)
		l.FilterType = pebble.TableFilter
		if i > 0 {
			l.TargetFileSize = opt.Levels[i-1].TargetFileSize * 2
		}
		l.EnsureDefaults()
	}

	// Disable seek compaction explicitly. Check https://github.com/ethereum/go-ethereum/pull/20130
	// for more details.
	opt.Experimental.ReadSamplingMultiplier = -1
	// blast adjustments
	opt.Experimental.MaxWriterConcurrency = runtime.NumCPU()
	opt.Experimental.ForceWriterParallelism = true
	opt.Experimental.SecondaryCacheSizeBytes = 64 << 20 // 64MB
	opt.TargetByteDeletionRate = 512 << 20              // 512 MB
	opt.Experimental.ReadCompactionRate = 50 << 20      // 50 mb
	// Open the db and recover any potential corruptions
	innerDB, err := pebble.Open(file, opt)
	if err != nil {
		return nil, err
	}
	db.db = innerDB

	db.compTimeMeter = metrics.NewRegisteredMeter(namespace+"compact/time", nil)
	db.compReadMeter = metrics.NewRegisteredMeter(namespace+"compact/input", nil)
	db.compWriteMeter = metrics.NewRegisteredMeter(namespace+"compact/output", nil)
	db.diskSizeGauge = metrics.NewRegisteredGauge(namespace+"disk/size", nil)
	db.diskReadMeter = metrics.NewRegisteredMeter(namespace+"disk/read", nil)
	db.diskWriteMeter = metrics.NewRegisteredMeter(namespace+"disk/write", nil)
	db.writeDelayMeter = metrics.NewRegisteredMeter(namespace+"compact/writedelay/duration", nil)
	db.writeDelayNMeter = metrics.NewRegisteredMeter(namespace+"compact/writedelay/counter", nil)
	db.memCompGauge = metrics.NewRegisteredGauge(namespace+"compact/memory", nil)
	db.level0CompGauge = metrics.NewRegisteredGauge(namespace+"compact/level0", nil)
	db.nonlevel0CompGauge = metrics.NewRegisteredGauge(namespace+"compact/nonlevel0", nil)
	db.seekCompGauge = metrics.NewRegisteredGauge(namespace+"compact/seek", nil)
	db.manualMemAllocGauge = metrics.NewRegisteredGauge(namespace+"memory/manualalloc", nil)
	log.Info("Opened pebble database", "format-major-version", innerDB.FormatMajorVersion())

	// Start up the metrics gathering and return
	go db.meter(metricsGatheringInterval, namespace)
	return db, nil
}
