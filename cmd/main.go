package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"net-cat/tools"
	"net-cat/handel"
)
//
func main() {
	var portNumber int
	var err error

	if len(os.Args) == 2 {
		portNumber,err = strconv.Atoi(os.Args[1])
		if err != nil {
			message := tools.ColorString(tools.COLOR_RED,fmt.Sprintf("Port number (%v) is invalid",os.Args[1]))
			fmt.Println(message)
			portNumber,_ = strconv.Atoi("8989")
			message = tools.ColorString(tools.COLOR_GREEN,"The port 8989 is defined by default")
			fmt.Println(message)
		}
	} else if len(os.Args) > 2  {
		message := tools.ColorString(tools.COLOR_RED,"[USAGE]: ./TCPChat $port") 
		fmt.Println(message)
		os.Exit(1)
	} else if len(os.Args) == 1 {
		portNumber = 8989
	}

	listener, err := net.Listen("tcp",fmt.Sprintf("localhost:%v",portNumber)) 
	defer listener.Close()
	if err != nil {
		message := tools.ColorString(tools.COLOR_RED,"Erreur de démarrage de serveur : ")
		fmt.Println(message,err)
		os.Exit(1)
	}
	
	fmt.Println(fmt.Sprintf("Listening on the port : %v",portNumber))
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur d'acceptation de connxion : ",err)
			continue
		}
		go handel.ConnectionManagement(conn)
	}
}