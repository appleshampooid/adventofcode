import sys

with open(sys.argv[1]) as instructions_input:
    instructions = list(map(lambda s: s.rstrip().split(), instructions_input.readlines()))

def run(instructions):
    maybe_corrupted = []
    instruction_pointer = 0
    accumulator = 0
    while True:
        if instruction_pointer == len(instructions):
            return accumulator
        try:
            if maybe_corrupted.index(instruction_pointer):
                return maybe_corrupted
        except ValueError:
            maybe_corrupted.append(instruction_pointer)
        instruction = instructions[instruction_pointer]
        step = 1
        arg = int(instruction[1])
        if instruction[0] == 'acc':
            accumulator += arg
        elif instruction[0] == 'jmp':
            step = arg
        instruction_pointer += step

def swap(instruction):
    if instruction[0] == 'jmp':
        instruction[0] = 'nop'
    else:
        instruction[0] = 'jmp'

maybe_corrupted = run(instructions)
while True:
    maybe_fix = maybe_corrupted.pop()
    if instructions[maybe_fix][0] == 'acc':
        continue
    swap(instructions[maybe_fix])
    ret = run(instructions)
    if type(ret) == int:
        print(f'We fixed it! The accumulator is {ret}')
        break
    else:
        swap(instructions[maybe_fix])
