import sys

known_combinations = {}

def all_connections(joltages):
    if joltages[0] in known_combinations:
        return known_combinations[joltages[0]]
    print(f'Examining {joltages}')
    if len(joltages) == 1:
        return 1
    combos = 0
    for i, joltage in enumerate(joltages[1:4]):
        if joltage - joltages[0] <= 3:
            combos += all_connections(joltages[i+1:])
    known_combinations[joltages[0]] = combos
    return combos

with open(sys.argv[1]) as joltage_input:
    joltages = sorted(map(int, joltage_input))
joltages.insert(0, 0)
joltages.append(joltages[-1] + 3)
print(joltages)
diffs = {1: 0,
         2: 0,
         3: 0}
for i, joltage in enumerate(joltages[1:], 1):
    diffs[joltage - joltages[i-1]] +=1
print(diffs)
print(f'Part 1 answer is {diffs[1] * diffs[3]}')

print(f'Part 2 answer is {all_connections(joltages)}')
