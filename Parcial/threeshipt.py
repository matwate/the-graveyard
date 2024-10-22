import random

class Ship:
    def __init__(self, length, orientation):
        self.length = length
        self.orientation = orientation
        self.hits = [False] * length

    def hit(self, position):
        self.hits[position] = True

    def is_sunk(self):
        return all(self.hits)

    def get_coordinates(self, board_size):
        coordinates = []
        if self.orientation == 'H':
            row = self.hits.index(True) if True in self.hits else 0
            for col in range(len(self.hits)):
                coordinates.append((row, col))
        else:
            col = self.hits.index(True) if True in self.hits else 0
            for row in range(len(self.hits)):
                coordinates.append((row, col))
        return coordinates


class GameBoard:
    def __init__(self, size):
        self.size = size
        self.board = [['-'] * size for _ in range(size)]
        self.ships = []

    def place_ship(self, ship, row, col):
        if ship.orientation == 'H':
            if col + ship.length > self.size:
                return False
            for i in range(col, col + ship.length):
                if self.board[row][i] != '-':
                    return False
            for i in range(col, col + ship.length):
                self.board[row][i] = 'S'
        else:
            if row + ship.length > self.size:
                return False
            for i in range(row, row + ship.length):
                if self.board[i][col] != '-':
                    return False
            for i in range(row, row + ship.length):
                self.board[i][col] = 'S'
        self.ships.append(ship)
        return True

    def get_shot(self, row, col):
        if self.board[row][col] == '-':
            return False
        elif self.board[row][col] == 'X':
            return False
        else:
            print(f"Attempting to find a ship at ({row}, {col})")
            for ship in self.ships:
                for r, c in ship.get_coordinates(self.size):
                    if r == row and c == col:
                        index = ship.orientation == 'H' and c - col or r - row
                        ship.hit(index)
                        if ship.is_sunk():
                            for r, c in ship.get_coordinates(self.size):
                                self.board[r][c] = 'X'
                            self.ships.remove(ship)
                        else:
                            self.board[row][col] = 'H'
                        return True
        return False

    def all_ships_sunk(self):
        return len(self.ships) == 0

    def is_position_available(self, row, col, length, orientation):
        if orientation == 'H':
            if col + length > self.size:
                return False
            for i in range(col, col + length):
                if self.board[row][i] != '-':
                    return False
        else:
            if row + length > self.size:
                return False
            for i in range(row, row + length):
                if self.board[i][col] != '-':
                    return False
        # Verificar que las celdas no estén ocupadas por otros barcos
        if orientation == 'H':
            for i in range(col, col + length):
                for ship in self.ships:
                    for r, c in ship.get_coordinates(self.size):
                        if r == row and c == i:
                            return False
        else:
            for i in range(row, row + length):
                for ship in self.ships:
                    for r, c in ship.get_coordinates(self.size):
                        if r == i and c == col:
                            return False
        return True
    def print_board(self):
        for row in self.board:
            print(' '.join(row))

class Player:
    def __init__(self, board_size):
        self.board = GameBoard(board_size)
        self.shots = []

    def place_ships(self, ship_lengths):
        for length in ship_lengths:
            attempts = 0
            max_attempts = 100  # Establecer un límite razonable de intentos
            ship_placed = False
            while attempts < max_attempts:
                orientation = random.choice(['H', 'V'])
                row, col = self.get_random_position(length, orientation)
                if self.board.is_position_available(row, col, length, orientation):
                    ship = Ship(length, orientation)
                    if self.board.place_ship(ship, row, col):
                        ship_placed = True
                        break
                attempts += 1
            if not ship_placed:
                print(f"No se pudo colocar un barco de longitud {length}")
                break

    def get_random_position(self, length, orientation):
        if orientation == 'H':
            row = random.randint(0, self.board.size - 1)
            col = random.randint(0, self.board.size - length)
        else:
            row = random.randint(0, self.board.size - length)
            col = random.randint(0, self.board.size - 1)
        return row, col

    def get_shot(self, row, col):
        if (row, col) in self.shots:
            return False
        self.shots.append((row, col))
        return self.board.get_shot(row, col)


class Game:
    def __init__(self, board_size, ship_lengths):
        self.players = [Player(board_size) for _ in range(3)]
        for player in self.players:
            player.place_ships(ship_lengths)

    def print_boards(self):
        for i, player in enumerate(self.players):
            print(f"Player {i + 1} board:")
            player.board.print_board()
            print()

    def play(self):
        current_player = 0
        print("This line executes")
        while True:
            player = self.players[current_player]
            self.print_boards()
            print(f"Player {current_player + 1}'s turn")
            print(f"Current player: {current_player + 1}")
            target_player = int(input("Enter player to attack (1-3, excluding yourself): "))
            while target_player == current_player + 1 or target_player < 1 or target_player > 3:
                target_player = int(input("Invalid input. Enter player to attack (1-3, excluding yourself): "))
            target_player -= 1
            print(f"Target player: {target_player + 1}")
            row = int(input("Enter row: "))
            col = int(input("Enter col: "))
            hit = self.players[target_player].get_shot(row, col)
            if hit:
                print(f"Player {current_player + 1} hit a ship!")
                if self.players[target_player].board.all_ships_sunk():
                    print(f"Player {current_player + 1} wins!")
                    return
                current_player = target_player
            else:
                print("Miss!")
                current_player = (current_player + 1) % 3

if __name__ == "__main__":
    print("Initializing game...")
    game = Game(10, [5, 4, 3, 2, 2])
    print("Game initialized")
    game.play()