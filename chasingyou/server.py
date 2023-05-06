"""
Game flow:
    Server starts
    Wait for players to join
    Input asks to start the game (yes)
    Start threads for each client
    send game state
    receive player info
    update the game state
    repeat
    if 1 player == is_alive
    announce winner and stop game 
"""
import random
import socket
import pickle
from time import sleep
from threading import Thread

from player_sprite import PlayerSprite
from utils import pprint, sysexit, END, ORANGE, GREEN

# pylint: disable=no-member

HOST = "192.168.1.5"  # socket.gethostbyname(socket.gethostname())
PORT = 9999
ADDR = (HOST, PORT)
BUFFERSIZE = 1024
# available sprite pngs
SPRITE_LIST = [
    "red.png",
    "blue.png",
    "brown.png",
    "yellow.png",
    "pink.png",
    "dark.png",
    "white.png",
]


def stop_listening():
    """A hack to stop socket from listening to connections when game is about to start"""

    socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect(ADDR)
    client_socket.close()


def gather_players():
    """Used by a thread to wait for functions without blocking"""

    global game_running, players_store

    # List of sprites that we remove from each time a sprite is allocated
    sprites_available = SPRITE_LIST

    while not game_running:
        # On client connection
        global client_socket
        pprint("WAITING FOR PLAYERS TO CONNECT")

        client_socket, client_address = server_socket.accept()

        client_username = client_socket.recv(BUFFERSIZE).decode()

        if not client_username:
            pprint("STOPPED WAITING FOR PLAYERS TO CONNECT")
            game_running = True
            break

        print("PLAYER ACCEPTED, ADDRESS:", client_address, "USERNAME:", client_username)

        player = PlayerSprite(
            # can get the initial x and y values depending on the screensize (request size firstly)
            x=random.randint(100, 700),
            y=random.randint(100, 700),
            name=client_username,
            file_name="assets/" + random.choice(sprites_available),
            alive=True,
            has_bomb=False,
        )

        sprites_available = [
            sprite for sprite in sprites_available if sprite != player.file_name
        ]
        players_store.append(
            {"player": player, "socket": client_socket, "socketAddress": client_address}
        )


def get_player_state(player_info: dict):
    """Receiving player state"""

    player_socket, player_obj = player_info["socket"], player_info["player"]

    while game_running:
        sleep(1)
        print(GREEN, "-> Getting game state from:", player_obj.name, END, flush=True)

        try:
            data = player_socket.recv(BUFFERSIZE)
            data = pickle.loads(data)
            print(data, flush=True)
            if data == "":
                print("Received empty payload when getting game state", flush=True)
                break
        except Exception as err:  # User might have disconnected
            print("Exception when getting game state:", err, flush=True)
            break

        for player in players_store:
            if player["player"].name == data["name"]:
                player["player"] = PlayerSprite(
                    # can get the initial x and y values depending on the screensize (request size firstly)
                    x=data["x"],
                    y=data["y"],
                    name=data["name"],
                    file_name=data["file_name"],
                    alive=data["alive"],
                    has_bomb=data["has_bomb"],
                )


def send_game_state(player_info: dict):
    """Sending game state to players"""

    player_socket, player = player_info["socket"], player_info["player"]

    while game_running:
        sleep(1)
        print(ORANGE, "-> Sending game state to:", player.name, END, flush=True)

        data = [vars(player["player"]) for player in players_store]
        for row in data:
            if "_Sprite__g" in row.keys():
                row.pop("_Sprite__g")
        print(data, flush=True)
        data = pickle.dumps(data)
        player_socket.send(data)

    print("Removed the disconnected player")
    players_store.remove(player_info)


def main() -> None:
    """Main server function"""
    pprint("GAME INITIALIZED")

    try:
        global server_socket
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.bind(ADDR)
    except socket.error as err:
        print("Error with socket:", err)
        sysexit()

    server_socket.listen(20)
    gather = Thread(target=gather_players)
    gather.start()

    _ = input("Type anything to start the game: ")
    _ = input("Are you sure you wanna start the game? : ")

    stop_listening()

    if not players_store:
        pprint("NO PLAYERS JOINED :( TERMINATING THE GAME")
        return

    pprint(f"STARTING THE GAME WITH {len(players_store)} PLAYERS")

    # Choose a random player to hold the bomb
    chosen_player = players_store[random.randint(0, len(players_store) - 1)]
    chosen_player["player"].has_bomb = True
    print("Player chosen to hold the bomb:", chosen_player["player"].name)

    for player in players_store:
        # Alert the players that the game started
        player["socket"].send("Game Started".encode())

        thread1 = Thread(target=send_game_state, args=(player,))
        thread2 = Thread(target=get_player_state, args=(player,))

        thread1.start()
        thread2.start()

        thread1.join()
        thread2.join()

    pprint("FINISHED THE GAME")


if __name__ == "__main__":
    global game_running, players_store

    game_running = False
    players_store = []

    try:
        main()
    except KeyboardInterrupt:
        server_socket.close()
        pprint("EXITING THE GAME")
        sysexit()
    server_socket.close()
