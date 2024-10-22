import unittest
from threeshipt import Ship, GameBoard, Player, Game

class TestShip(unittest.TestCase):
    def test_hit(self):
        ship = Ship(5, 'H')
        ship.hit(2)
        self.assertTrue(ship.hits[2])

    def test_is_sunk(self):
        ship = Ship(3, 'V')
        self.assertFalse(ship.is_sunk())
        ship.hit(0)
        ship.hit(1)
        ship.hit(2)
        self.assertTrue(ship.is_sunk())

    def test_get_coordinates(self):
        ship = Ship(4, 'H')
        ship.hit(1)
        ship.hit(3)
        expected_coordinates = [(0, 1), (0, 2), (0, 3), (0, 4)]
        self.assertEqual(ship.get_coordinates(10), expected_coordinates)

class TestGameBoard(unittest.TestCase):
    def test_place_ship(self):
        board = GameBoard(10)
        ship = Ship(3, 'H')
        self.assertTrue(board.place_ship(ship, 0, 0))
        self.assertFalse(board.place_ship(ship, 0, 0))

    def test_get_shot(self):
        board = GameBoard(10)
        ship = Ship(3, 'H')
        board.place_ship(ship, 0, 0)
        self.assertTrue(board.get_shot(0, 0))
        self.assertFalse(board.get_shot(0, 0))

    def test_is_position_available(self):
        board = GameBoard(10)
        ship = Ship(3, 'H')
        board.place_ship(ship, 0, 0)
        self.assertFalse(board.is_position_available(0, 0, 3, 'H'))
        self.assertTrue(board.is_position_available(0, 3, 3, 'H'))

    def test_all_ships_sunk(self):
        board = GameBoard(10)
        ship1 = Ship(3, 'H')
        ship2 = Ship(4, 'V')
        board.place_ship(ship1, 0, 0)
        board.place_ship(ship2, 2, 2)
        self.assertFalse(board.all_ships_sunk())
        board.get_shot(0, 0)
        board.get_shot(0, 1)
        board.get_shot(0, 2)
        board.get_shot(2, 2)
        board.get_shot(3, 2)
        board.get_shot(4, 2)
        board.get_shot(5, 2)
        self.assertTrue(board.all_ships_sunk())

class TestPlayer(unittest.TestCase):
    def test_place_ships(self):
        player = Player(10)
        player.place_ships([5, 4, 3, 2, 2])
        ships_placed = sum(len(row) for row in player.board.board for cell in row if cell == 'S')
        self.assertEqual(ships_placed, 16)

    def test_get_shot(self):
        player = Player(10)
        ship = Ship(3, 'H')
        player.board.place_ship(ship, 0, 0)
        self.assertTrue(player.get_shot(0, 0))
        self.assertFalse(player.get_shot(0, 0))

class TestGame(unittest.TestCase):
    def test_play(self):
        # Aquí puedes simular una partida completa y verificar el resultado
        pass

if __name__ == '__main__':
    unittest.main()