"""
Server starts
Wait for players to join
Input asks to start the game (yes)
Start threads for each client and 

"""
import socket
import sys
import os
from threading import Thread

HOST = socket.gethostname(socket.gethostname())
PORT = 9999
ADDR = (HOST, PORT)

players_info = {

}

def pprint(*args):
    print("======== ", *args ," ========")

def sysexit():
    try:
        sys.exit(0)
    except SystemExit:
        os._exit(0)

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

    while True:
        # On client connection 
        client_socket, client_address = server_socket.accept()
        print("CLIENT ACCEPTED, address:", client_address)
        # client_thread = Thread(target=handle_client, args=(client_name))
        # client_thread.start()

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        pprint("EXITING THE GAME")
        sysexit()