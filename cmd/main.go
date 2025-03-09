package main

import (
	"fmt"
	"net"
	"os"
	"net-cat/tools"
	"net-cat/handel"
	"net-cat/config"
)

func main() {
	var portNumber uint
	var err error

	portNumber = config.GetPort()

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