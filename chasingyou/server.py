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
import socket
from utils import pprint, sysexit
from threading import Thread
from player_sprite import PlayerSprite
import random
import concurrent

HOST = socket.gethostbyname(socket.gethostname())
PORT = 9999
ADDR = (HOST, PORT)
BUFFERSIZE = 1024
# available sprite pngs
SPRITE_LIST = ["red.png", "blue.png", "brown.png", "yellow.png", "pink.png", "dark.png", "white.png"]

game_running = True
players_store = []

def stop_listening():
    """A hack to stop socket from listening to connections when game is about to start"""

    socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect(ADDR)
    client_socket.close()

def gather_players():
    """Used by a thread to wait for functions without blocking"""
    
    # List of sprites that we remove from each time a sprite is allocated
    sprites = SPRITE_LIST

    while True:
        # On client connection
        global client_socket
        pprint("WAITING FOR PLAYERS TO CONNECT")
        client_socket, client_address = server_socket.accept()

        try:
            client_username = client_socket.recv(BUFFERSIZE)
        except socket.error as e:
            pprint("STOPPED WAITING FOR PLAYERS TO CONNECT")
            break

        print("PLAYER ACCEPTED, ADDRESS:", client_address, "USERNAME:", client_username)
        # TODO: Create player based on a class and the username given
        player = PlayerSprite(
            x = random.randint(100, 700), # can get the initial x and y values depending on the screensize (request size firstly)
            y = random.randint(100, 700),
            name=client_username,
            file_name= "assets/" + random.choice(sprites),
            alive = True,
            has_bomb=False
        )
        sprites = [sprite for sprite in sprites if sprite != player.file_name]
        players_store.append({"player": player, "socket": client_socket, "socketAddress": client_address})

def get_player_state(player_info: dict):
    s, player = player_info["socket"], player_info["player"]
    while game_running:
        print("-> Getting game state to:", player.name)
        data = s.recv(BUFFERSIZE).decode()
        print(data)

def send_game_state(player_info: dict):
    s, player = player_info["socket"], player_info["player"]
    while game_running:
        print("-> Sending game state to:", player.name)
        try:
            s.send("players state".encode())
        except OSError:
            print(f"Player {player.name} disconnected:")
            try:
                s.close()
            except Exception as e:
                print("Exception closing the socket:", e)
            break
    print("Removed the disconnected player")
    players_store.remove(player_info)


def main():

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
        return None

    pprint(f"STARTING THE GAME WITH {len(players_store)} PLAYERS")

    # Choose a random player to hold the bomb
    chosen_player = players_store[random.randint(0, len(players_store)-1)]
    chosen_player["player"].has_bomb = True
    print("Player chosen to hold the bomb:", chosen_player["player"].name)

    thrds_get, thrds_set = [], []
    for player in players_store:
        a = Thread(target=send_game_state, args=(player,))
        b = Thread(target=get_player_state, args=(player,))
        
        a.start()
        b.start()

        a.join()
        b.join()

    pprint("FINISHED THE GAME")

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        server_socket.close()
        pprint("EXITING THE GAME")
        sysexit()