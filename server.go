package main

import (
	"net"
	"fmt"
	"bufio"
	"bitbucket.org/bertimus9/systemstat"
	"strconv"
	"strings"
)

var mem systemstat.MemSample

func main() {
	
	fmt.Println("	waiting for server, wait for it!")
	
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		msg := strings.ToLower(string(message))
		mem = systemstat.GetMemSample()

		if strings.TrimSpace(msg) == "t" {
			returnmsg := "Total available RAM in kb: " + strconv.FormatUint(mem.MemTotal, 10) + "k total.";
			conn.Write([]byte(returnmsg + "\n"))
		} else if strings.TrimSpace(msg) == "u" {
			returnmsg := "Used RAM in kb: " + strconv.FormatUint(mem.MemUsed, 10) + "k used.";
			conn.Write([]byte(returnmsg + "\n"))
		} else if strings.TrimSpace(msg) == "f" {
			returnmsg := "Used RAM in kb: " + strconv.FormatUint(mem.MemUsed, 10) + "k free.";
			conn.Write([]byte(returnmsg + "\n"))
		} else {
			returnmsg := "Error: Unidentified command.";
			conn.Write([]byte(returnmsg + "\n"))
		}
	}
}
