package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"goemailalert/dto"
	"goemailalert/mailtemplate"

	brevo "github.com/getbrevo/brevo-go/lib"
)

const (
	// Your valid email id or brevo account email id
	fromEmail = "example@gmail.com"
	// Your name or alias
	fromName = "Jon"
	// Valid email address
	toEmail = "alice@gmail.com"
	// Name of the recipient or email address
	toName  = "alice"
	subject = "Weather Alert!"
	// threshold temperature in Celcius
	thresholdValue = 28
	// Valid City
	city = "New Delhi"

	// Email service api key
	EmailAPIKey = "xxxxxxx-xxxxxxxxxxxxxxxxxx-xxxxxxxxxxx"
	// Weather API key
	WeatherAPIKey = "xxxxxxxxxxxxxxxxxxxxxxxx"
)

func main() {
	fmt.Println("alerting system is online....")

	for {
		time.Sleep(10 * time.Second)

		// get city temperature
		temperature, err := getCityTemperature(city)
		if err != nil {
			log.Println("error while fetching city temperature", err)
			continue
		}

		// check the condition
		if temperature > thresholdValue {
			log.Println("threshold of city temperature is crossed. sending an alert email...",
				"threshold temp:", thresholdValue,
				"temperature", temperature)

			// send email
			err := sendEmail(toEmail, toName, (temperature))
			if err != nil {
				log.Println("error while sending alert email", err)

				continue
			}

			log.Println("alert email successfully sent.")
		}
	}
}

func getCityTemperature(city string) (int, error) {
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", WeatherAPIKey, city)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	v := &dto.WeatherAPIResponse{}

	err = json.Unmarshal(body, v)
	if err != nil {
		log.Println("error while unmarshalling weather api response", err)
		return 0, err
	}

	fmt.Println(string(body), v.Current.Temperature)
	return v.Current.Temperature, nil
}

func sendEmail(toEmail, toName string, data int) error {
	var ctx context.Context
	cfg := brevo.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", EmailAPIKey)

	// Create New API Client using the configuration
	br := brevo.NewAPIClient(cfg)

	// Email metadata
	meta := brevo.SendSmtpEmail{
		To: []brevo.SendSmtpEmailTo{
			{Email: toEmail, Name: toName},
		},
		Subject:     subject,
		HtmlContent: fmt.Sprintf(mailtemplate.AlertTemplate, thresholdValue, data),
		Sender: &brevo.SendSmtpEmailSender{
			Name: fromName, Email: fromEmail,
		},
	}

	// send a transactional email
	cr, res, err := br.TransactionalEmailsApi.SendTransacEmail(ctx, meta)

	if err != nil {
		return err
	}

	log.Println(res.Status, cr.MessageId)
	return nil
}
