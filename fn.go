package contactmebackend

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Email struct defines the data coming via REST API
type email struct {
	EmailAddress string `json:"emailAddress"`
	Name         string `json:"name"`
	Message      string `json:"message"`
}

// SendEmail function sents email with the data coming from REST API
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var e email
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)

	if err != nil {
		panic(err)
	}

	// Prepare the email content
	from := mail.NewEmail(e.Name, e.EmailAddress)
	subject := "Contacted via portfolio website"
	to := mail.NewEmail("Akshay Milmile", "akshay.milmile@gmail.com")
	plainTextContent := e.Message

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, plainTextContent)

	// Get the SendGrid API key from environment variable
	sendGridAPIKey := os.Getenv("SENDGRID_API_KEY")
	// Generate a SendGrid Send Client
	client := sendgrid.NewSendClient(sendGridAPIKey)
	// Send the message
	response, err := client.Send(message)

	if response.StatusCode != http.StatusAccepted {
		if response.StatusCode != http.StatusOK {
			// Request failed
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something bad happened!"))
			w.Write([]byte(err.Error()))
		}
	} else {
		// Request successful

		// Setting header to allow cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Success"))
	}
}
