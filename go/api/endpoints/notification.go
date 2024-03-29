package endpoints

import (
	"api/data"
	"fmt"
	"lib/utils"
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

	var notification data.Notification
	err := utils.UnmarshalJSON(writer, request, &notification)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	//Sender data
	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")

	//Receiver data
	to := []string{
		*notification.To,
	}

	//smtp Server
	smtpHost := "smtp.gmail.com"
	smptPort := "587"

	//Message
	message := []byte("From: archecapstoneteam@gmail.com\r\n" +
		"To: " + *notification.To + "\r\n" +
		"Subject: Booking Confirmation\r\n\r\n" +
		"Your booking has been confirmed!\n\n" +
		"Start Date: " + *notification.StartDate + "\n" +
		"Start Time: " + *notification.StartTime + "\n" +
		"End Date: " + *notification.EndDate + "\n" +
		"End Time: " + *notification.EndTime + "\r\n")

	//Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//Sending email
	err = smtp.SendMail(smtpHost+":"+smptPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Notification email sent")

	utils.Ok(writer, request)
}

func SendNotification(message string, recipient string) {
	//Receiver data
	to := []string{
		recipient,
	}

	//Sender data
	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")

	//smtp Server
	smtpHost := "smtp.gmail.com"
	smptPort := "587"

	//Message
	theMessage := []byte("From: archecapstoneteam@gmail.com\r\n" +
		"To: " + recipient + "\r\n" + message)

	//Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//Sending email
	err := smtp.SendMail(smtpHost+":"+smptPort, auth, from, to, theMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Notification email sent")

}
