"""
Performance check on the fluid module
Settings: step() executed 5 times

fluid               : 119_677_044 function calls in 37.972 seconds
cached by lru_cache : 7_529_542 function calls in 11.470 seconds
diy memoization     : 40_259_525 function calls (40226733 primitive calls) in 20.035 seconds
numba               : 33_436_012 function calls (32812989 primitive calls) in 13.193 seconds
numpy               : TODO
cython              : TODO
"""
from fluid import Fluid
import cProfile
import pstats

fluid1 = Fluid(128, 0.1, 0, 0, 15)

fluid1.add_density(10, 10, 100)


def main():
    for _ in range(5):
        fluid1.step()


cProfile.run("main()", "output.dat")

with open("outpout_time.txt", "w") as f:
    p = pstats.Stats("output.dat", stream=f)
    p.sort_stats("time").print_stats()

with open("output_calls.txt", "w") as f:
    p = pstats.Stats("output.dat", stream=f)
    p.sort_stats("calls").print_stats()
