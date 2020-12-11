# I must say it once again: ONLY KNUTH CAN JUDGE ME

import sys
from copy import deepcopy

def occupiable(seats, i, j):
    all_empty = True
    for x in range(-1, 2):
        for y in range(-1, 2):
            # negative index wrapping is for once an annoyance
            if (x == 0 and y == 0) or (i+x < 0 or j+y < 0):
                continue
            try:
                if seats[i+x][j+y] == '#':
                    all_empty = False
                    break
            except IndexError:
                continue
        if not all_empty:
            break
    return all_empty

def crowded(seats, i, j):
    # print(f'examining {i}, {j}')
    num_occupied = 0
    for x in range(-1, 2):
        for y in range(-1, 2):
            if (x == 0 and y == 0) or (i+x < 0 or j+y < 0):
                continue
            try:
                if seats[i+x][j+y] == '#':
                    # print(f'examining {i}, {j}, with {x} and {y} mods there is an occupied seat')
                    num_occupied += 1
                    if num_occupied >= 4:
                        break
            except IndexError:
                continue
        if num_occupied >= 4:
            break
    return num_occupied >= 4

def occupiable2(seats, i, j, height, width):
    all_empty = True
    for x in range(-1, 2):
        for y in range(-1, 2):
            if (x == 0 and y == 0):
                continue
            multiplier = 1
            while True:
                a = i + x * multiplier
                b = j + y * multiplier
                if a < 0 or b < 0 or a >= height or b >= width:
                    break
                if seats[a][b] == '#':
                    all_empty = False
                    break
                if seats[a][b] == 'L':
                    break
                multiplier +=1
            if not all_empty:
                break
        if not all_empty:
            break
    return all_empty

def crowded2(seats, i, j, height, width):
    num_occupied = 0
    for x in range(-1, 2):
        for y in range(-1, 2):
            if (x == 0 and y == 0):
                continue
            multiplier = 1
            while True:
                a = i + x * multiplier
                b = j + y * multiplier
                if a < 0 or b < 0 or a >= height or b >= width:
                    break
                if seats[a][b] == '#':
                    num_occupied += 1
                    break
                if seats[a][b] == 'L':
                    break
                multiplier +=1
            if num_occupied >= 5:
                break
        if num_occupied >= 5:
            break
    return num_occupied >= 5

def run_model(seats, height, width):
    new_seats = deepcopy(seats)
    for i in range(0, height):
        for j in range(0, width):
            if seats[i][j] == '.':
                continue
            elif seats[i][j] == 'L' and occupiable(seats, i, j):
                new_seats[i][j] = '#'
            elif seats[i][j] == '#' and crowded(seats, i, j):
                new_seats[i][j] = 'L'
    return new_seats

def run_model2(seats, height, width):
    new_seats = deepcopy(seats)
    for i in range(0, height):
        for j in range(0, width):
            if seats[i][j] == '.':
                continue
            elif seats[i][j] == 'L' and occupiable2(seats, i, j, height, width):
                new_seats[i][j] = '#'
            elif seats[i][j] == '#' and crowded2(seats, i, j, height, width):
                new_seats[i][j] = 'L'
    return new_seats

def print_seats(seats):
    occupied = 0
    for row in seats:
        for seat in row:
            print(seat, end='')
            if seat == '#':
                occupied += 1
        print('\n', end='')
    print('\n', end='')
    return occupied

seats = []
width = None
height = None

with open(sys.argv[1]) as seats_input:
    for seat_line in seats_input:
        seat_line = seat_line.rstrip()
        if not width:
            width = len(seat_line)
        else:
            if width != len(seat_line):
                raise Exception('Lines of unequal length! WTF, mate?!')
        seats.append(list([seat for seat in seat_line]))
height = len(seats)
print_seats(seats)

original_seats = deepcopy(seats)

part1_occupied = None
while True:
    new_seats = run_model(seats, height, width)
    part1_occupied = print_seats(new_seats)
    if new_seats == seats:
        break
    seats = new_seats

seats = original_seats
part2_occuppied = None
while True:
    new_seats = run_model2(seats, height, width)
    part2_occupied = print_seats(new_seats)
    if new_seats == seats:
        break
    seats = new_seats

print(f'Part 1: there are {part1_occupied} occupied seats at steady state')
print(f'Part 2: there are {part2_occupied} occupied seats at steady state')
