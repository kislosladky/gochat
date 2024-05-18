package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func listenAndPrint() {
	p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 8080,
        IP: nil,
    }
    socket, err := net.ListenUDP("udp", &addr) //Слушаю все ip на 5050 порте
    if err != nil {
        fmt.Printf("Error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := socket.ReadFromUDP(p) //читаю сообщение из сокета
        fmt.Printf("From %v: %s", remoteaddr, p)
		clear(p)
        if err !=  nil {
            fmt.Printf("Error %v", err)
            continue
        }

		fmt.Print("Write an IP and message\n> ")
    }
}
func main() {
	go listenAndPrint() //запускаю поток, слушающий сокет

	var ip string
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Write an IP (type \"all\" for broadcast) and message\n> ")
		fmt.Scanln(&ip)
		fmt.Print("> ")
		ip = strings.ToLower(ip)
		if ip == "all" {             //ставлю broadcast адрес
			ip = "192.168.0.255"
		}
		message, msgErr := in.ReadString('\n')
		if msgErr != nil {
			panic(msgErr)
		}
		
		socket, err := net.Dial("udp", ip + ":8080") // создаю сокет для отправки
		if err != nil {
			fmt.Printf("Error %v", err)
			return
		}
		
		_, socketErr := socket.Write([]byte(message)) //отправляю сообщение
		if socketErr != nil {
			fmt.Printf("Error %v", socketErr)
		}
		
		socket.Close()
	}
}
