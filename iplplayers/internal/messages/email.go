package messages

import (
	"fmt"
	"net/smtp"
	"time"
)

func SendEmail(activationCode string) {
	currTime := time.Now()
	currStr := currTime.Format("2006-02-01")
	// Sender data.
	from := "dkgosql@gmail.com"
	password := "test"

	// Receiver email address.
	to := []string{
		"dhananjayksharma@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "localhost"
	smtpPort := "1025"
	fmt.Println("currTime: ", currStr)
	// Message.
	msg := fmt.Sprintf("Player registraion done at time is :%s, ACTIVATION CODE:%s", currStr, activationCode)
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
