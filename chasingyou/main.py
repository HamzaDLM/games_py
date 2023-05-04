import pygame
import socket
from utils import sysexit, pprint
from player_sprite import PlayerSprite

pygame.init()

# DECLARING USEFUL CONSTANTS
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
PLAYER_NAME = input("Provide your ingame name: ")
HOST = socket.gethostbyname(socket.gethostname()) #input("Provide the host server IP: ")
PORT = 9999 #int(input("Provide the host server PORT: "))
ADDR = (HOST, PORT)
SCREEN_HEIGHT = 876 #pygame.display.Info().current_h * 0.9
SCREEN_WIDTH = 1542 #pygame.display.Info().current_w * 0.9
SPEED = 4
FPS = 40
BUFFERSIZE = 1024
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
pygame.display.set_caption("Chasing You") # Setting the window title

# CREATING PLAYER SPRITES
class Player(PlayerSprite):

    def __init__(self):
        super().__init__()
    
    def move(self, key: str):
        if self.alive:
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
    
    def display_player(self, sp_x: int, sp_y: int, sp_w: int, sp_h: int, game_display):
        """sp_x, sp_y are related to the position of sprite in sheet"""
        # TODO: intialize those sp variables here depending on if player alive or dead 
        sprite_image = pygame.image.load(self.file_name).convert_alpha()
        # self.rect = self.image.get_rect()
        render_name = h3_text.render(self.name, True, WHITE)
        game_display.blit(render_name, (self.x+10, self.y-25))
        sprite = pygame.Surface((sp_w, sp_h), pygame.SRCALPHA)
        sprite.blit(sprite_image, (0, 0), (sp_x, sp_y, sp_w, sp_h))
        scaled_sprite = pygame.transform.scale(sprite, (70, 70))
        scaled_sprite.set_colorkey((135,132,181))
        
        if self.has_bomb:
            scaled_bomb = pygame.transform.scale(bomb, (50, 50))
            scaled_sprite.blit(scaled_bomb, (30, 10), (0, 0, 50, 50))

        return scaled_sprite


# GAME LOGIC
def get_game_state():
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
    payload = sock.recv(BUFFERSIZE).decode()
    print(payload)

def send_game_state():
    """Send movement data to server
        Payload sent:
            {
                
            }
    """
    data = "something"
    sock.send(data.encode())

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

# INITIALIZING OBJECTS & GAMEPLAY-RELATED VARIABLES
players = []
# players.append(Player(x=100, y=100, name="Blue", file_name="assets/blue.png", has_bomb=True))
# players.append(Player(x=400, y=300, name="Red", file_name="assets/red.png"))
# players.append(Player(x=800, y=300, name="Yellow", file_name="assets/yellow.png"))
# players.append(Player(x=300, y=500, name="Brown", file_name="assets/brown.png"))
def init():
    """Initialize the first state of the game (creating players...)"""
    pass

def main():
    tick_counter = pygame.time.get_ticks()
    pygame.mixer.music.play(-1)
    alpha = 255
    game_started = False
    running = True
    clock = pygame.time.Clock()
    
    # MAIN GAME LOOP
    while running:
    
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        # game_display.fill(WHITE)
        game_display.blit(BACKGROUND_IMAGE, (0, 0))
        game_display.blit(TITLE, (SCREEN_WIDTH/2 - TITLE.get_size()[0]/2, 40)) # x, y

        # Check game status
        get_game_state()
        
        # Update characters
        # handle_movement(players[0])
        # game_display.blit(player1.scaled_assets, (0, 0))
        
        # for player in players:
        #     game_display.blit(
        #         player.display_player(*SPRITE_POSITION["normal"]), 
        #         (player.x, player.y)
        #     )
        print("- DOING SOMETHING")
        # build payload to send to server
        # payload = {
        #     player.x,
        #     player.y,
        #     player.name,
        #     player.alive
        # }
        # Send current status
        send_game_state()

        pygame.display.flip()

        clock.tick(FPS)

    pygame.QUIT


if __name__ == "__main__":
    # open the socket connection
    try:
        global sock
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        pprint("ATTEMPTING TO CONNECT TO SERVER")
        sock.connect(ADDR)
        pprint("CONNECTED SUCCESFULLY TO SERVER")
        if PLAYER_NAME:
            sock.send(PLAYER_NAME.encode())
        else:
            PLAYER_NAME = input("Provide your ingame name: ")
            sock.send(PLAYER_NAME.encode())
        pprint("STARTING MAIN GAME LOOP")
        init()
        main()
    except Exception as e:
        pprint("PROBLEM EXECUTING THE PROGRAM:", e)
        sysexit
