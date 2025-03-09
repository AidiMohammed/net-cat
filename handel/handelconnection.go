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

  if len(users) > 9 {
    conn.Write([]byte("The maximum number of connections allowed has been reached. Please try again later."))
    return
  }

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
    found := false
    for _,nameValue := range users {
      if nameValue == name {
        found = true
        break
      }
    }

    if !found {
      users[conn] = name
      fmt.Printf("Add new user %v address %v \n",name,conn.RemoteAddr())
      
      for valueConn,valueName := range users {
        if valueName != name {
          valueConn.Write([]byte(fmt.Sprintf("%v has joined our chat...\n",name)))
        }
      }
    } else {
      fmt.Printf("User %s deconnecté %v",conn.RemoteAddr())
      return
    }
	mutex.Unlock()

  broadCastHistoriqueMessage(conn)

	for {
		now := time.Now()
		message,err := reader.ReadString('\n')
		if err != nil {
      broadCastMessage(fmt.Sprintf("%v to disconnect from chat ...",name),conn)
      mutex.Lock()
        delete(users,conn)
        fmt.Printf(fmt.Sprintf("delete user : %v\n",name))
			mutex.Unlock()
			fmt.Printf(fmt.Sprintf("User %s déconnecté.\n",name))
			break
		}

		message = strings.TrimSpace(message)

		if message != "" {
			messageBroadcast := fmt.Sprintf("[%v][%v]:%v\n",now.Format("2006-01-02 15:04:05"),name,message)	
      saveHistoriqueMessage(messageBroadcast)
      broadCastMessage(messageBroadcast,conn)
		}
	}
}

func broadCastMessage(message string,sender net.Conn) {
	if message != "" {
    mutex.Lock()
		  for keyConn,valueName := range users {
        if keyConn == sender {
          keyConn.Write([]byte(fmt.Sprintf("\033[1A\033[2K%v",message)))
          continue
        }
		  	keyConn.Write([]byte(message))
        fmt.Printf("brodcast message to user %v addres : %v\n",keyConn.RemoteAddr(),valueName)
		  }
    mutex.Unlock()
	}
}

func saveHistoriqueMessage(message string) {
  mutex.Lock()
    historiqueMessage = append(historiqueMessage,message)
  mutex.Unlock()
}

func broadCastHistoriqueMessage(conn net.Conn){
  mutex.Lock()
    for _,message := range historiqueMessage {
      conn.Write([]byte(message))
      fmt.Printf("Brodcast Historique message to ip %v message (%v) \n",conn.RemoteAddr(),strings.TrimSuffix(message,"\n"))
    }
  mutex.Unlock()
}