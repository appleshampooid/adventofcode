expenses = []
answer_part1 = None
answer_part2 = None
with open('./input') as expenses_input:
    for expense in expenses_input:
        current_expense = int(expense)
        for past_expense in expenses:
            if past_expense + current_expense == 2020:
                answer_part1 = past_expense * current_expense
                print(f'The answer to part 1 is {answer_part1} from {past_expense} and {current_expense}')
                break
        expenses.append(current_expense)

for i, first_expense in enumerate(expenses):
    for j, second_expense in enumerate(expenses[i+1:]):
        for third_expense in expenses[i+j+2:]:
            # print(f'i is {i}, j is {j}')
            # print(f'trying {first_expense} and {second_expense} and {third_expense}')
            if first_expense + second_expense + third_expense == 2020:
                answer_part2 = first_expense * second_expense * third_expense
                print(f'The answer to part 2 is {answer_part2} , from {first_expense} and {second_expense} and {third_expense}')
                break
        if answer_part2:
            break
    if answer_part2:
        break
