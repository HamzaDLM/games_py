"""
Fluid simulation in PyGame.
Reference: https://mikeash.com/pyblog/fluid-simulation-for-dummies.html
"""
import random
import pygame

from fluid import Fluid, IX

pygame.init()

# Animation vars
N: int = 64
iter = 16
SCALE = 10

# Window vars
SCREEN_HEIGHT: int = N * SCALE
SCREEN_WIDTH: int = N * SCALE
SCREEN_SIZE = (SCREEN_WIDTH, SCREEN_HEIGHT)
FPS: int = 40

# Colors
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
GREY = (50, 50, 50)
RED = (255, 0, 0)
BLUE = (0, 0, 255)
GREEN = (0, 255, 0)

# Text fonts
H3 = pygame.font.Font(None, 28)

canvas = pygame.display.set_mode(SCREEN_SIZE)
clock = pygame.time.Clock()
pygame.display.set_caption("Fluid simulation")
dragging = False

# Initialize a fluid instance
fluid1 = Fluid(N=N, dt=0.1, diffussion=0, viscosity=0, iter=iter)


def constrain(value, minimum, maximum):
    if value < minimum:
        return minimum
    elif value > maximum:
        return maximum
    else:
        return value


def render_density(density: list[float]):
    for i in range(0, N):
        for j in range(0, N):
            x = i * SCALE
            y = j * SCALE
            d = density[IX(i, j)]
            d = int(constrain(d, 0, 255))
            color = (d, d, d)
            rect = pygame.Rect(x, y, SCALE, SCALE)
            pygame.draw.rect(canvas, color, rect)


while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            break
        canvas.fill(BLACK)
        if event.type == pygame.MOUSEBUTTONDOWN:
            if event.button == 1:
                dragging = True
        if event.type == pygame.MOUSEBUTTONUP:
            if event.button == 1:
                dragging = False
        if event.type == pygame.MOUSEMOTION:
            if dragging:
                mouse_x, mouse_y = event.pos
                if 0 < mouse_x < N * SCALE and 0 < mouse_y < N * SCALE:
                    cx = int(mouse_x / SCALE)
                    cy = int(mouse_y / SCALE)
                    fluid1.add_density(cx, cy, 100)
                    fluid1.add_velocity(cx, cy, 1, 1)

    fluid1.step()
    render_density(fluid1.density)
    fluid1.fade_density()

    f = str(clock.get_fps())
    fps_text = H3.render(f, True, WHITE)
    canvas.blit(fps_text, (0, 0))

    pygame.display.flip()
    clock.tick(FPS)
