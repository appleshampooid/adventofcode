import sys
from seat import Seat

highest = 0
ids = [None] * 1024
with open(sys.argv[1]) as seats_input:
    for seat in seats_input:
        seat = Seat(seat.rstrip())
        print(seat)
        if seat.id >= highest:
            highest = seat.id
        ids[seat.id] = 1
started = False
for i, seat_id in enumerate(ids):
    if not started:
        if seat_id:
            started = True
    else:
        if not seat_id:
            my_seat = i
            break

print(f'{ids}')
print(f'the highest id is {highest}')
print(f'my seat is {my_seat}')
