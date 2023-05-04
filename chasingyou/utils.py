import os
import sys

# ANSI COLORS
MAIN = "\001\033[38;5;85m\002"
GREEN = "\001\033[38;5;82m\002"
GRAY = PLOAD = "\001\033[38;5;246m\002"
NAME = "\001\033[38;5;228m\002"
RED = "\001\033[1;31m\002"
FAIL = "\001\033[1;91m\002"
ORANGE = "\033[0;38;5;214m\002"
LRED = "\033[0;38;5;202m\002"
BOLD = "\001\033[1m\002"
UNDERLINE = "\001\033[4m\002"
END = "\001\033[0m\002"


def pprint(*args):
    print(RED, "======== ", *args, END, flush=True)


def sysexit():
    try:
        sys.exit(0)
    except SystemExit:
        os._exit(0)
