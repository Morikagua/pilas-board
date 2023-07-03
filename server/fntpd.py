#!/usr/bin/env python3

import socket as s
import sys
import os


def listen():
    socket = s.socket()
    socket.bind(("0.0.0.0", int(sys.argv[1])))
    socket.listen(0)
    while True:
        client, addr = socket.accept()
        response = client.recv(3333)
        if response == b'GET_TIME':
            time  = os.popen("cmd /c time /t").read()
            date  = os.popen("cmd /c date /t").read()
            client.send((time + date).encode())
            


if __name__ == "__main__":
    try:
        listen()
    except:
        pass
