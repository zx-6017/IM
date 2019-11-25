package chatroom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func Server(request *gin.Context){
	fmt.Println("server listening 8089")

	listener,err := net.Listen("tcp","0.0.0.0:8089")
	defer  listener.Close()

	if err != nil{
		fmt.Println("server listen err ",err)
		return
	}
	for{
		fmt.Println("server wait connect...")

		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("listener accept error ",err)
			continue
		}

		go process(conn)

	}

}

func process(conn net.Conn){


}