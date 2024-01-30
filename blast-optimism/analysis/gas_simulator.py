from tabulate import tabulate

class GasClaimSimulator:
    def __init__(self, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate):
        self.zeroClaimRate = _zeroClaimRate
        self.baseGasSecondsPerWei = _baseGasSeconds
        self.baseClaimRate = _baseClaimRate
        self.ceilGasSecondsPerWei = _ceilGasSeconds
        self.ceilClaimRate = _ceilClaimRate

        self.balance = 0
        self.gasSeconds = 0
        self.timestamp = 0

        self.data = []
        self.data.append((self.timestamp, self.gasSeconds, self.gasSeconds, self.balance, 0, 0, 0))

    def update(self, timestamp):
        if timestamp < self.timestamp:
            raise ValueError("Timestamp cannot be less than the current timestamp")
        self.gasSeconds += (timestamp - self.timestamp)*self.balance
        self.timestamp = timestamp 

    def earn(self, balanceToEarn, timestamp, verbose = True):
        oldGasSeconds = self.gasSeconds
        self.update(timestamp)
        self.balance += balanceToEarn
        if verbose: 
            self.data.append((self.timestamp, oldGasSeconds, self.gasSeconds, self.balance, 0, 0, 0))

    def getClaimRateBps(self, gasSecondsToConsume, gasToClaim):
        secondsStaked = gasSecondsToConsume / gasToClaim
        if secondsStaked < self.baseGasSecondsPerWei:
            return self.zeroClaimRate
        if secondsStaked > self.ceilGasSecondsPerWei:
            return self.ceilClaimRate

        rateDiff = self.ceilClaimRate - self.baseClaimRate
        secondsDiff = self.ceilGasSecondsPerWei - self.baseGasSecondsPerWei
        secondsStakedDiff = secondsStaked - self.baseGasSecondsPerWei
        additionalClaimRate = rateDiff * secondsStakedDiff / secondsDiff
        return self.baseClaimRate + additionalClaimRate

    def claimGas(self, gasToClaim, gasSecondsToConsume, timestamp):
        if gasToClaim == 0:
            return

        oldGasSeconds = self.gasSeconds
        self.update(timestamp)

        # check validity requirements
        if gasToClaim <= 0:
            raise ValueError("must withdraw non-zero amount")
        if gasToClaim > self.balance:
            raise ValueError("too much to withdraw")
        if gasSecondsToConsume > self.gasSeconds:
            raise ValueError("not enough gas seconds")

        # get claim rate
        claim_rate = self.getClaimRateBps(gasSecondsToConsume, gasToClaim)
        
        # calculate tax
        user_ether = gasToClaim * claim_rate / 10000
        penalty = gasToClaim - user_ether

        self.balance -= gasToClaim
        self.gasSeconds -= gasSecondsToConsume


        self.data.append((self.timestamp, oldGasSeconds, self.gasSeconds, self.balance, user_ether, penalty, claim_rate))

    def claim_gas_at_min_claim_rate(self, min_claim_rate_bps, timestamp):
        if min_claim_rate_bps <= self.zeroClaimRate:
            return self.claim_all(timestamp)

        # set minClaimRate to baseClaimRate in this case
        if min_claim_rate_bps < self.baseClaimRate:
            min_claim_rate_bps = self.baseClaimRate

        bps_diff = min_claim_rate_bps - self.baseClaimRate
        seconds_diff = self.ceilGasSecondsPerWei - self.baseGasSecondsPerWei
        rate_diff = self.ceilClaimRate - self.baseClaimRate
        min_seconds_staked = self.baseGasSecondsPerWei + (bps_diff * seconds_diff / rate_diff)
        max_ether_claimable = self.gasSeconds / min_seconds_staked
        if max_ether_claimable > self.balance:
            max_ether_claimable = self.balance

        seconds_to_consume = max_ether_claimable * min_seconds_staked # TODO: check for rounding errors(?)
        return self.claimGas(max_ether_claimable, seconds_to_consume, timestamp)

    def claim_all(self, timestamp):
        self.claimGas(self.balance, self.gasSeconds, timestamp)

    def claim_free(self, timestamp):
        gasToClaim = self.gasSeconds / self.ceilGasSecondsPerWei
        gasToClaim = min(gasToClaim, self.balance)
        gasToConsume = gasToClaim * self.ceilGasSecondsPerWei
        self.claimGas(gasToClaim, gasToConsume, timestamp)

    def pretty_print_data(self, verbose = True):
        print(tabulate(self.data, tablefmt="plain", headers=['T', 'old_sec', 'new_sec', 'bal', 'earned', 'penalty', 'bps']))


# default values
zeroClaimRate = .1
baseGasSeconds = 1000
baseClaimRate = .5
ceilGasSeconds = 10000
ceilClaimRate = 1


def sim1():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    sim.earn(5, 1)
    sim.earn(5, 2)
    sim.earn(5, 3)
    sim.pretty_print_data()

def sim2():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    for t in range(0, 10000, 10):
        sim.earn(5, t)
    sim.claim_all(10000)
    sim.pretty_print_data()

def sim3():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    for t in range(0, 100000, 10):
        sim.earn(5, t)
    sim.claim_all(100000)
    sim.pretty_print_data()

def sim4():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    for t in range(0, 100000, 10):
        sim.earn(5, t, False)
        if t % 10000 == 0:
            sim.claim_all(t)
    sim.claim_all(100000)
    sim.pretty_print_data()

def sim5():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    for t in range(0, 100000, 10):
        sim.earn(5, t, False)
        if t % 10000 == 0:
            sim.claim_free(t)
    sim.claim_all(100000)
    sim.pretty_print_data()

def sim6():
    sim =  GasClaimSimulator(zeroClaimRate*10000, baseGasSeconds, baseClaimRate*10000, ceilGasSeconds, ceilClaimRate*10000)
    for t in range(0, 100000, 100):
        sim.earn(1000, t, False)
        if t % 10000 == 0:
            sim.claim_gas_at_min_claim_rate(9000, t)
    sim.pretty_print_data()

sim4()