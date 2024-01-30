import math
from tabulate import tabulate
import matplotlib.pyplot as plt



class YieldSimulator:
    def __init__(self, _initSharePrice, _users, _value):

        self.sharePrice = _initSharePrice
        self.users = _users;
        avgPricePerUser = math.floor(_value / self.users)
        self.totalValueLocked = avgPricePerUser * self.users
        self.sharesPerUser = avgPricePerUser // self.sharePrice
        self.remainderPerUser = avgPricePerUser % self.sharePrice

        self.count = self.sharesPerUser * self.users
        self.pending = 0
        self.timestamp = 0
        self.data = []

    def addYield(self, amount, timestamp):
        self.pending += amount
        self.totalValueLocked += amount
        self.timestamp = timestamp
        if self.pending >= self.count:
            self.sharePrice += self.pending // self.count
            self.pending = self.pending % self.count
        wasted = self.totalValueLocked - self.getYieldBearingValue()
        self.data.append((self.timestamp, round(self.sharePrice / 1e9,  6), round(self.pending / 1e18, 3), self.count, round(self.totalValueLocked / 1e18, 1), round(self.getYieldBearingValue() / 1e18, 1), round(wasted / 1e18, 1)))

    def addThreePercent(self, timestamp):
        if timestamp <= self.timestamp:
            return
        daysElapsed = timestamp - self.timestamp
        rate = 3.7 / 100
        continuousCompoundInterest = self.totalValueLocked * math.exp(rate * daysElapsed / 365) - self.totalValueLocked
        if timestamp % 20 == 0:
            print(continuousCompoundInterest / 1e18)
        self.addYield(continuousCompoundInterest, timestamp)

    def getYieldBearingValue(self):
        return self.sharePrice * self.count

    def pretty_print_data(self):
        d = [a for a in self.data if a[0] % 100 == 0]
        print(tabulate(d, tablefmt="plain", headers=['day#', 'priceInGwei', 'pendingEther', 'count', 'tvl', 'yieldBearing', 'non']))

    def plot(self):
        timestamps = [a[0] for a in self.data][0:100]
        pending = [a[5] for a in self.data][0:100]
        plt.plot(timestamps, pending)
        plt.show()


blastUsers = 100_000
blastTvl = 1e9*1e18 # 1 bil

def simWrapper(initSharePrice, initUsers, initTvl):
    sim = YieldSimulator(initSharePrice, initUsers, initTvl)
    for t in range(1, 365*10):
        if t % 100 == 0:
            sim.addThreePercent(t)
        else:
            sim.addThreePercent(t)
    print("initial sharePrice gwei:", round(initSharePrice/ 1e9, 6))
    print("initial tvl", round(blastTvl / 1e18, 0))
    print("initial users", blastUsers)
    sim.pretty_print_data()
    sim.plot()
    # Assuming 'x' and 'y' are the two lists of data

def sim1():
    sharePrice = 1e9
    simWrapper(sharePrice, blastUsers, blastTvl)

def sim2():
    sharePrice = 1e4
    simWrapper(sharePrice, blastUsers, blastTvl)

def sim3():
    sharePrice = 1e9
    users = 1_000_000
    tvl = 1e10 * 1e18
    simWrapper(sharePrice, users, tvl)

sim3()

