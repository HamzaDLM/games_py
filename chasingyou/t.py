import json


class Smt:
    pass


class Something(Smt):
    def __init__(self, x, y):
        super().__init__()
        self.x = x
        self.y = y


l = []
l.append(Something(1, 2))
l.append(Something(1, 3))


data = [vars(s) for s in l]
for c in data:
    c.pop("x")


# print("Raw data:", data, type(data))

# data = json.dumps(data)
# print("Json dumps:", data, type(data))

# data = data.encode()
# print("Encoded:", data, type(data))
# ######## SENT VIA SOCKET

# ######## RECEIVED VIA SOCKET
# data = data.decode()
# print("Decoded:", data, type(data))

# data = json.loads(data)
# print("Json loads:", data, type(data))


# Smaller version

# Send
data = json.dumps(data).encode()
print(data)
# Receive
data = json.loads(data.decode())
print(data)
