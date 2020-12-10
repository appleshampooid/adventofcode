class Seat:
    def __init__(self, definition):
        self.definition = definition
        if len(self.definition) != 10:
            raise Exception('Incorrect number of characters in seat definition!')
        self.row_spec = self.definition[0:7]
        self.column_spec = self.definition[7:]
        self.binary_spec = self.definition.replace('F', '0').replace('B', '1').replace('L', '0').replace('R', '1')
        self.id = int(self.binary_spec, 2)

    def __str__(self):
        return f'rows {self.row_spec} columns {self.column_spec} binary spec {self.binary_spec} id {self.id}'
