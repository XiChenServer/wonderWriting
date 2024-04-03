package verification

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

func init() {
	// 加载 .env 文件中的环境变量
	if err := godotenv.Load("/home/zwm/GolandProjects/wonderWriting/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var TimeExpiration = 7 * 24 * 60 * 60

func SendEmailVerificationCode(email, verificationCode1 string) error {

	//发送对象
	recipient := email
	// 生成验证码
	verificationCode := verificationCode1

	// 构建邮件内容
	subject := "验证码"
	body := fmt.Sprintf("你的验证码是：%s", verificationCode)
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", "15294440097@163.com")
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)
	// 转换 SMTP_PORT 为整数
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		fmt.Println("Error converting SMTP_PORT to integer:", err)
		return err
	}
	// 创建SMTP客户端
	dialer := gomail.NewDialer(smtpServer, smtpPort, senderEmail, senderPassword)

	// 发送邮件
	err = dialer.DialAndSend(message)
	if err != nil {
		fmt.Println("发送邮件失败:", err)
		return err
	}

	return nil
}
