import sys

with open(sys.argv[1]) as joltage_input:
    joltages = sorted(map(int, joltage_input))
joltages.insert(0, 0)
joltages.append(joltages[-1] + 3)
diffs = {1: 0,
         2: 0,
         3: 0}
for i, joltage in enumerate(joltages[1:], 1):
    diffs[joltage - joltages[i-1]] +=1
print(diffs)
print(f'Part 1 answer is {diffs[1] * diffs[3]}')
