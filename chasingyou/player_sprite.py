import pygame

class PlayerSprite(pygame.sprite.Sprite):
    """
    x, y used in init are related to the position of the player in game
    """
    def __init__(self, x: int, y: int, name: str, file_name: str, alive: bool = True, has_bomb: bool = False):
        super().__init__()
        self.x = x
        self.y = y
        self.file_name = file_name
        self.name = name
        self.has_bomb = has_bomb
        self.alive = alive
