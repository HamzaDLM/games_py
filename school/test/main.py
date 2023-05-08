import pygame

pygame.init()

screen_width = 1300
screen_height = 900
screen_size = (screen_width, screen_height)
FPS = 40
game_display = pygame.display.set_mode(screen_size)
clock = pygame.time.Clock()

bg = pygame.image.load("bg.jpg")
bg_smaller = pygame.transform.scale(bg, (screen_width, screen_height))

music = pygame.mixer.music.load("song.mp3")
pygame.mixer.music.play(-1)


class Character:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.coords = (634, 346, 154, 160)
        self.size = (634, 346)

    def display_character(self):
        sprite_image = pygame.image.load("dofus_characters.png").convert_alpha()
        sprite = pygame.Surface(self.size, pygame.SRCALPHA)
        sprite.blit(sprite_image, (0, 0), (self.coords))
        scaled_sprite = pygame.transform.scale(sprite, (400, 400))
        scaled_sprite.set_colorkey((230, 230, 230))

        return scaled_sprite


boss1 = Character(0, 0)


def move(character):
    pressed_keys = pygame.key.get_pressed()
    if pressed_keys[pygame.K_RIGHT]:
        character.x += 4
    if pressed_keys[pygame.K_LEFT]:
        character.x -= 4
    if pressed_keys[pygame.K_UP]:
        character.y -= 4
    if pressed_keys[pygame.K_DOWN]:
        character.y += 4


while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            break

    # game_display.fill((33, 128, 207))
    game_display.blit(bg_smaller, (0, 0))

    game_display.blit(boss1.display_character(), (boss1.x, boss1.y))
    move(boss1)

    pygame.display.flip()
    clock.tick(FPS)
