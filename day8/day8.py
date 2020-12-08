import sys

with open(sys.argv[1]) as instructions_input:
    instructions = list(map(lambda s: s.rstrip().split(), instructions_input.readlines()))

accumulator = 0
visited_instructions = [0] * len(instructions)
#print(visited_instructions)
instruction_pointer = 0
while True:
    instruction = instructions[instruction_pointer]
    visited_instructions[instruction_pointer] += 1
    if visited_instructions[instruction_pointer] == 2:
        break
    step = 1
    arg = int(instruction[1])
    if instruction[0] == 'acc':
        accumulator =+ arg
    elif instruction[0] == 'jmp':
        step = arg
    instruction_pointer += step
print(f'the accumulator is {accumulator}')
