import os 
import sys


def pprint(*args):
    print("======== ", *args ," ========")

def sysexit():
    try:
        sys.exit(0)
    except SystemExit:
        os._exit(0)
