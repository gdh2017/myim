package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){

	conn,err := net.Dial("tcp","0.0.0.0:9000")
	defer conn.Close()
	if err != nil {
		fmt.Println("conn err:",err)
		return
	}

	inputReader := bufio.NewReader(os.Stdin)

	for{
		input,err := inputReader.ReadString('\n')
		if err != nil{
			fmt.Println("read form stdin err:",err)
			break
		}

		trimmedInput := strings.TrimSpace(input)

		_,err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("write err :",err)
			break
		}
	}
}
