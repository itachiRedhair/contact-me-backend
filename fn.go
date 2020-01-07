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

// Email struct is for the data coming from contact me form
type email struct {
	EmailAddress string `json:"emailAddress"`
	Name         string `json:"name"`
	Message      string `json:"message"`
}

// SendEmail function sents email coming from contact me form of portfolio
func SendEmail(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e email

	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}

	from := mail.NewEmail(e.Name, e.EmailAddress)
	subject := "Contacted via portfolio website"
	to := mail.NewEmail("Akshay Milmile", "akshay.milmile@gmail.com")
	plainTextContent := e.Message
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, plainTextContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if response.StatusCode != http.StatusAccepted {
		if response.StatusCode != http.StatusOK {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Success")
	}
}
