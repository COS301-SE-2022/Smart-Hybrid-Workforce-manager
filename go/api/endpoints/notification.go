package endpoints

import (
	"api/utils"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gorilla/mux"
)

func NotificationHandlers(router *mux.Router) error {
	router.HandleFunc("/send", SendNotificationHandler).Methods("POST")
	return nil
}

func SendNotificationHandler(writer http.ResponseWriter, request *http.Request) {

	//Sender data
	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")

	//Receiver data
	to := []string{
		"tash2814@gmail.com",
	}

	//smtp Server
	smtpHost := "smtp.gmail.com"
	smptPort := "587"

	//Message
	message := []byte("From: archecapstoneteam@gmail.com\r\n" +
		"To: tash2814@gmail.com\r\n" +
		"Subject: Booking Confirmation\r\n\r\n" +
		"Your booking has been confirmed!\r\n")

	//Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//Sending email
	err := smtp.SendMail(smtpHost+":"+smptPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Notification email sent")

	utils.Ok(writer, request)
}
