package caihsms

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// List of invalid response codes from API.
var invalidResponseCode = map[string]string{
	"00001": "System Internal Error",
	"00002": "Invalid Request Parameter",
	"00003": "Unauthorized Information",
	"00004": "The API Exceeds the Access Frequency Limit",
	"00005": "The API Call Returns an Exception",
	"02001": "Number is Wrong",
	"02002": "Send Failure",
	"03001": "Query Failure",
	"99003": "SMS Send Failed",
}

// SendRequest sends request to API.
func SendRequest(apiEndpoint string, parameters interface{}, response interface{}) {
	// Get the encoded md5Sum and request parameters
	var md5SumString, requestParameters = SetRequestBody(parameters)
	var req, err = http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestParameters))

	if err != nil {
		log.Fatalln(err)
	}

	// Set the request headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("md5Sum", md5SumString)

	var client = &http.Client{}
	var resp, errRequest = client.Do(req)

	if errRequest != nil {
		log.Fatalln(errRequest)
	}

	defer resp.Body.Close()

	// Read the response body
	var body, errResponse = ioutil.ReadAll(resp.Body)

	if errResponse != nil {
		log.Fatalln(errResponse)
	}

	// Decode the response JSON into target struct
	json.Unmarshal(body, &response)
}

// SetRequestBody generates the encoded request parameter and its md5Sum.
func SetRequestBody(parameters interface{}) (string, []byte) {
	// Encode the request parameters to JSON
	requestParameters, err := json.Marshal(parameters)

	if err != nil {
		log.Fatalln(err)
	}

	// Combine the JSON string with ChannelKey
	jsonString := string(requestParameters) + Conf.ChannelKey

	// Generate MD5Sum of JSON string
	md5SumByte := md5.Sum([]byte(jsonString))
	md5SumString := hex.EncodeToString(md5SumByte[:])

	return md5SumString, requestParameters
}

// ValidateResponse for check if response from API is not an error.
func ValidateResponse(responseCode string) error {
	textError, valid := invalidResponseCode[responseCode]

	if valid {
		return errors.New(textError)
	}

	return nil
}
