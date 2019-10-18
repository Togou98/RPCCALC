package main

import (
	"./encode"
	"flag"
	"fmt"
	"net"
)

func main() {
	host := flag.String("host", "localhost:9999", "useage ./prog hostname:port to connect to server")
	flag.Parse()
 //	var emptyslice []byte
	conn, err := net.Dial("tcp", *host)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	msg := make([]byte, 16)
	retmsg := make([]byte, 16)
	fmt.Println("报文格式为 <+ - * /> = <add sub mul div> <int> <int> like: add 100 100")
	for {
		var fn encode.Data
		fmt.Scanln(&fn.Operator, &fn.Lval, &fn.Rval)
		if len(fn.Operator) != 3 {
			fmt.Println("func name error")
			continue
		}
		msg = fn.ToByte()
		conn.Write(msg)
		conn.Read(retmsg)
		val := encode.Byte2Int(retmsg)
		fmt.Printf("%d %s %d is: %d \n", fn.Lval, string(fn.Operator), fn.Rval, val)
	}
}
