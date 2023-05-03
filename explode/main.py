import pygame
import socket
from contextlib import contextmanager


pygame.init()

# DECLARING USEFUL CONSTANTS
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
# HOST = input("Provide the host server IP: ")
# PORT = int(input("Provide the host server PORT: "))
# ADDR = (HOST, PORT)
SCREEN_HEIGHT = pygame.display.Info().current_h * 0.9
SCREEN_WIDTH = pygame.display.Info().current_w * 0.9
FPS = 60
WIN = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
BUFFERSIZE = 1024
SPEED = 1

SPRITE_POSITION = {
    "normal": (688, 102, 409, 506),
    "sliced": (712, 722, 409, 506),
    "dead": (1445, 731, 409, 506)
}


h2_text = pygame.font.Font(None, 40)
h3_text = pygame.font.Font(None, 30)

game_started = False
running = True
clock = pygame.time.Clock()
pygame.mouse.set_visible(False)

# CREATING PLAYER SPRITES
class Player(pygame.sprite.Sprite):
    """
    x, y used in init are related to the position of the player in game
    sp_x, sp_y are related to the position of sprite in sheet
    """
    def __init__(self, x, y, name, filename):
        super().__init__()
        self.x = x
        self.y = y
        self.filename = filename
        self.image = pygame.image.load(filename).convert_alpha()
        self.rect = self.image.get_rect()
        self.name = h3_text.render(name, True, BLACK)
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
    
    def display_player(self, sp_x, sp_y, sp_w, sp_h):
        # WIN.blit(self.name, (self.x, self.y-40))
        sprite = pygame.Surface((sp_w, sp_h), pygame.SRCALPHA)
        # sprite.set_colorkey((135,132,181))
        sprite.blit(self.image, (0, 0), (sp_x, sp_y, sp_w, sp_h))
        scaled_sprite = pygame.transform.scale(sprite, (70, 70))
        scaled_sprite.set_colorkey((135,132,181))

        return scaled_sprite

# GAME LOGIC
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
    pass

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

# INITIALIZING OBJECTS
players = []
players.append(Player(100, 100, "Hamza", "img/blue.png"))
players.append(Player(400, 300, "Hamza", "img/red.png"))
players.append(Player(800, 300, "Hamza", "img/yellow.png"))
players.append(Player(300, 500, "Hamza", "img/brown.png"))


# MAIN GAME LOOP
while running:
 
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    WIN.fill(WHITE)

    # Check game status
    # get_game_status()
    
    # Update characters
    handle_movement()
    # WIN.blit(player1.scaled_img, (0, 0))
    
    for player in players:
        WIN.blit(
            player.display_player(*SPRITE_POSITION["normal"]), 
            (player.x, player.y)
        )

    # Send current status
    # send_game_status()



    pygame.display.flip()

    clock.tick(FPS)

pygame.QUIT
