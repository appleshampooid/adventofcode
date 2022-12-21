import sys
import re

coordinates = [0, 0]
#ugh, DEGREES.
direction = 0
instruction_regex = re.compile('([NSEWRLF])(\d+)')
with open(sys.argv[1]) as instructions_input:
    for instruction in instructions_input:
        (action, amount) = instruction_regex.match(instruction).groups()
        amount = int(amount)
        if action == 'N':
            coordinates[1] += amount
        elif action == 'S':
            coordinates[1] -= amount
        elif action == 'W':
            coordinates[0] -= amount
        elif action == 'E':
            coordinates[0] += amount
        elif action == 'L':
            direction  = (direction + amount) % 360
        elif action == 'R':
            direction = (direction - amount) % 360
        elif action == 'F':
            if direction == 0:
                coordinates[0] += amount
            elif direction == 90:
                coordinates[1] += amount
            elif direction == 180:
                coordinates[0] -= amount
            elif direction == 270:
                coordinates[1] -= amount

print(f'coordinates are {coordinates}')
distance = abs(coordinates[0]) + abs(coordinates[1])
print(f'Manhattan distance is {distance}')
