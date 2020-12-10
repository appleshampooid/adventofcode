# once again, this code suuuuuucks. Facking kids
import sys
from functools import reduce

with open(sys.argv[1]) as forms_input:
    group_forms = forms_input.read().split('\n\n')
any_yes_sum = 0
every_yes_sum = 0
for forms in group_forms:
    # print(f'{forms}')
    all_person_answers = []
    any_yes_answers = {}
    person_answers = forms.rstrip().split('\n')
    # print(f'{person_answers}')
    for answers in person_answers:
        my_answers = {}
        for answer in answers:
            any_yes_answers[answer] = True
            my_answers[answer] = True
        all_person_answers.append(my_answers)
    any_yes_sum += len(any_yes_answers)
    for char in 'abcdefghijklmnopqrstuvwxyz':
        yes = list(map(lambda answer_dict: char in answer_dict, all_person_answers))
        all_yes = reduce(lambda x, y: x and y, yes)
        if all_yes:
            # print(f'for {char}, all are yes!')
            every_yes_sum += 1

print(f'Part 1 sum is {any_yes_sum}')
print(f'Part 2 sum is {every_yes_sum}')
