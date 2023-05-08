import pygame


pygame.init()

SCREEN_WIDTH = 1440
SCREEN_HIGHT = 960
SCREEN = (SCREEN_WIDTH, SCREEN_HIGHT)
FPS = 40
bg = pygame.image.load("bg.webp")
game_display = pygame.display.set_mode(SCREEN)
clock = pygame.time.Clock()
music = pygame.mixer.music.load("wmd_ficus.mp3")


class Character:
    def __init__(self, x, y, image, coords, size):
        self.x = x
        self.y = y
        self.coords = coords
        self.size = size
        self.image = image

    def display_character(self):
        sprite_image = pygame.image.load(self.image).convert_alpha()
        sprite = pygame.Surface(self.size, pygame.SRCALPHA)
        sprite.blit(sprite_image, (0, 0), (*self.coords, *self.size))
        scaled_sprite = pygame.transform.scale(sprite, (150, 150))
        scaled_sprite.set_colorkey((255, 255, 255))
        return scaled_sprite


ferry = Character(0, 0, "characters.png", (101, 658), (317, 506))

pygame.mixer.music.play(-1)
# MAIN GAME LOOP
while True:
    if pygame.event in pygame.event.get():
        if event.type == game.QUIT:
            break

    # game_display.fill((200, 200, 200))
    game_display.blit(bg, (0, 0))

    game_display.blit(ferry.display_character(), (ferry.x, ferry.y))

    pygame.display.flip()
    clock.tick(FPS)

pygame.QUIT()
