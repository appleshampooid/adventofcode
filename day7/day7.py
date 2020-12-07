# only Knuth can judge me, okay?
import sys
import re
from bag import Bag

sub_bag_re = re.compile('^(\d+) ([\w ]+) bags?$')
root_bags = {}
with open(sys.argv[1]) as bag_rules_input:
    for bag_rule in bag_rules_input:
        (color, sub_bag_rules) = bag_rule.split(' bags contain ')
        sub_bags = {}
        for sub_bag_rule in sub_bag_rules.rstrip('\n.').split(', '):
            # print(sub_bag_rule)
            if sub_bag_rule != 'no other bags':
                (count, sub_color) = sub_bag_re.match(sub_bag_rule).groups()
                sub_bags[Bag(sub_color, {})] = int(count)
                # print(f'count is {count}, color is {color}')
        # print(sub_bags)
        root_bags[color] = Bag(color, sub_bags)

def contains(bag, color):
    for bag in bag.bags.keys():
        if bag.color == color:
            return True
        if contains(root_bags[bag.color], color):
            return True

def total_inside(bag):
    total = 0
    for (bag, count) in bag.bags.items():
        total += count
        total += count * total_inside(root_bags[bag.color])
    return total

total_can_hold_shiny = 0
for bag in root_bags.values():
    if contains(bag, 'shiny gold'):
        total_can_hold_shiny += 1
        print(bag)
print(f'{total_can_hold_shiny} bags can hold a shiny gold bag')

total_inside_shiny = total_inside(root_bags['shiny gold'])
print(f'{total_inside_shiny} bags are inside a shiny gold bag')
