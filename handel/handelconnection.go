package handel

import (
	"sync"
	"fmt"
	"net"
	"net-cat/tools"
	"bufio"
	"strings"
	"time"
	//"os"
)

var (
	users = make(map[net.Conn]string)
	historiqueMessage []string
	mutex sync.Mutex
)

func ConnectionManagement(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connecté : ", conn.RemoteAddr())
	messageWelcom,err := tools.WelcomMessage()
	if err != nil {
		messageErr := tools.ColorString(tools.COLOR_RED,fmt.Sprintf("Il y a eu une erreur lors de l'envoi du message welcom [%v]: %v",err,conn.RemoteAddr()))
		fmt.Println(messageErr)
		return
	}
	conn.Write([]byte(messageWelcom))

	reader := bufio.NewReader(conn)
	name,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture : ",err)
		return
	}

	name = strings.TrimSpace(name)

	mutex.Lock()
	users[conn] = name
	fmt.Printf("Add new user %v address %v \n",name,conn.RemoteAddr())

	for valueConn,valueName := range users {
		if valueName != name {
			valueConn.Write([]byte(fmt.Sprintf("%v has joined our chat...\n",name)))
		}
	}
	mutex.Unlock()

	mutex.Lock()
	for _,histo := range historiqueMessage {
		conn.Write([]byte(histo))
	}
	mutex.Unlock()

	for {
		now := time.Now()
		message,err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf(fmt.Sprintf("User %s déconnecté.\n",name))
			mutex.Lock()
			delete(users,conn)
			fmt.Printf(fmt.Sprintf("delete user : %v\n",name))
			mutex.Unlock()
			return
		}

		message = strings.TrimSpace(message)

		if len(message) > 0 {
			messageToUser := fmt.Sprintf("[%v][%v]:%v\n",now.Format("2006-01-02 15:04:05"),name,message)
			conn.Write([]byte(messageToUser))	
			mutex.Lock()
			historiqueMessage = append(historiqueMessage,messageToUser)
			mutex.Unlock()

			for keyConn,valueName := range users {
				if valueName != name {
					keyConn.Write([]byte(messageToUser))
				}
			}
		}

	}

}
