package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "192.168.126.136:8081")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\t t,		total RAM in kilobytes of the server \n")
	fmt.Print("\t u,		free RAM in kilobytes of the server \n")
	fmt.Print("\t f,		used RAM in kilobytes of the server \n")

	for {
		fmt.Print("\n Choose your command: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nResults from server: \n\t" + message + "\n")
	}

}
