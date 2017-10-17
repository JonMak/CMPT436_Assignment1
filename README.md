# CMPT436_Assignment1
Multi-client/Multi-chatroom chat server

A basic multi client, multiroom tcp chat server written in GO. Supports multiple room and multiple simultaneous connections.

The clients can:
  -create chat-rooms        ("/create")
  -list all existing rooms  ("/list")
  -join existing chat-rooms ("/join")
  -leave a chat-room        ("/leave")
  -send messages
  
 Upon joining a chat room, all previous messages from that chat room will be displayed. Joining and Leaving chat rooms are
 announced inside the chat rooms.

## Build

```
go build main.go
```

## Usage

* Starting the server:
```
/server --help
Usage of ./server:
  -ip string
    	IP address to listen on (default "127.0.0.1")
  -port string
    	Port to listen on (default "8000")

./server
```

* Connecting to the server:

```
telnet 127.0.0.1 8000
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
Please Enter Name: Jon
---------------------------
Welcome Jon
---------------------------
---------------------------
/leave : leave the current room
/quit : quit
/list : list all rooms
/create : create a new room
/join : join a room
/help : prints all available commands
---------------------------
```
