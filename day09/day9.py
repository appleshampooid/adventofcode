import sys
from itertools import combinations
preamble_size = int(sys.argv[2])
ciphertext = []

def valid(x, addends):
    # print(f'x is {x}, addends are {addends}')
    for (y, z) in combinations(addends, 2):
        if y + z == x:
            return True
    return False

invalid = None
with open(sys.argv[1]) as xmas_input:
    for i, line in enumerate(xmas_input):
        x = int(line)
        ciphertext.append(x)
        if i >= preamble_size and not invalid:
            if not valid(x, ciphertext[i - preamble_size:i]):
                invalid = x
                print(f'Part 1: the invalid number is {invalid}')

for i, x in enumerate(ciphertext):
    sum = x
    smallest = x
    largest = x
    for y in ciphertext[1+i:]:
        if y > largest:
            largest = y
        if y < smallest:
            smallest = y
        sum += y
        if sum == invalid:
            print(f'Part 2: sum is {smallest + largest} from {smallest} and {largest}; range is from {x} to {y}')
            sys.exit(0)
        elif sum > invalid:
            break
