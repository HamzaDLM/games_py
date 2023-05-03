import pygame
import socket
from contextlib import contextmanager


pygame.init()

# DECLARING USEFUL CONSTANTS
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
HOST = input("Provide the host server IP: ")
PORT = int(input("Provide the host server PORT: "))
ADDR = (HOST, PORT)
SCREEN_HEIGHT = 876 #pygame.display.Info().current_h * 0.9
SCREEN_WIDTH = 1542 #pygame.display.Info().current_w * 0.9
FPS = 40
BUFFERSIZE = 1024
SPEED = 4
SPRITE_POSITION = {
    "normal": (688, 102, 409, 506),
    "sliced": (712, 722, 409, 506),
    "dead": (1445, 731, 409, 506)
}
MUSIC = pygame.mixer.music.load("assets/wmd_ficus.mp3")
BACKGROUND_IMAGE = pygame.image.load("assets/maptest.gif")
TITLE = pygame.image.load("assets/title.png")

game_display = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
h2_text = pygame.font.Font(None, 40)
h3_text = pygame.font.Font(None, 28)
pygame.mouse.set_visible(False)
bomb = pygame.image.load("assets/bomb.png")

# CREATING PLAYER SPRITES
class Player(pygame.sprite.Sprite):
    """
    x, y used in init are related to the position of the player in game
    sp_x, sp_y are related to the position of sprite in sheet
    """
    def __init__(self, x: int, y: int, name: str, file_name: str, current_player: bool = False, alive: bool = True, has_bomb: bool = False):
        super().__init__()
        self.x = x
        self.y = y
        self.file_name = file_name
        self.image = pygame.image.load(file_name).convert_alpha()
        self.rect = self.image.get_rect()
        self.name = name
        self.has_bomb = has_bomb
        self.current_player = current_player
        self.alive = alive

    def move(self, key: str):
        if key == "UPRIGHT":
            self.x += SPEED
            self.y -= SPEED
        elif key == "UPLEFT":
            self.x -= SPEED
            self.y -= SPEED
        elif key == "DOWNRIGHT":
            self.x += SPEED
            self.y += SPEED
        elif key == "DOWNLEFT":
            self.x -= SPEED
            self.y += SPEED
        elif key == "UP":
            self.y -= SPEED
        elif key == "DOWN":
            self.y += SPEED
        elif key == "RIGHT":
            self.x += SPEED
        elif key == "LEFT":
            self.x -= SPEED
        else:
            pass

    def switch_bomb(self):
        self.has_bomb = not self.has_bomb
    
    def display_player(self, sp_x: int, sp_y: int, sp_w: int, sp_h: int):
        render_name = h3_text.render(self.name, True, WHITE)
        game_display.blit(render_name, (self.x+10, self.y-25))
        sprite = pygame.Surface((sp_w, sp_h), pygame.SRCALPHA)
        sprite.blit(self.image, (0, 0), (sp_x, sp_y, sp_w, sp_h))
        scaled_sprite = pygame.transform.scale(sprite, (70, 70))
        scaled_sprite.set_colorkey((135,132,181))
        
        if self.has_bomb:
            scaled_bomb = pygame.transform.scale(bomb, (50, 50))
            scaled_sprite.blit(scaled_bomb, (30, 10), (0, 0, 50, 50))

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

def handle_movement(player: Player):
    """Handle player movement"""
    # Track key strokes
    pressed_keys = pygame.key.get_pressed()
    if pressed_keys[pygame.K_UP] and pressed_keys[pygame.K_RIGHT]:
        player.move("UPRIGHT")
    elif pressed_keys[pygame.K_UP] and pressed_keys[pygame.K_LEFT]:
        player.move("UPLEFT")
    elif pressed_keys[pygame.K_DOWN] and pressed_keys[pygame.K_RIGHT]:
        player.move("DOWNRIGHT")
    elif pressed_keys[pygame.K_DOWN] and pressed_keys[pygame.K_LEFT]:
        player.move("DOWNLEFT")
    elif pressed_keys[pygame.K_UP]:
        player.move("UP")
    elif pressed_keys[pygame.K_DOWN]:
        player.move("DOWN")
    elif pressed_keys[pygame.K_RIGHT]:
        player.move("RIGHT")
    elif pressed_keys[pygame.K_LEFT]:
        player.move("LEFT")
    else:
        pass

# INITIALIZING OBJECTS & GAMEPLAY-RELAED VARIABLES
players = []
players.append(Player(x=100, y=100, name="Blue", file_name="assets/blue.png", current_player=True, has_bomb=True))
players.append(Player(x=400, y=300, name="Red", file_name="assets/red.png"))
players.append(Player(x=800, y=300, name="Yellow", file_name="assets/yellow.png"))
players.append(Player(x=300, y=500, name="Brown", file_name="assets/brown.png"))

tick_counter = pygame.time.get_ticks()
pygame.mixer.music.play()
alpha = 255
game_started = False
running = True
clock = pygame.time.Clock()

def main():
    # MAIN GAME LOOP
    while running:
    
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        # game_display.fill(WHITE)
        game_display.blit(BACKGROUND_IMAGE, (0, 0))
        game_display.blit(TITLE, (SCREEN_WIDTH/2 - TITLE.get_size()[0]/2, 40)) # x, y

        # Check game status
        # get_game_status()
        
        # Update characters
        handle_movement(players[0])
        # game_display.blit(player1.scaled_assets, (0, 0))
        
        for player in players:
            game_display.blit(
                player.display_player(*SPRITE_POSITION["normal"]), 
                (player.x, player.y)
            )

        # build payload to send to server
        payload = {
            player.x,
            player.y,
            player.name,
            player.alive
        }
        # Send current status
        # send_game_status()

        pygame.display.flip()

        clock.tick(FPS)

    pygame.QUIT


if __name__ == "__main__":
    # open the socket connection

    main()