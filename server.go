package main
import(
		"fmt"
		"./encode"
		"net"
		"flag"
	//	"time"
		"bytes"
	//	"os"
)

var add = []byte{'a','d','d'}
var substract = []byte{'s','u','b'}
var	multiply = []byte{'m','u','l'}
var divide = []byte{'d','i','v'}


func main(){
			listenaddr := flag.String("server host","localhost:9999","useage ./prog hostname:port")
			flag.Parse()
	ln , err := net.Listen("tcp",*listenaddr)
	if err != nil{
			fmt.Println(err)
	}
	for{
			conn , err := ln.Accept()
			if err != nil{
					fmt.Println(err)
			}
			fmt.Println("Connect ok")
			go handleconn(conn)
	}

}

func handleconn(conn net.Conn) {
		defer conn.Close()
		var Lval,Rval int = 0 , 0
		msg := make([]byte,16)
		var fn []byte
		var result int
		var bytearr [][]byte
		//var data encode.Data
		for true {
				ln, err := conn.Read(msg)
				if err != nil{
						fmt.Println(err)
				}
				
				if ln ==0 {
						fmt.Println("客户端已经关闭连接了")
						return 
				}

				bytearr = encode.Bytesplit(msg)
				if len(bytearr[0]) != 3{
						continue
				}
				fn = bytearr[0]
				Lval = encode.Byte2Int(bytearr[1])
				Rval = encode.Byte2Int(bytearr[2])
				if bytes.Equal(fn,add){
						result = Lval + Rval
						conn.Write(encode.Int2Byte(result))
				}else if bytes.Equal(fn,substract){
						result = Lval - Rval
						conn.Write(encode.Int2Byte(result))
				}else if bytes.Equal(fn,multiply){
						result = Lval * Rval
						conn.Write(encode.Int2Byte(result))
				}else if bytes.Equal(fn,divide){
						result = Lval / Rval
						conn.Write(encode.Int2Byte(result))
				}else {
						errmsg := string("no such a function call")
						fmt.Println(errmsg)
				}
				//msg = emptyslice
			}
		}


