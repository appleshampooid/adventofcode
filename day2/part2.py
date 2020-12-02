valid_passwords = 0
with open('./input') as passwords_input:
    for password_spec in passwords_input:
        (valid_range, letter, password) = password_spec.split(' ')
        letter = letter[0]
        (first_letter_string, second_letter_string) = valid_range.split('-')
        # print(f'range is {valid_range} ; letter is {letter} ; password is {password}')
        first_letter = int(first_letter_string)
        second_letter = int(second_letter_string)
        # print(f'min_letter {min_letter} ; max_letter is {max_letter}')
        if (password[first_letter - 1] == letter) ^ (password[second_letter - 1] == letter):
            print(f'password {password} is valid')
            valid_passwords += 1

print(f'There are {valid_passwords} valid passwords')
