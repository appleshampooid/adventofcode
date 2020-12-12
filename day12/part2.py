# my kingdom for 30 minutes to refactor but the kids' movie just ended

import sys
import re

ship_coordinates = [0, 0]
waypoint_relative = [10, 1]

instruction_regex = re.compile('([NSEWRLF])(\d+)')
with open(sys.argv[1]) as instructions_input:
    for instruction in instructions_input:
        (action, amount) = instruction_regex.match(instruction).groups()
        amount = int(amount)
        if action == 'N':
            waypoint_relative[1] += amount
        elif action == 'S':
            waypoint_relative[1] -= amount
        elif action == 'W':
            waypoint_relative[0] -= amount
        elif action == 'E':
            waypoint_relative[0] += amount
        elif action in ('L', 'R'):
            new_waypoint = [None, None]
            if action == 'L':
                amount = (amount * -1) % 360
            if amount == 90:
                new_waypoint[0] = waypoint_relative[1]
                new_waypoint[1] = -waypoint_relative[0]
            elif amount == 180:
                new_waypoint[0] = -waypoint_relative[0]
                new_waypoint[1] = -waypoint_relative[1]
            elif amount == 270:
                new_waypoint[0] = -waypoint_relative[1]
                new_waypoint[1] = waypoint_relative[0]
            waypoint_relative = new_waypoint
        elif action == 'F':
            ship_coordinates[0] += waypoint_relative[0] * amount
            ship_coordinates[1] += waypoint_relative[1] * amount

print(f'ship coordinates are {ship_coordinates}')
distance = abs(ship_coordinates[0]) + abs(ship_coordinates[1])
print(f'Manhattan distance is {distance}')
