package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"log"
)

// R type represents an object of Recaptcha and has public property Secret,
// which is secret obtained from google recaptcha tool admin interface
type R struct {
	Secret    string
	lastError []string
}

// Struct for parsing json in google's response
type googleResponse struct {
	Success    bool `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname string `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

// url to post submitted re-captcha response to
var postURL = "https://www.google.com/recaptcha/api/siteverify"

// Verify method, verifies if current request have valid re-captcha response and returns true or false
// This method also records any errors in validation.
// These errors can be received by calling LastError() method.
func (r *R) Verify(req http.Request) *googleResponse {
	response := req.FormValue("g-recaptcha-response")
	return r.VerifyResponse(response)
}

// VerifyResponse is a method similar to `Verify`; but doesn't parse the form for you.  Useful if
// you're receiving the data as a JSON object from a javascript app or similar.
func (r *R) VerifyResponse(response string) *googleResponse {
	r.lastError = make([]string, 1)
	client := &http.Client{Timeout: 5 * time.Second}
	log.prin
	resp, err := client.PostForm(postURL,
		url.Values{"secret": {r.Secret}, "response": {response}})
	if err != nil {
		r.lastError = append(r.lastError, err.Error())
		//return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.lastError = append(r.lastError, err.Error())
		//return false
	}
	gr := new(googleResponse)
	err = json.Unmarshal(body, gr)
	if err != nil {
		r.lastError = append(r.lastError, err.Error())
		//return false
	}
	if !gr.Success {
		r.lastError = append(r.lastError, gr.ErrorCodes...)
	}
	return gr
}

// LastError returns errors occurred in last re-captcha validation attempt
func (r R) LastError() []string {
	return r.lastError
}
