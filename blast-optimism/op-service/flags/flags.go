package flags

import (
	"fmt"
	"strings"

	"github.com/ethereum-optimism/optimism/op-node/chaincfg"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	openum "github.com/ethereum-optimism/optimism/op-service/enum"
	"github.com/ethereum-optimism/optimism/op-service/forks"
	"github.com/urfave/cli/v2"
)

const (
	RollupConfigFlagName                  = "rollup.config"
	NetworkFlagName                       = "network"
	CanyonOverrideFlagName                = "override.canyon"
	DeltaOverrideFlagName                 = "override.delta"
	EcotoneOverrideFlagName               = "override.ecotone"
	TaigaOverrideFlagName                 = "override.taiga"
	PectraBlobScheduleOverrideFlagName    = "override.pectrablobschedule"
	OsakaBlobScheduleOverrideFlagName     = "override.osakablobschedule"
	Bpo1BlobScheduleOverrideFlagName      = "override.bpo1blobschedule"
	Bpo2BlobScheduleOverrideFlagName      = "override.bpo2blobschedule"
	Bpo2BlastBlobScheduleOverrideFlagName = "override.bpo2Blastblobschedule"
)

func CLIFlags(envPrefix string, category string) []cli.Flag {
	return []cli.Flag{
		&cli.Uint64Flag{
			Name:     CanyonOverrideFlagName,
			Usage:    "Manually specify the Canyon fork timestamp, overriding the bundled setting",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_CANYON"),
			Hidden:   false,
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     DeltaOverrideFlagName,
			Usage:    "Manually specify the Delta fork timestamp, overriding the bundled setting",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_DELTA"),
			Hidden:   false,
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     EcotoneOverrideFlagName,
			Usage:    "Manually specify the Ecotone fork timestamp, overriding the bundled setting",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_ECOTONE"),
			Hidden:   false,
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     TaigaOverrideFlagName,
			Usage:    "Manually specify the Taiga fork timestamp, overriding the bundled setting",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_TAIGA"),
			Hidden:   false,
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     PectraBlobScheduleOverrideFlagName,
			Usage:    "Manually specify the PectraBlobSchedule fork timestamp, overriding the bundled setting",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_PECTRABLOBSCHEDULE"),
			Hidden:   false,
			Category: category,
		},
		&cli.StringFlag{
			Name:     OsakaBlobScheduleOverrideFlagName,
			Usage:    "Manually specify the osaka blob schedule, overriding the bundled setting valid options " + openum.EnumString(forks.All),
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_FUSAKABLOBSCHEDULE"),
			Hidden:   false,
			Category: category,
		},
		&cli.StringFlag{
			Name:     Bpo1BlobScheduleOverrideFlagName,
			Usage:    "Manually specify the bpo1 blob schedule, overriding the bundled setting valid options " + openum.EnumString(forks.All),
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_BPO1BLOBSCHEDULE"),
			Hidden:   false,
			Category: category,
		},
		&cli.StringFlag{
			Name:     Bpo2BlobScheduleOverrideFlagName,
			Usage:    "Manually specify the bpo2 blob schedule, overriding the bundled setting valid options " + openum.EnumString(forks.All),
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_BPO2BLOBSCHEDULE"),
			Hidden:   false,
			Category: category,
		},
		&cli.Uint64Flag{
			Name:     Bpo2BlastBlobScheduleOverrideFlagName,
			Usage:    "Manually specify the bpo2 blast blob timestamp, overriding the bundled setting valid options ",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "OVERRIDE_BPO2BLASTBLOBSCHEDULE"),
			Hidden:   false,
			Category: category,
		},

		CLINetworkFlag(envPrefix, category),
		CLIRollupConfigFlag(envPrefix, category),
	}
}

func CLINetworkFlag(envPrefix string, category string) cli.Flag {
	return &cli.StringFlag{
		Name:     NetworkFlagName,
		Usage:    fmt.Sprintf("Predefined network selection. Available networks: %s", strings.Join(chaincfg.AvailableNetworks(), ", ")),
		EnvVars:  opservice.PrefixEnvVar(envPrefix, "NETWORK"),
		Category: category,
	}
}

func CLIRollupConfigFlag(envPrefix string, category string) cli.Flag {
	return &cli.StringFlag{
		Name:     RollupConfigFlagName,
		Usage:    "Rollup chain parameters",
		EnvVars:  opservice.PrefixEnvVar(envPrefix, "ROLLUP_CONFIG"),
		Category: category,
	}
}

// This checks flags that are exclusive & required. Specifically for each
// set of flags, exactly one flag must be set.
var requiredXorFlags = [][]string{
	// TODO(client-pod#391): Re-enable this check at a later point
	// {
	// 	RollupConfigFlagName,
	// 	NetworkFlagName,
	// },
}

func CheckRequiredXor(ctx *cli.Context) error {
	for _, flagSet := range requiredXorFlags {
		var setCount int
		for _, flagName := range flagSet {
			if ctx.IsSet(flagName) {
				setCount++
			}
		}
		if setCount != 1 {
			return fmt.Errorf("exactly one of the following flags must be set: %s", strings.Join(flagSet, ", "))
		}
	}
	return nil
}
