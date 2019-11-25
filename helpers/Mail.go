package helpers

import (
	"IM/config"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/smtp"
	"strings"
)
type mailConn struct {
	User string
	Pass string
	Host string
	NickName string
	Port int
}

func SendMail(mailTo []string,subject string,body string){
	debug,err := config.Conf.Bool("APP_DEBUG")
	if err != nil{
		fmt.Println("read app debug err",err)
	}
	// 调试模式不需要发送邮件
	if debug {
		return
	}
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mail_conn := mailConn{
		User: config.Conf.String("MAIL_USER"),
		Pass: config.Conf.String("MAIL_PASS"),
		Host: config.Conf.String("MAIL_HOST"),
	}
	mail_conn.Port,_ = config.Conf.Int("MAIL_PORT")

	mail_msg := gomail.NewMessage()

	mail_msg.SetHeader("From",mail_conn.User)
	mail_msg.SetHeader("To",mailTo...)
	mail_msg.SetHeader("Subject",subject)
	mail_msg.SetBody("text/html",body)

	sender := gomail.NewDialer(mail_conn.Host,mail_conn.Port,mail_conn.User,mail_conn.Pass)

	err = sender.DialAndSend(mail_msg)

	if err != nil {
		fmt.Println("mail send err",err)
		return
	}


}


func SendMailSmtp(mailTo []string,subject string,body string){
	debug,err := config.Conf.Bool("APP_DEBUG")
	if err != nil{
		fmt.Println("read app debug err",err)
	}
	// 调试模式不需要发送邮件
	if debug {
		return
	}
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mail_conn := mailConn{
		User: config.Conf.String("MAIL_USER"),
		Pass: config.Conf.String("MAIL_PASS"),
		Host: config.Conf.String("MAIL_HOST"),
		NickName:config.Conf.String("MAIL_NICKNAME"),
	}
	mail_conn.Port,_ = config.Conf.Int("MAIL_PORT")

	auth := smtp.PlainAuth("",mail_conn.User,mail_conn.Pass,mail_conn.Host)

	nickname := mail_conn.NickName
	contentType := "Content-Type: text/html; charset=UTF-8"

	msg := []byte("To: " + strings.Join(mailTo, ",") + "\r\nFrom: " + nickname + "<" + mail_conn.User + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)


	err = smtp.SendMail("smtp.exmail.qq.com:25",auth,mail_conn.User,mailTo,msg)


	if err != nil{
		fmt.Println("smtp send mail err ",err)
	}



}
