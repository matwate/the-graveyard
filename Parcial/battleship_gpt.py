import random

class Player:
    def __init__(self, name):
        self.name = name
        self.board = [['-' for _ in range(10)] for _ in range(10)]
        self.ships = {'A': 5, 'B': 4, 'S': 3, 'D': 3, 'P': 2}
        self.remaining_ships = len(self.ships)

    def place_ships(self):
        for ship, length in self.ships.items():
            while True:
                orientation = random.choice(['horizontal', 'vertical'])
                if orientation == 'horizontal':
                    x = random.randint(0, 9)
                    y = random.randint(0, 9 - length)
                    if all(self.board[x][y + i] == '-' for i in range(length)):
                        for i in range(length):
                            self.board[x][y + i] = ship[0]
                        break
                else:
                    x = random.randint(0, 9 - length)
                    y = random.randint(0, 9)
                    if all(self.board[x + i][y] == '-' for i in range(length)):
                        for i in range(length):
                            self.board[x + i][y] = ship[0]
                        break

    def check_guess(self, x, y):
        if self.board[x][y] != '-':
            ship_hit = self.board[x][y]
            self.board[x][y] = 'X'
            self.ships[ship_hit + ""] -= 1
            if self.ships[ship_hit] == 0:
                print(f"{self.name} has sunk the {ship_hit}")
                self.remaining_ships -= 1
                if self.remaining_ships == 0:
                    print(f"{self.name} has won!")
            return True
        else:
            print("Miss!")
            return False

    def __str__(self):
        return f"{self.name} has {self.remaining_ships} ships remaining."
class BattleshipGame:
    def __init__(self, players):
        self.players = players
        self.current_player_index = 0
        self.turn = 0
        self.reverse = False
        self.next_player = 1
    def switch_turn(self, reverse=False):
        self.reverse = reverse
        print(f"order reversed: {self.reverse}")
        if reverse:
            self.current_player_index = (self.current_player_index - 1 ) % len(self.players)
            
        else:
            self.current_player_index = (self.current_player_index + 1) % len(self.players)
        self.turn += 1

    def take_turn(self):
        next_player = self.next_player
        if self.reverse:
            next_player = -next_player
        print(self.current_player_index)
        print(*self.players, sep="\n")
        current_player = self.players[self.current_player_index]
        print(f"It's {current_player.name}'s turn to attack.")
        print(f"Board of {current_player.name} before attack:")
        for row in current_player.board:
            print(" ".join(row))
        print()
        
        for i, player in enumerate(self.players):
            if i != self.current_player_index:
                print(f"Player {i + 1}'s board:")
                for row in player.board:
                    print(" ".join(row))
                print()
        print(f"Attacking  {self.players[(self.current_player_index + next_player) % len(self.players)].name}'s  board:")
        guess = input("Enter coordinates to guess (e.g., A5): ").upper()
        x = ord(guess[0]) - ord('A')
        y = int(guess[1:]) - 1
        
        
        if self.players[(self.current_player_index + next_player) % len(self.players)].check_guess(x, y):
            self.next_player *= -1
            self.switch_turn(reverse= not (next_player < 0)  )
        else:
            self.switch_turn()

    def play(self):
        for player in self.players:
            player.place_ships()
            print(f"{player.name}'s board after placing ships:")
            for row in player.board:
                print(" ".join(row))
            print()

        while True:
            self.take_turn()
            if self.players[(self.current_player_index + 1) % len(self.players)].remaining_ships == 0:
                break
            

# Test the game with three players
player1 = Player("Player 1")
player2 = Player("Player 2")
player3 = Player("Player 3")
game = BattleshipGame([player1, player2, player3])
print("--------------------------------------------")
game.play()
