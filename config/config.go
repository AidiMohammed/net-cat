package config

import (
	"os"
	"strconv"
	"net-cat/tools"
	"fmt"
)

func GetPort() uint {
	var portNumber int
	var err error
	if len(os.Args) == 2 {
		portNumber,err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Port %v is invalid the port 8989 is set by default\n",portNumber)
			return 8989
		}
		if portNumber < 0{
			fmt.Printf("Port %v is invalid the port 8989 is set by default\n",portNumber)
			return 8989
		} 
		return uint(portNumber)
	} else if len(os.Args) > 2 {
		message := tools.ColorString(tools.COLOR_RED,"[USAGE]: ./TCPChat $port") 
		fmt.Println(message)
		os.Exit(1)
	} 
	return 8989	
}