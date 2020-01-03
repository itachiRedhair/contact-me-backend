package contactmebackend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type email struct {
	EmailAdress string `json:"emailAddress"`
	Name        string `json:"name"`
	Message     string `json:"message"`
}

// SendEmail function sents email coming from contact me form of portfolio
func SendEmail(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e email
	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Email Address: %s!", e.EmailAdress)
	fmt.Fprintf(w, "Name: %s!", e.Name)
	fmt.Fprintf(w, "Message: %s!", e.Message)

	from := mail.NewEmail(e.Name, e.EmailAdress)
	subject := "Contacted via portfolio website"
	to := mail.NewEmail("Akshay Milmile", "akshay.milmile@gmail.com")
	plainTextContent := e.Message
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, plainTextContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "Status code: %d", response.StatusCode)
		fmt.Fprintf(w, "Body: %s", response.Body)
		fmt.Fprintf(w, "Headers: %s", response.Headers)
	}
}
