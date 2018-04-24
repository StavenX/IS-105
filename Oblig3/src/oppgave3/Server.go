package main

import "net"
import (
	"fmt"
	"bufio"
	"os"
	"time"
)

func main() {
	go tcp()
	go udp()
	time.Sleep(time.Minute *10)
}

func tcp() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":17")

	// accept connection on port
	conn, _ := ln.Accept()

	quote := "The weather is nice today, and so are you!"
	// run loop forever (or until ctrl-c)
	for {
		bufio.NewReader(conn).ReadString('\n')
		conn.Write([]byte(quote + "\n"))
	}
}

func udp() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr,err := net.ResolveUDPAddr("udp",":17")
	CheckError(err)
	quote := "The world is nice. Appreciate it!"

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		_,addr,err := ServerConn.ReadFromUDP(buf)
		// send new string back to client
		ServerConn.WriteToUDP([]byte(quote + "\n"),addr)

		if err != nil {
			fmt.Println("Error: ",err)
		}
	}
}
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}}