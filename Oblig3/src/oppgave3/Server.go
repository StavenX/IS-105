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

	quote := "Eg er i Arendal pga. noko AUF-greier, så eg er nok dessverre litt opptatt i dag"
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
	quote := "Har noko AUF-greier som eg må ta hånd om, så eg kan ofre meg for laget"

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