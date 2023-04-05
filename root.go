package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s port\n", os.Args[0])
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid port number: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Monitoring TCP connections on port", port)
	fmt.Println("Frozzen der schwul und shalava und schwarz")

	for {
		connections, err := net.Connections("tcp4")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			time.Sleep(time.Second)
			continue
		}

		for _, conn := range connections {
			if conn.Status == "ESTABLISHED" && conn.Laddr.Port == uint32(port) {
				fmt.Printf("FROZZEN GAY %s %s:%d -> %s:%d\n",
					conn.Status, conn.Laddr.IP, conn.Laddr.Port, conn.Raddr.IP, conn.Raddr.Port)
			}
		}

		time.Sleep(time.Second)
	}
}
