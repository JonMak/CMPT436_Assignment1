# Assignment 1
Multi-client/Multi-chatroom chat server

A basic multi client, multiroom tcp chat server written in GO. Supports multiple room and multiple simultaneous connections.

The clients can:

1. Create chat rooms       
2. List all existing rooms  
3. Join existing chat rooms 
4. Leave a chat room        
5. Send messages
  
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

* Creating a chat room:

```
/leave : leave the current room
/quit : quit
/list : list all rooms
/create : create a new room
/join : join a room
/help : prints all available commands
---------------------------
/create
Please enter room name: Test
---------------------------
* Jon has joined! *
---------------------------
* room Test has been created *
---------------------------
---------------------------
============ Test ============
---------------------------
---------------------------
```

* Joining a chat room:

```
/join
Please enter room name:
Test
---------------------------
========= Test =========
---------------------------
---------------------------
Jon has joined Test
---------------------------
---------------------------
* Jon has joined! *
---------------------------

```

* Leaving a chat room:

```
---------------------------
============ Test ============
---------------------------
---------------------------
/leave
---------------------------
leaving Test
---------------------------
```

* List all existing chat rooms:

```
/list
-------------------
Test : online members(0)
Test2 : online members(0)
Test3 : online members(1)
-------------------

```

* Sending a message to a chat room:

```
Hello
---------------------------
17/10/2017 15:22:05 * (Jon): "Hello"
---------------------------
```

* Quiting the program:

```
/quit
---------------------------
leaving Test3
---------------------------
```



