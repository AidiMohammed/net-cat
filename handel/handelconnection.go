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
  	var nameClient string
  	if len(users) == 10 {
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
	
	for {
		conn.Write([]byte("\n[ENTER YOUR NAME]:"))
		name,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture : ",err)
			return
		}
		if name == ""  {
			conn.Write([]byte("Notice: Provide a non-blank name: "))
			continue
		}
		name = strings.ReplaceAll(name, "\n", "")
 		if (!tools.IsVisibleString(name)) {
			conn.Write([]byte("Notice: Provide a name that holds only printable characters:"))
			continue 
		}
		name = strings.TrimSpace(name)
		foundName := false
		mutex.Lock()
		for _,nameValue := range users {
			if nameValue == name {
				foundName = true
				break
			}
		}
		if !foundName {

			if len(name) > 50  {
				conn.Write([]byte("Notice: : the size name is very long (max size : 50)"))
				continue
			}
			users[conn] = name
			mutex.Unlock()
			fmt.Printf("Add new user %v address %v \n",name,conn.RemoteAddr())
			broadCastMessage(fmt.Sprintf("%v has joined our chat ...\n",name),conn);
			nameClient = name
			break
		} else {
			mutex.Unlock()
			conn.Write([]byte(fmt.Sprintf("Notice: This name (%v) is already taken, try another one: ",name)))
			continue
		}
	}

  	setHistoriqueMessage(conn)

	for {
		now := time.Now()
		headerMessage := fmt.Sprintf("[%v][%v]:",now.Format("2006-01-02 15:04:05"),nameClient)
		message,err := reader.ReadString('\n')
		if err != nil {
      	broadCastMessage(fmt.Sprintf("%v to disconnect from chat ...",nameClient),conn)
      	mutex.Lock()
        delete(users,conn)
        fmt.Printf(fmt.Sprintf("delete user : %v\n",nameClient))
			mutex.Unlock()
			fmt.Printf(fmt.Sprintf("User %s déconnecté.\n",nameClient))
			break
		}
		message = strings.ReplaceAll(message, "\n", "")
		if(!tools.IsVisibleString(message)) {
			conn.Write([]byte("Notice: Provide a messsage that holds only printable characters:"))
			continue
		}
		message = strings.TrimSpace(message)

		if len(message) > 500  {
			conn.Write([]byte("Notice: : the size message is very long (max size : 500)"))
			continue
		}

		if message != "" {
      		saveHistoriqueMessage(headerMessage+message)
      		broadCastMessage(headerMessage+message,conn)
			conn.Write([]byte(headerMessage+message))
		}
	}
}

// broadCastMessage sends a message to all users except the sender
func broadCastMessage(message string,sender net.Conn) {
	if message == "" {
		return
	}
    mutex.Lock()
	defer mutex.Unlock()

	for keyConn,valueName := range users {
        if keyConn == sender {
		  continue
        }
		_,err := keyConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erreur d'envoi à %v : %v",valueName, err)
			continue
		}
        fmt.Printf("brodcast message to user %v addres : %v\n",keyConn.RemoteAddr(),valueName)
	}
}

func saveHistoriqueMessage(message string) {
  mutex.Lock()
    historiqueMessage = append(historiqueMessage,message)
  mutex.Unlock()
}

func setHistoriqueMessage(conn net.Conn){
  mutex.Lock()
    for _,message := range historiqueMessage {
      conn.Write([]byte(message))
      fmt.Printf("Brodcast Historique message to ip %v message (%v) \n",conn.RemoteAddr(),strings.TrimSuffix(message,"\n"))
    }
  mutex.Unlock()
}