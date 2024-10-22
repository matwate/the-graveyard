import tkinter as tk
import random

from threeshipt import Ship, GameBoard, Player, Game

class BattleshipGUI(tk.Tk):
    def __init__(self):
        super().__init__()
        self.title("Battleship")
        self.geometry("600x400")

        self.game = Game(8, [4, 3, 3, 2, 2])  # Reducir el tamaño del tablero a 8x8
        self.current_player = 0
        self.target_player = None

        self.boards_frame = tk.Frame(self)
        self.boards_frame.pack(side=tk.LEFT, padx=10, pady=10)

        self.create_boards()

        self.controls_frame = tk.Frame(self)
        self.controls_frame.pack(side=tk.RIGHT, padx=10, pady=10)

        self.create_controls()

    def create_boards(self):
        for i, player in enumerate(self.game.players):
            board_frame = tk.Frame(self.boards_frame)
            board_frame.pack(side=tk.TOP, pady=10)

            title = tk.Label(board_frame, text=f"Player {i + 1} Board")
            title.pack()

            board_canvas = tk.Canvas(board_frame, width=240, height=240)
            board_canvas.pack()

            self.draw_board(board_canvas, player.board)

    def draw_board(self, canvas, board):
        canvas.delete("all")
        cell_size = 30
        for row in range(board.size):
            for col in range(board.size):
                x1 = col * cell_size
                y1 = row * cell_size
                x2 = x1 + cell_size
                y2 = y1 + cell_size
                value = board.board[row][col]
                if value == '-':
                    color = "blue"
                elif value == 'S':
                    color = "green"
                elif value == 'H':
                    color = "yellow"
                else:
                    color = "red"
                canvas.create_rectangle(x1, y1, x2, y2, fill=color, outline="white")

    def create_controls(self):
        target_label = tk.Label(self.controls_frame, text="Target Player:")
        target_label.pack()

        self.target_player_var = tk.IntVar()
        for i in range(1, 4):
            if i != self.current_player + 1:
                radio_button = tk.Radiobutton(self.controls_frame, text=f"Player {i}", variable=self.target_player_var, value=i - 1)
                radio_button.pack()

        self.row_entry = tk.Entry(self.controls_frame, width=5)
        self.row_entry.pack()

        self.col_entry = tk.Entry(self.controls_frame, width=5)
        self.col_entry.pack()

        fire_button = tk.Button(self.controls_frame, text="Fire!", command=self.fire)
        fire_button.pack()

    def fire(self):
        row = int(self.row_entry.get()) - 1  # Ajustar índices de 0 a 7
        col = int(self.col_entry.get()) - 1
        target_player = self.target_player_var.get()

        hit = self.game.players[target_player].get_shot(row, col)
        if hit:
            print(f"Player {self.current_player + 1} hit a ship!")
            if self.game.players[target_player].board.all_ships_sunk():
                print(f"Player {self.current_player + 1} wins!")
                # Agregar lógica para finalizar el juego
            self.current_player = target_player
        else:
            print("Miss!")
            self.current_player = (self.current_player + 1) % 3

        self.update_boards()

    def update_boards(self):
        for i, player in enumerate(self.game.players):
            board_canvas = self.boards_frame.winfo_children()[i].winfo_children()[1]
            self.draw_board(board_canvas, player.board)

if __name__ == "__main__":
    app = BattleshipGUI()
    app.mainloop()