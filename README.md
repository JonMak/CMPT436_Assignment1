# Assignment 1
Multi-client/Multi-chatroom chat server

A basic multi client, multiroom tcp chat server written in GO. Supports multiple room and multiple simultaneous connections. Uses telnet to connect client to the server.

The clients can:

1. Create chat rooms       
2. List all existing rooms  
3. Join existing chat rooms 
4. Leave a chat room        
5. Send messages to other clients
  
 Upon joining a chat room, all previous messages from that chat room will be displayed. Joining and Leaving chat rooms are
 announced inside the chat rooms.

## Build

```
go build main.go
```

## Usage

* Starting the server:
```
/main --help
Usage of ./main:
  -ip string
    	IP address to listen on (default "127.0.0.1")
  -port string
    	Port to listen on (default "8000")

./main
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

# Testing

For testing I choose to do iterative testing throughout development. I wrote each piece and tested for compatibility. A final testing method I did was to check to see if 10 concurrent clients were able to complete the required tasks outlined in the assignment document.
Throughout the source code, logs are outputted to the console to keep track of everything the server does so that bugs are easy to define and fix.

#### Here is the server logs for testing 10 concurrent clients with sending/recieving messages, creating chat rooms, and leaving/joining chat rooms:
```
2017/10/17 16:12:19 listening on:  127.0.0.1:8000
2017/10/17 16:12:24 createclient: remote connection from: 127.0.0.1:49263
2017/10/17 16:12:26 createclient: remote connection from: 127.0.0.1:49265
2017/10/17 16:12:30 createclient: remote connection from: 127.0.0.1:49267
2017/10/17 16:12:33 createclient: remote connection from: 127.0.0.1:49269
2017/10/17 16:12:40 createclient: remote connection from: 127.0.0.1:49273
2017/10/17 16:12:43 createclient: remote connection from: 127.0.0.1:49275
2017/10/17 16:12:47 createclient: remote connection from: 127.0.0.1:49277
2017/10/17 16:12:50 createclient: remote connection from: 127.0.0.1:49279
2017/10/17 16:12:54 createclient: remote connection from: 127.0.0.1:49281
2017/10/17 16:12:58 createclient: remote connection from: 127.0.0.1:49283
2017/10/17 16:13:02 new client created: 127.0.0.1:49263 c1
2017/10/17 16:13:05 new client created: 127.0.0.1:49265 c2
2017/10/17 16:13:07 new client created: 127.0.0.1:49267 c3
2017/10/17 16:13:08 new client created: 127.0.0.1:49269 c4
2017/10/17 16:13:10 new client created: 127.0.0.1:49273 c5
2017/10/17 16:13:12 new client created: 127.0.0.1:49275 c6
2017/10/17 16:13:14 new client created: 127.0.0.1:49277 c7
2017/10/17 16:13:15 new client created: 127.0.0.1:49279 c8
2017/10/17 16:13:17 new client created: 127.0.0.1:49281 c9
2017/10/17 16:13:20 new client created: 127.0.0.1:49283 c10
2017/10/17 16:13:33 creating room r0
2017/10/17 16:13:33 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:13:33 receive: client(127.0.0.1:49263) recvd msg: * c1 has joined! * 
2017/10/17 16:13:40 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:13:40 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:13:40 receive: client(127.0.0.1:49265) recvd msg: * c2 has joined! * 
2017/10/17 16:13:40 receive: client(127.0.0.1:49263) recvd msg: * c2 has joined! * 
2017/10/17 16:13:52 creating room r1
2017/10/17 16:13:52 createroom: broadcasting msg in room: r1 to member: c3
2017/10/17 16:13:52 receive: client(127.0.0.1:49267) recvd msg: * c3 has joined! * 
2017/10/17 16:13:57 createroom: broadcasting msg in room: r1 to member: c3
2017/10/17 16:13:57 createroom: broadcasting msg in room: r1 to member: c4
2017/10/17 16:13:57 receive: client(127.0.0.1:49269) recvd msg: * c4 has joined! * 
2017/10/17 16:13:57 receive: client(127.0.0.1:49267) recvd msg: * c4 has joined! * 
2017/10/17 16:14:05 creating room r2
2017/10/17 16:14:05 createroom: broadcasting msg in room: r2 to member: c5
2017/10/17 16:14:05 receive: client(127.0.0.1:49273) recvd msg: * c5 has joined! * 
2017/10/17 16:14:13 createroom: broadcasting msg in room: r2 to member: c5
2017/10/17 16:14:13 createroom: broadcasting msg in room: r2 to member: c6
2017/10/17 16:14:13 receive: client(127.0.0.1:49273) recvd msg: * c6 has joined! * 
2017/10/17 16:14:13 receive: client(127.0.0.1:49275) recvd msg: * c6 has joined! * 
2017/10/17 16:14:23 creating room r3
2017/10/17 16:14:23 createroom: broadcasting msg in room: r3 to member: c7
2017/10/17 16:14:23 receive: client(127.0.0.1:49277) recvd msg: * c7 has joined! * 
2017/10/17 16:14:50 createroom: broadcasting msg in room: r3 to member: c7
2017/10/17 16:14:50 createroom: broadcasting msg in room: r3 to member: c8
2017/10/17 16:14:50 receive: client(127.0.0.1:49277) recvd msg: * c8 has joined! * 
2017/10/17 16:14:50 receive: client(127.0.0.1:49279) recvd msg: * c8 has joined! * 
2017/10/17 16:14:56 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:14:56 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:14:56 createroom: broadcasting msg in room: r0 to member: c9
2017/10/17 16:14:56 receive: client(127.0.0.1:49263) recvd msg: * c9 has joined! * 
2017/10/17 16:14:56 receive: client(127.0.0.1:49281) recvd msg: * c9 has joined! * 
2017/10/17 16:14:56 receive: client(127.0.0.1:49265) recvd msg: * c9 has joined! * 
2017/10/17 16:15:03 createroom: broadcasting msg in room: r0 to member: c10
2017/10/17 16:15:03 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:15:03 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:15:03 createroom: broadcasting msg in room: r0 to member: c9
2017/10/17 16:15:03 receive: client(127.0.0.1:49281) recvd msg: * c10 has joined! * 
2017/10/17 16:15:03 receive: client(127.0.0.1:49283) recvd msg: * c10 has joined! * 
2017/10/17 16:15:03 receive: client(127.0.0.1:49263) recvd msg: * c10 has joined! * 
2017/10/17 16:15:03 receive: client(127.0.0.1:49265) recvd msg: * c10 has joined! * 
2017/10/17 16:15:06 send: msg: Hello from: c1
2017/10/17 16:15:06 receive: client(127.0.0.1:49263) recvd msg: 17/10/2017 16:15:06 * (c1): "Hello" 
2017/10/17 16:15:06 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:15:06 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:15:06 createroom: broadcasting msg in room: r0 to member: c9
2017/10/17 16:15:06 createroom: broadcasting msg in room: r0 to member: c10
2017/10/17 16:15:06 receive: client(127.0.0.1:49283) recvd msg: 17/10/2017 16:15:06 * (c1): "Hello" 
2017/10/17 16:15:06 receive: client(127.0.0.1:49265) recvd msg: 17/10/2017 16:15:06 * (c1): "Hello" 
2017/10/17 16:15:06 receive: client(127.0.0.1:49281) recvd msg: 17/10/2017 16:15:06 * (c1): "Hello" 
2017/10/17 16:15:35 send: msg: Hello from: c4
2017/10/17 16:15:35 createroom: broadcasting msg in room: r1 to member: c4
2017/10/17 16:15:35 createroom: broadcasting msg in room: r1 to member: c3
2017/10/17 16:15:35 receive: client(127.0.0.1:49267) recvd msg: 17/10/2017 16:15:35 * (c4): "Hello" 
2017/10/17 16:15:35 receive: client(127.0.0.1:49269) recvd msg: 17/10/2017 16:15:35 * (c4): "Hello" 
2017/10/17 16:15:50 leave: removing user c8 from room r3: current members: map[127.0.0.1:49277:0xc4200d8100]
2017/10/17 16:15:50 createroom: broadcasting msg in room: r3 to member: c7
2017/10/17 16:15:50 receive: client(127.0.0.1:49277) recvd msg: * c8 has left.. * 
2017/10/17 16:15:59 leave: removing user c10 from room r0: current members: map[127.0.0.1:49263:0xc420060400 127.0.0.1:49265:0xc4200b6080 127.0.0.1:49281:0xc4200d8140]
2017/10/17 16:15:59 createroom: broadcasting msg in room: r0 to member: c1
2017/10/17 16:15:59 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:15:59 createroom: broadcasting msg in room: r0 to member: c9
2017/10/17 16:15:59 receive: client(127.0.0.1:49281) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c5
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c6
2017/10/17 16:15:59 receive: client(127.0.0.1:49265) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49283) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49273) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49263) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49275) recvd msg: * c10 has left.. * 
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c10
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c5
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c6
2017/10/17 16:15:59 createroom: broadcasting msg in room: r2 to member: c10
2017/10/17 16:15:59 receive: client(127.0.0.1:49273) recvd msg: * c10 has joined! * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49283) recvd msg: * c10 has joined! * 
2017/10/17 16:15:59 receive: client(127.0.0.1:49275) recvd msg: * c10 has joined! * 
2017/10/17 16:21:53 leave: removing user c1 from room r0: current members: map[127.0.0.1:49265:0xc4200b6080 127.0.0.1:49281:0xc4200d8140]
2017/10/17 16:21:53 c1 has left..
2017/10/17 16:21:53 createroom: broadcasting msg in room: r0 to member: c2
2017/10/17 16:21:53 createroom: broadcasting msg in room: r0 to member: c9
2017/10/17 16:21:53 receive: client(127.0.0.1:49281) recvd msg: * c1 has left.. * 
2017/10/17 16:21:53 receive: client(127.0.0.1:49265) recvd msg: * c1 has left.. * 
2017/10/17 16:21:56 c8 has left..
```


