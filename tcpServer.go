package main

import (
	"fmt"
	"net"

)

func process(conn net.Conn){
	defer conn.Close()
	for {
		var buf [1024]byte
		n,err := conn.Read(buf[:])
		if err != nil{
			fmt.Println("read err",err)
			break
		}
		str := string(buf[:n])
		fmt.Println("recv content :",str)
	}

}
//tcp servers
func TcpServer(){
	listener,err := net.Listen("tcp","0.0.0.0:9000")
	if err != nil {
		fmt.Println("listener err",err)
		return
	}
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("conn err",err)
			continue
		}
		go process(conn)
	}


}
func main(){
	TcpServer()
}
