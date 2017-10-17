/*
	Jonathan Mak
	11157071
	jom262
	CMPT436 Assignment 1
	Oct. 17 2017

 */
package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Clients
type client struct {
	Conn    net.Conn
	Name    string
	Message chan string
	Room    string
}

// Chat rooms
type chatRoom struct {
	name     string
	messages chan string
	members  map[string]*client
	history  []string
}

// Keep track of rooms created.
var roomList = map[string]*chatRoom{}

// TCP server addresses and ports
var flagIP = flag.String("ip", "127.0.0.1", "IP address to listen on")
var flagPort = flag.String("port", "8181", "Port to listen on")

// Time layout to parse
var customTime = "02/01/2006 15:04:05"

// Server Commands
var help = map[string]string{
	"/quit":      "quit\n",
	"/list":   	  "list all rooms\n",
	"/create":    "create a new room\n",
	"/join":      "join a room\n",
	"/help":      "prints all available commands\n",
	"/leave":	  "leave the current room\n",
}

func main() {
	flag.Parse()

	//start TCP listener
	listener, err := net.Listen("tcp", *flagIP+":"+*flagPort)
	if err != nil {
		log.Fatalf("could not listen on interface %v:%v error: %v ", *flagIP, *flagPort, err)
	}
	defer listener.Close()
	log.Println("listening on: ", listener.Addr())

	//main listen accept loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("could not accept connection %v ", err)
		}
		//create new client on connection
		go createclient(conn)
	}
}

/**
	Creates a client upon connecting to the server
	Input:
		- conn: connection to the server.
 */
func createclient(conn net.Conn) {

	log.Printf("createclient: remote connection from: %v", conn.RemoteAddr())

	name, err := readInput(conn, "Please Enter Name: ")
	if err != nil {
		panic(err)
	}

	writeFormattedMsg(conn, "Welcome "+name)

	//initialize client struct
	client := &client{
		Message: make(chan string),
		Conn:    conn,
		Name:    name,
		Room:    "",
	}

	log.Printf("new client created: %v %v", client.Conn.RemoteAddr(), client.Name)

	//spin off separate send, receive
	go client.send()
	go client.receive()

	//print help
	writeFormattedMsg(conn, help)
}

/*
	Closes clients connection to the server
 */
func (c *client) close() {
	c.leave()
	c.Conn.Close()
	c.Message <- "/quit"
}

/*
	Allows clients to receive messages from the server.
 */
func (c *client) receive() {

	// Read message from the client received from the server
	for {
		msg := <-c.Message

		if msg == "/quit" {
			roomList[c.Room].announce(c.Name + " has left..")
			break
		}

		log.Printf("receive: client(%v) recvd msg: %s ", c.Conn.RemoteAddr(), msg)

		// Print out message to the current client.
		writeFormattedMsg(c.Conn, msg)
	}
}

/*
	Send Loop that checks the users input if it is a regular message or command.
 */
func (c *client) send() {
Loop:
	for {
		msg, err := readInput(c.Conn, "")
		if err != nil {
			panic(err)
		}

		if msg == "\\quit" {
			c.close()
			log.Printf("%v has left..", c.Name)
			break Loop
		}

		// Check if the message is a command.
		if c.command(msg) {
			// If the client is currently not connected to a room
			if c.Room == ""{
				log.Printf("Error: " + c.Name + " is not in a room")
				writeFormattedMsg(c.Conn, "You are currently not in a room, please join a room.")
			}else {
				// Client is connected to a room and the input is a regular message.
				log.Printf("send: msg: %v from: %s", msg, c.Name)
				send := time.Now().Format(customTime) + " * (" + c.Name + "): \"" + msg + "\""

				// Add message to the chat rooms history.
				roomList[c.Room].history = append(roomList[c.Room].history, send)

				// Send out message to all clients connected to the current room.
				for _, v := range roomList {
					for k := range v.members {
						if k == c.Conn.RemoteAddr().String() {
							v.messages <- send
						}
					}
				}
			}
		}
	}
}

/*
	Parse message to see if it is one of the server commands.
	Input:
		- msg: client input that will be parsed by the server.
	Returns: True if regular message, False if it is a server command.
 */
func (c *client) command(msg string) bool {
	switch {
	case msg == "/list":
		c.Conn.Write([]byte("-------------------\n"))
		for k := range roomList {
			count := 0
			for range roomList[k].members {
				count++
			}
			c.Conn.Write([]byte(k + " : online members(" + strconv.Itoa(count) + ")\n"))
		}
		c.Conn.Write([]byte("-------------------\n"))
		return false
	case msg == "/join":
		c.join()
		return false
	case msg == "/help":
		writeFormattedMsg(c.Conn, help)
		return false
	case msg == "/create":
		c.create()
		return false
	case msg == "/leave":
		c.leave()
		return false
	}
	return true
}

/*
	Joins client to inputted room.
 */
func (c *client) join() {

	// If there is no rooms, user cannot join a room.
	if len(roomList) == 0 {
		writeFormattedMsg(c.Conn, "There is currently no rooms created, " +
			"Please create a new room with the command \"/create")
	} else {
		roomName, err := readInput(c.Conn, "Please enter room name:\n")
		if err != nil {
			panic(err)
		}

		// User cannot join a room that it is currently part of.
		if c.Room == roomName {
			writeFormattedMsg(c.Conn, "You are already part in that channel")
		} else if cr := roomList[roomName]; cr != nil {

			//Create a new room
			cr.members[c.Conn.RemoteAddr().String()] = c

			// Leave client's current room to join new room
			if c.Room != "" {
				c.leave()
				cr.announce(c.Name + " has left..")
			}

			// set client's room to new room
			c.Room = roomName
			writeFormattedMsg(c.Conn, "========= " + cr.name + " =========")
			writeFormattedMsg(c.Conn, c.Name+" has joined "+cr.name)

			// Display all previous messages in the chat room when joining new room
			for _,m := range roomList[c.Room].history {
				writeFormattedMsg(c.Conn, m)
			}

			// Announce to all members of room of a new client joining.
			cr.announce(c.Name + " has joined!")

		} else {
			writeFormattedMsg(c.Conn, "error: could not join room")
		}
	}
}

/*
	leave the current room
 */
func (c *client) leave() {

	//only if room is not empty, cannot leave a room if not part of one.
	if c.Room != "" {
		delete(roomList[c.Room].members, c.Conn.RemoteAddr().String())
		log.Printf("leave: removing user %v from room %v: current members: %v", c.Name, c.Room, roomList[c.Room].members)
		writeFormattedMsg(c.Conn, "leaving " + c.Room)
		roomList[c.Room].announce(c.Name + " has left..")
		c.Room = ""

	} else {
		writeFormattedMsg(c.Conn, "* error: You are currently not in a room")
		writeFormattedMsg(c.Conn, help)
	}
}

/*
	Creates a new room.
 */
func (c *client) create() {
	roomName, err := readInput(c.Conn, "Please enter room name: ")
	if err != nil {
		panic(err)
	}

	// Cannot make a room that already exists.
	if _, ok := roomList[roomName];ok {
		writeFormattedMsg(c.Conn,"A room already exists with that name, try again later")
	} else if roomName != "" {
		cr := createRoom(roomName)
		cr.members[c.Conn.RemoteAddr().String()] = c

		// leave current room to join newly create room
		if c.Room != "" {
			roomList[c.Room].announce(c.Name + " has left..")
			c.leave()
			}

		// set clients room to new room
		c.Room = cr.name

		// add new room to map
		roomList[cr.name] = cr
		cr.announce(c.Name + " has joined!")
		writeFormattedMsg(c.Conn, "* room "+cr.name+" has been created *")
		writeFormattedMsg(c.Conn, "============ " + cr.name + " ============")

		} else {
			writeFormattedMsg(c.Conn, "* error: could not create room \""+roomName+"\" *")
			}

}

/*
	Prints message in formatted style.
	Input:
		- conn: clients connection to server
		- msg: message to be formatted and printed to other clients.

 */
func writeFormattedMsg(conn net.Conn, msg interface{}) error {
	_, err := conn.Write([]byte("---------------------------\n"))
	t := reflect.ValueOf(msg)
	switch t.Kind() {
	case reflect.Map:
		for k, v := range msg.(map[string]string) {
			_, err = conn.Write([]byte(k + " : " + v))
		}
		break
	case reflect.String:
		v := reflect.ValueOf(msg).String()
		_, err = conn.Write([]byte(v + "\n"))
		break
	} //switch
	conn.Write([]byte("---------------------------\n"))

	if err != nil {
		return err
	}
	return nil
}

/*
	Reads Input so that server can parse message.
	Input:
		- conn: clients connection to server
		- qst:  question to ask client.
 */
func readInput(conn net.Conn, qst string) (string, error) {
	conn.Write([]byte(qst))
	s, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("readinput: could not read input from stdin: %v from client %v", err, conn.RemoteAddr().String())
		return "", err
	}
	s = strings.Trim(s, "\r\n")
	return s, nil
}

/*
	Constructs a new chatRoom
	Input: name = name of the chatRoom to be created.
	Returns: a new chatRoom
 */
func createRoom(name string) *chatRoom {
	// Initialize chat room struct
	c := &chatRoom{
		name:     name,
		messages: make(chan string),
		members:  make(map[string]*client, 0),
		history:  []string{},
	}
	log.Printf("creating room %v", c.name)

	//spin off new routine to listen for messages
	go func(c *chatRoom) {
		for {
			out := <-c.messages
			for _, v := range c.members {
				v.Message <- out
				log.Printf("createroom: broadcasting msg in room: %v to member: %v", c.name, v.Name)
			}
		}
	}(c)

	return c
}

/*
	Send an announcement to all members of the chatRoom.
 */
func (c *chatRoom) announce(msg string) {
	c.messages <- "* " + msg + " *"
}