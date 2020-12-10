import sys

tree_map = []
width = None

with open(sys.argv[1]) as trees_input:
    for tree_line in trees_input:
        tree_line = tree_line.rstrip()
        if not width:
            width = len(tree_line)
        else:
            if width != len(tree_line):
                raise Exception('Lines of unequal length! WTF, mate?!')
        tree_map.append(tree_line)
print(tree_map)
print(width)

slopes = [[3,1],
          [1,1],
          [5,1],
          [7,1],
          [1,2]]
slope_product = 1
for slope in slopes:
    tree_count = 0
    (x, y) = (0, 0)
    while y < len(tree_map):
        # print(f'x {x} y {y}')
        if tree_map[y][x % width] == '#':
            # print('hit')
            tree_count += 1
        x += slope[0]
        y += slope[1]
    slope_product *= tree_count
    print(f'For slope {slope}, tree_count is {tree_count}')
print(f'slope_product {slope_product}')
