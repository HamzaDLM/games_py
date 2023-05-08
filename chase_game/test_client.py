from networking import TCPhandler
from time import sleep
import asyncio

tcphandler = TCPhandler()
print("CLIENT")


async def sending(data):
    await tcphandler.send_data(data)


async def reading():
    data = await tcphandler.read_data()
    return data


async def main():
    data = input("-> ")
    while True:
        sleep(1)

        await sending(data)

        # sleep(0.1)

        # data = await reading()
        # print(data)


asyncio.run(main(), debug=True)
