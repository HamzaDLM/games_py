import pygame
import socket
from contextlib import contextmanager

pygame.init()

# Colors
BLACK = (0, 0, 0)

HOST = input("Provide the host server IP: ")
PORT = input("Provide the host server PORT: ")
ADDR = (HOST, PORT)
SCREEN_HEIGHT = pygame.display.Info().current_h * 0.9
SCREEN_WIDTH = pygame.display.Info().current_w * 0.9
FPS = 60
WIN = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
BUFFERSIZE = 1024
SPEED = 1

game_started = False
running = True
clock = pygame.time.Clock()

class Player:
    def __init__(self, x, y, name):
        self.x = x
        self.y = y
        self.name = name
        self.has_bomb = False
    def move(self, key):
        if key == "UP":
            self.y += SPEED
        elif key == "DOWN":
            self.y -= SPEED
        elif key == "RIGHT":
            self.x += SPEED
        elif key == "LEFT":
            self.x -= SPEED
        else:
            pass
    def switch_bomb(self):
        self.has_bomb = not self.has_bomb

@contextmanager
def opensocket(*args, **kw):
    """Context manager for sockets"""
    s = socket.socket(*args, **kw)
    try:
        yield s
    finally:
        s.close()

def get_game_status():
    """Get data of other players from server
        Payload returned:
        {
            game_started: False,
            speed: 1:
            connected: True,
            players: {
                name: "XXXX",
                position: (100, 200)
                has_bomb: False
                }
       }
    """
    with opensocket(socket.AF_INET, socket.SOCK_STREAM) as s:
        try:
            s.bind(ADDR)
            print("UDP server up and listening")
            payload = s.recvfrom(BUFFERSIZE)
            print(payload)

        except socket.error as err:
            print("Problem with socket:", err)
            exit()

def send_game_status():
    """Send movement data to server
        Payload sent:
            {
                
            }
    """


def handle_movement():
    """Handle player movement"""
    # Track key strokes
    pressed_keys = pygame.key.get_pressed()
    if pressed_keys[pygame.K_UP]:
        pass 
    elif pressed_keys[pygame.K_DOWN]:
        pass 
    elif pressed_keys[pygame.K_RIGHT]:
        pass
    elif pressed_keys[pygame.K_LEFT]:
        pass

players = []

while running:
 
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    WIN.fill(BLACK)

    # Check game status
    get_game_status()
    
    # Update characters
    handle_movement()

    # Send current status
    send_game_status()




    pygame.display.flip()

    clock.tick(FPS)

pygame.QUIT
