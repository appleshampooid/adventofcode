class Bag:
    def __init__(self, color, bags):
        self.color = color
        self.bags = bags

    def __str__(self):
        ret = f'color: {self.color}; bags: '
        ret += '{'
        for bag, count in self.bags.items():
            ret += f'{count}: {str(bag)}'
        ret += '} '
        return ret
    # these are lazy, and depend on the input being "nice", that is
    # every bag is only defined *once*
    def __hash__(self):
        hashcode = hash(self.color)
        return hashcode

    def __eq__(self, other):
        return self.color == other.color
