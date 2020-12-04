# This code sucks. LOL.
import sys
import re

required_fields = set(['byr',
                       'iyr',
                       'eyr',
                       'hgt',
                       'hcl',
                       'ecl',
                       'pid'])
#cid
height_regex = re.compile('(\d+)(cm|in)')
hcl_regex = re.compile('#[0-9a-f]{6}')
ecl_regex = re.compile('amb|blu|brn|gry|grn|hzl|oth')
pid_regex = re.compile('\d{9}')
with open(sys.argv[1]) as passports_input:
    passport_data= passports_input.read()
passports = passport_data.split('\n\n')
valid_passports = 0
for (i, passport) in enumerate(passports):
    print(f'working on {i}')
    valid = True
    fields = passport.split()
    keys = set()
    for field in fields:
        (key,value) = field.split(':')
        if key == 'byr':
            year = int(value)
            if not (len(value) == 4 and year >= 1920 and year <= 2002):
                valid = False
                print(f'invalid byr on {i}')
                continue
        elif key == 'iyr':
            year = int(value)
            if not (len(value) == 4 and year >= 2010 and year <= 2020):
                valid = False
                print(f'invalid iyr on {i}')
                continue
        elif key == 'eyr':
            year = int(value)
            if not (len(value) == 4 and year >= 2020 and year <= 2030):
                valid = False
                print(f'invalid eyr on {i}')
                continue
        elif key == 'hgt':
            match = height_regex.match(value)
            if match:
                unit = match.group(2)
                size = int(match.group(1))
                if unit == 'cm':
                    if not (size >= 150 and size <= 193):
                        valid = False
                        print(f'invalid hgt:cm on {i}')
                        continue
                elif unit == 'in':
                    if not (size >= 59 and size <= 76):
                        valid = False
                        print(f'invalid hgt:in on {i}')
                        continue
            else:
                print(f'invalid hgt on {i}')
                valid = False
                continue
        elif key == 'hcl':
            if not hcl_regex.match(value):
                print(f'invalid hcl on {i}')
                valid = False
                continue
        elif key == 'ecl':
            if not ecl_regex.match(value):
                print(f'invalid ecl on {i}')
                valid = False
                continue
        elif key == 'pid':
            if not (pid_regex.match(value) and len(value) == 9):
                print(f'invalid pid on {i}')
                valid = False
                continue
        keys.add(key)
    if not valid:
        print(f'invalid values {i}')
        continue
    if keys.issuperset(required_fields):
        valid_passports += 1
        print(f'A valid one at {i}: {fields}')
    else:
        print(f'invalid fields {i}')

print(f'There are {valid_passports} valid passports.')
# 146 too high
