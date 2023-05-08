import asyncio


class TCPhandler:
    # TCP handler for the clients
    def __init__(self, host: str = "localhost", port: int = 9000):
        self.host = host
        self.port = port

    async def send_data(self, data):
        """"""
        _, writer = await asyncio.open_connection(self.host, self.port)

        print("Send data", flush=True)
        writer.write(str(data).encode())
        await writer.drain()

        print("Close connection", flush=True)
        writer.close()
        await writer.wait_closed()

    async def read_data(self):
        reader, _ = await asyncio.open_connection(self.host, self.port)

        received = await reader.read(255)
        received = received.decode()
        print("Received data", received, flush=True)

        print("Close connection", flush=True)
        reader.close()


# Code for the server
async def handle_client(reader, writer):
    request = None
    while request != "quit":
        request = (await reader.read(255)).decode("utf8")
        response = str(request) + "\n"
        print(response)
        writer.write(response.encode("utf8"))
        await writer.drain()
    writer.close()


async def run_server():
    server = await asyncio.start_server(handle_client, "localhost", 9000)
    async with server:
        await server.serve_forever()


asyncio.run(run_server())
