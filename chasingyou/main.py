"""Main script for players"""
import pickle
import socket
from threading import Thread
from time import sleep
import asyncio

import pygame

from player_sprite import PlayerSprite
from utils import END, GREEN, ORANGE, pprint, sysexit

# pylint: disable=no-member

pygame.init()

# DECLARING USEFUL CONSTANTS
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
PLAYER_NAME = input("Provide your ingame name: ")
HOST = "localhost"  # input("Provide the host server IP: ")
PORT = 9999  # int(input("Provide the host server PORT: "))
ADDR = (HOST, PORT)
SCREEN_HEIGHT = 876  # pygame.display.Info().current_h * 0.9
SCREEN_WIDTH = 1542  # pygame.display.Info().current_w * 0.9
SPEED = 4
FPS = 40
BUFFERSIZE = 1024 * 10
SPRITE_POSITION = {
    "normal": (688, 102, 409, 506),
    "sliced": (712, 722, 409, 506),
    "dead": (1445, 731, 409, 506),
}
MUSIC = pygame.mixer.music.load("assets/wmd_ficus.mp3")
BACKGROUND_IMAGE = pygame.image.load("assets/maptest.gif")
TITLE = pygame.image.load("assets/title.png")

game_display = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
h2_text = pygame.font.Font(None, 40)
h3_text = pygame.font.Font(None, 28)
pygame.mouse.set_visible(False)
bomb = pygame.image.load("assets/bomb.png")
pygame.display.set_caption("Chasing You")  # Setting the window title


# CREATING PLAYER SPRITES
class Player(PlayerSprite):
    """Main player object"""

    def __init__(
        self,
        x: int,
        y: int,
        name: str,
        file_name: str,
        alive: bool = True,
        has_bomb: bool = False,
    ):
        super().__init__(
            x,
            y,
            name,
            file_name,
            alive,
            has_bomb,
        )

    def move(self, key: str):
        """Handle player movement"""
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
        """Used to assign bomb to a player"""
        self.has_bomb = not self.has_bomb

    def display_player(self, sp_x: int, sp_y: int, sp_w: int, sp_h: int):
        """sp_x, sp_y are related to the position of sprite in sheet"""
        # TODO: intialize those sp variables here depending on if player alive or dead
        sprite_image = pygame.image.load(self.file_name).convert_alpha()
        # self.rect = self.image.get_rect()
        render_name = h3_text.render(self.name, True, WHITE)
        game_display.blit(render_name, (self.x + 10, self.y - 25))
        sprite = pygame.Surface((sp_w, sp_h), pygame.SRCALPHA)
        sprite.blit(sprite_image, (0, 0), (sp_x, sp_y, sp_w, sp_h))
        scaled_sprite = pygame.transform.scale(sprite, (70, 70))
        scaled_sprite.set_colorkey((135, 132, 181))

        if self.has_bomb:
            scaled_bomb = pygame.transform.scale(bomb, (50, 50))
            scaled_sprite.blit(scaled_bomb, (30, 10), (0, 0, 50, 50))

        return scaled_sprite


# GAME LOGIC
def get_game_state():
    """Get data of other players from server
    If first time inserting players data, insert all regardless
    If not, then only update certain things
    """
    while True:
        sleep(1)
        print(GREEN, "-> Getting game state from", END, flush=True)
        global players_store
        payload = sock.recv(BUFFERSIZE)
        data = pickle.loads(payload)
        print("G", data, flush=True)
        # If players_store empty, then append players in it
        if not players_store:
            for row in data:
                players_store.append(
                    Player(
                        x=row["x"],
                        y=row["y"],
                        name=row["name"],
                        file_name=row["file_name"],
                        alive=row["alive"],
                        has_bomb=row["has_bomb"],
                    )
                )
        # If players already added to store, update them all except current player
        else:
            for row in data:
                for player in players_store:
                    if row["name"] == player.name and row["name"] != PLAYER_NAME:
                        player = Player(
                            x=row["x"],
                            y=row["y"],
                            name=row["name"],
                            file_name=row["file_name"],
                            alive=row["alive"],
                            has_bomb=row["has_bomb"],
                        )


def send_game_state():
    """Send movement data of the current player to server"""
    while True:
        sleep(1)
        print(ORANGE, "-> Sending game state to", END, flush=True)
        for player in players_store:
            if player.name == PLAYER_NAME:
                print("S", vars(player), flush=True)
                data = pickle.dumps(vars(player))
                sock.send(data)


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


def main():
    """Main game loop function"""
    pprint("STARTED MAIN LOOP")
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

        game_display.blit(BACKGROUND_IMAGE, (0, 0))
        game_display.blit(
            TITLE, (SCREEN_WIDTH / 2 - TITLE.get_size()[0] / 2, 40)
        )  # x, y

        # Update characters
        for player in players_store:
            if player.name == PLAYER_NAME:
                handle_movement(player)

        # Display players
        for player in players_store:
            game_display.blit(
                player.display_player(*SPRITE_POSITION["dead"]), (player.x, player.y)
            )

        pygame.display.flip()
        clock.tick(FPS)

    pygame.QUIT


if __name__ == "__main__":
    global sock, players_store
    players_store = []

    # open the socket connection
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    pprint("ATTEMPTING TO CONNECT TO SERVER")
    sock.connect(ADDR)
    pprint("CONNECTED SUCCESFULLY TO SERVER")

    # Send player username to server
    if PLAYER_NAME:
        sock.send(PLAYER_NAME.encode())
    else:
        PLAYER_NAME = input("Provide your ingame name: ")
        sock.send(PLAYER_NAME.encode())

    # Receive indication that game started
    started = sock.recv(BUFFERSIZE).decode()
    if started == "Game Started":
        pprint("GAME STARTED:", started)
    else:
        print("Error receiving game starting alert.")
        sysexit()

    pprint("START THREADS FOR EXCHANGING DATA WITH SERVER")
    t1 = Thread(target=get_game_state)
    t2 = Thread(target=send_game_state)

    t1.start()
    t2.start()

    main()

    t1.join()  # FIXME: this will never join, need to stop from inside
    t2.join()

    sock.close()
