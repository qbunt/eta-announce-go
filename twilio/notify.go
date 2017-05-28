package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Notify(phoneNumber string, eta string) error {
	accountSid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	v.Set("To", phoneNumber)
	v.Set("From", os.Getenv("TWILIO_PHONE"))
	message := strings.Replace(os.Getenv("MESSAGE"), "{{NAME}}", os.Getenv("NAME"), 1)
	etaMessage := strings.Replace(message, "{{ETA}}", eta, 1)
	v.Set("Body", etaMessage)

	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
			err = nil
			return err
		} else {
			return err
		}
	} else {
		fmt.Println(resp.Status)
	}
    return nil
}
