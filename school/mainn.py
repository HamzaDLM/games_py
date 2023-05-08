import pygame


pygame.init()

SCREEN_WIDTH = 1200
SCREEN_HIGHT = 900
SCREEN = (SCREEN_WIDTH, SCREEN_HIGHT)
FPS = 40

game_display = pygame.display.set_mode(SCREEN)

clock = pygame.time.Clock()

bg = pygame.image.load("bg.jpg")
# bg = pygame.transform.scale(bg, (SCREEN_HIGHT, SCREEN_HIGHT))

# tree = pygame.image.load("tree.png")
# tree = pygame.transform.scale(tree, (90, 120))


class Character:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def display_character(self):
        sp_x, sp_y, sp_w, sp_h = (688, 102, 409, 506)
        sprite_image = pygame.image.load("blue.png").convert_alpha()
        # self.rect = self.image.get_rect()
        sprite = pygame.Surface((sp_w, sp_h), pygame.SRCALPHA)
        sprite.blit(sprite_image, (0, 0), (sp_x, sp_y, sp_w, sp_h))
        scaled_sprite = pygame.transform.scale(sprite, (70, 70))
        scaled_sprite.set_colorkey((135, 132, 181))

        return scaled_sprite


riven = Character(100, 200)

while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            break

    # game_display.fill((112, 91, 207))
    game_display.blit(bg, (0, 0))
    game_display.blit(riven.display_character(), (100, 100))
    pygame.display.flip()
    clock.tick(FPS)

pygame.QUIT()
