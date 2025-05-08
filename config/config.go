package config

import (
	"os"
	"strconv"
	"net-cat/tools"
	"fmt"
	"errors"
)

func GetPort() (uint,error) {
	var portNumber int
	var err error
	if len(os.Args) == 2 {
		portNumber,err = strconv.Atoi(os.Args[1])
		if err != nil || portNumber < 1024 || portNumber > 49151 {
			fmt.Printf("Port %v is invalid the port 8989 is set by default\n",os.Args[1])
			return 8989,nil
		}
		return uint(portNumber),nil
	} else if len(os.Args) > 2 {
		return 0,errors.New(tools.ColorString(tools.COLOR_RED,"[USAGE]: ./TCPChat $port"))
	} 
	return 8989,nil	
}