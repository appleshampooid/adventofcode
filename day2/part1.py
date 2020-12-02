valid_passwords = 0
with open('./input') as passwords_input:
    for password_spec in passwords_input:
        (valid_range, letter, password) = password_spec.split(' ')
        letter = letter[0]
        (min_letter_string, max_letter_string) = valid_range.split('-')
        # print(f'range is {valid_range} ; letter is {letter} ; password is {password}')
        min_letter = int(min_letter_string)
        max_letter = int(max_letter_string)
        # print(f'min_letter {min_letter} ; max_letter is {max_letter}')
        count = 0
        for char in password:
            # print(f'evaluating {char}')
            if char == letter:
                count += 1
        # print(f'count is {count}')
        if count <= max_letter and count >= min_letter:
            print(f'password {password} is valid')
            valid_passwords += 1

print(f'There are {valid_passwords} valid passwords')
