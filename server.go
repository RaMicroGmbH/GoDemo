package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	res string
	ipadd = "ServerIPAddr"
	port  = ":1234"
	defaultBytes = 2048
)
// A Simple function to verify error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func InitUDPClient() string {
	conn, err := net.Dial("udp", ipadd + port)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return ""
	}
	defer conn.Close()
	i:= 0
	for {
		msg := strconv.Itoa(i) + " :: Hi Server !"
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)		

		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
		return msg
	}

}

func InitUDPServer() string {

	ServerAddr, err := net.ResolveUDPAddr("udp4", ":1234")
	CheckError(err)

	//listen at selected port
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 2048)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		res = string(buf[0:n]) + " from " + addr.String()
		if err != nil {
			fmt.Println("Error : ", err)
		}
		go sendResponse(ServerConn, addr)
		return res	

	}		

}

func main(){
 //call to udp server
 InitUDPServer()
 
 
}
