package sendgrid

import (
	"fmt"
	"log"

	vsend "github.com/sendgrid/sendgrid-go"
	vhelper "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(email string, subj string, key string, txtMsg string, htmlMsg string) (err error) {

	//	fmt.Printf("<<<< MAIL >>>>\n% 15s : %s\n% 15s : %s\n% 15s : %s\n% 15s : %s\n",
	//		"ADDR", email, "SUBJ", subj, "TXT", txtMsg, "HTML", htmlMsg)

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", email)

	/*

		from := vvmail.NewEmail("Raspberry of Arduino", "raspberry@dev-box.com")
		subject := subj // "Raspberry of Arduino Contact Info"
		//	to := mail.NewEmail("Example User", "test@example.com")
		to := vvmail.NewEmail("Administrator", email)
		//	plainTextContent := "and easy to do anywhere, even with Go"
		//	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
		plainTextContent := txtMsg
		htmlContent := htmlMsg
		message := vvmail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

		if err = os.Setenv("SENDGRID_API_KEY", "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"); err != nil {
			fmt.Printf("Sakmoto-Yakmoto! ===> nevaru turpināt\n%v\n", err)
			return
		}

		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", subject)
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", plainTextContent)
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", email)
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", htmlContent)
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", email)

		client := vvgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)

		if err != nil {
			log.Println("*** ERROR ***\n", err)
		} else {
			fmt.Println("\n*** STATUS ***\n", response.StatusCode)
			fmt.Println("\n*** BODY ***\n", response.Body)
			fmt.Println("\n*** HEADERS ***\n", response.Headers)
		}
	*/

	from := vhelper.NewEmail("Raspberry of Arduino", "raspberry@dev-box.com")
	subject := subj //"Sending with SendGrid is Fun"
	//	to := mail.NewEmail("Example User", "test@example.com")
	to := vhelper.NewEmail("Example User", email)
	plainTextContent := txtMsg
	//htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	htmlContent := htmlMsg
	message := vhelper.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	//	if err = os.Setenv("SENDGRID_API_KEY", "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"); err != nil {
	//		fmt.Printf("Sakmoto-Yakmoto! ===> nevaru turpināt\n%v\n", err)
	//		return
	//	}

	//client := vsend.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	client := vsend.NewSendClient(key)

	response, err := client.Send(message)
	if err != nil {
		fmt.Println("*** ERROR ***")
		log.Println(err)
	} else {
		fmt.Println("*** SUCCESS ***")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return
}
