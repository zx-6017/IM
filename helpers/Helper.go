package helpers

import (
	"IM/config"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

//request uuid
func GetUUid() string{
	var id uuid.UUID
	var err error

	for i:=0; i<3;i++ {
		id,err = uuid.NewUUID()
		if err == nil{
			return id.String()
		}
	}

	return ""
}

func WritePid(){

	file,e := os.Create(config.Conf.String("PID_FILE"))
	defer file.Close()
	if e != nil{
		log.Fatalln("pid write failed...")
		return
	}
	pid := strconv.Itoa(syscall.Getpid())
	_,e = file.WriteString(pid)
	if e!= nil{
		log.Fatalln("pid write failed...")
		return
	}
	fmt.Println("app pid is",pid)
	return
}

func DelPid(){
	err := os.Remove(config.Conf.String("PID_FILE"))
	if err != nil{
		log.Fatalln("remove pid file error...")
	}
}

func SignalExe(){
	log.Println("The program is going to shutdown,save clicks,waiting for 5s")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}