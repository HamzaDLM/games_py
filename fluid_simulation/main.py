"""Fluid simulation in PyGame."""
import random

import pygame

from fluid import Fluid, IX, constrain

pygame.init()

# Generic vars
N: int = Fluid.N
iter = Fluid.iterations
SCALE = 4

# Window vars
SCREEN_HEIGHT: int = N * SCALE
SCREEN_WIDTH: int = N * SCALE
SCREEN_SIZE = (SCREEN_WIDTH, SCREEN_HEIGHT)
FPS: int = 120

# Colors
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
GREY = (50, 50, 50)
RED = (255, 0, 0)
BLUE = (0, 0, 255)
GREEN = (0, 255, 0)

# Text fonts
H3 = pygame.font.Font(None, 24)

canvas = pygame.display.set_mode(SCREEN_SIZE)
clock = pygame.time.Clock()
pygame.display.set_caption("Fluid simulation")
dragging = False

stats = f"""
Size: {N}
Iterations: {iter} 
"""
info = H3.render("Drag to change velocity/density.", True, WHITE)

# Initialize a fluid instance
fluid1 = Fluid(dt=0.2, diffussion=0, viscosity=0.0000001)


def multiline_text(t: str, x: int, y: int, font_size):
    t = t.strip().split("\n")
    offset = 0
    for line in t:
        text = H3.render(line, True, WHITE)
        canvas.blit(text, (x, y + offset))
        offset += font_size


def render_density(density: list[float]) -> None:
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
                    fluid1.add_density(cx, cy, random.randint(50, 150))
                    fluid1.add_velocity(cx, cy, 1, 1)

    fluid1.step()
    render_density(fluid1.density)
    fluid1.fade_density(amount=12)

    f = f"FPS: {round(clock.get_fps(), 1)}"
    fps_text = H3.render(f, True, WHITE)
    canvas.blit(fps_text, (10, 10))
    canvas.blit(info, (10, SCREEN_WIDTH - 28))

    multiline_text(stats, 10, 28, 20)

    pygame.display.flip()
    clock.tick(FPS)
