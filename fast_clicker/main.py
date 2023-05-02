import pygame
import time
import random

# Initialize the game
pygame.init()

# Define constants
RED = (255, 0, 0)
GREEN = (0, 255, 51)
BLUE = (0, 0, 255)
BLACK = (0, 0, 0)
DARK_BLUE = (0, 0, 100)
YELLOW = (255, 255, 0)
LIGHT_BLUE = (80, 80, 255)
CYON = (65, 223, 220)
FPS = 40
main_font = pygame.font.Font(None, 40)
secondary_font = pygame.font.Font(None, 30)

# Initialize pygame
screen_height = 500
screen_width = 500

mw = pygame.display.set_mode((screen_height, screen_width)) 
running = True
clock = pygame.time.Clock()

# Making the yellow cards
class Area:
    def __init__(self, backgound_color, x, y, height, width):
        self.backgound_color = backgound_color
        self.rectangle = pygame.Rect(x, y, height, width)
    def draw_rectangle(self, bg_color: tuple = None):
        if bg_color == None:
            bg_color = self.backgound_color
        pygame.draw.rect(mw, bg_color, self.rectangle)
    def is_clicked(self):
        return pygame.mouse.get_pressed()[0] and self.rectangle.collidepoint(pygame.mouse.get_pos())

# Make the Label class (text)
class Label(Area):
    def __init__(self, text, x, y):
        self.label = secondary_font.render(text, True, BLACK)
        self.x = x
        self.y = y
    def write_text(self):
        mw.blit(self.label, (self.x, self.y))

class Timer():
    def __init__(self, x, y):
        self.label = main_font.render("Timer:", True, BLACK)
        self.x = x
        self.y = y
        self.timer = 0
    def write_text(self):
        mw.blit(self.label, (self.x, self.y))
        timer_text = main_font.render(str(self.timer), True, BLACK)
        mw.blit(timer_text, (self.x*2, self.y*5))
    def increment(self):
        self.timer += 1
    

class Points:
    def __init__(self, x, y):
        self.label = main_font.render("Points:", True, BLACK)
        self.x = x
        self.y = y
        self.counter = 0
    def write_text(self):
        mw.blit(self.label, (self.x, self.y))
        score_text = main_font.render(str(self.counter), True, BLACK)
        mw.blit(score_text, (self.x+30, self.y*5))
    def increment(self):
        self.counter += 1
    def decremennt(self):
        self.counter -= 1

# Creat the rectangles and the texts from the classes above
areas_list = []
labels_list = []
w, h = 100, 150
x, y, = 25, screen_height/2 - h/2
for i in range(4):
    # Create areas
    area = Area(YELLOW, x, y, w, h)
    areas_list.append(area)
    # Create labels 
    text = "CLICK"
    x1 = x + (area.rectangle.width - secondary_font.size(text)[0]) / 2 
    y1 = screen_height / 2 -  secondary_font.size(text)[1]/ 2
    label = Label(text, x1, y1)
    labels_list.append(label)
    x += 120

points = Points(screen_width-30-main_font.size("Points:")[0], 10)
timer = Timer(20, 10)

def randchoice(length: int, current: int | None = None) -> int:
    choices = [i for i in range(length) if i != current]
    return random.choice(choices)

first_tick = True
randomcounter = 0
choice = 0
# TODO: add disable showing text for small time after (improvement)
show = True

# The main game Loop
while running:

    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    mw.fill(CYON)

    # Refresh choice after half sec
    change_time = int(FPS/1.5)
    if first_tick or randomcounter > change_time:
        choice = randchoice(len(labels_list), choice)
        randomcounter = 0
        first_tick = False
    randomcounter += 1

    # Draw the rectangles
    for i in range(len(areas_list)):
        if areas_list[i].is_clicked():
            if choice == i:
                points.increment()
                areas_list[i].draw_rectangle(bg_color=GREEN)
            else:
                points.decremennt()
                areas_list[i].draw_rectangle(bg_color=RED)
        else:
            areas_list[i].draw_rectangle()

    # Draw the labels
    for i in range(len(labels_list)):
        if i == choice:
            labels_list[i].write_text()

    timer.write_text()
    points.write_text()

    







    pygame.display.flip()

    clock.tick(FPS)

pygame.QUIT
