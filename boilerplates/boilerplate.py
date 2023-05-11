import pygame

pygame.init()

SCREEN_HEIGHT = 800
SCREEN_WIDTH = 1000
SCREEN_SIZE = (SCREEN_WIDTH, SCREEN_HEIGHT)
FPS = 60

# Colors
BLACK = (0, 0, 0)
WHITE = (255, 255 ,255)
RED = (255, 0, 0)
BLUE = (0, 0, 255)
GREEN = (0, 255, 0)

H3 = pygame.font.Font(None, 28)

display_window = pygame.display.set_mode(SCREEN_SIZE)
clock = pygame.time.Clock()

while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            break


    display_window.fill(BLACK)
    



    f = str(clock.get_fps())
    fps_text = H3.render(f, True, WHITE)
    display_window.blit(fps_text, (0, 0))

    pygame.display.flip()
    clock.tick(FPS)
