package sms

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// Conf is an instance of Config struct.
var Conf = Config{
	baseURL: "http://sms.caihcom.com",
}

// The validator instance.
var validation = validator.New()

// BatchSend is function to send SMS messages in batches.
func (batchSendRequest *BatchSendRequest) BatchSend() (BatchSendResponse, error) {
	batchSendRequest.Token = Conf.token
	var err = validation.Struct(batchSendRequest)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var apiEndpoint = Conf.baseURL + "/sms/batchSend"
	var response = BatchSendResponse{}

	SendRequest(apiEndpoint, batchSendRequest, &response)

	var responseError = ValidateResponse(response.RespCode)

	return response, responseError
}

// BatchQueryStatus is function to check the sending status of SMS messages in batches.
func (batchQueryStatusRequest *BatchQueryStatusRequest) BatchQueryStatus() (BatchQueryStatusResponse, error) {
	batchQueryStatusRequest.Token = Conf.token
	var err = validation.Struct(batchQueryStatusRequest)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var apiEndpoint = Conf.baseURL + "/sms/batchQueryStatus"
	var response = BatchQueryStatusResponse{}

	SendRequest(apiEndpoint, batchQueryStatusRequest, &response)

	var responseError = ValidateResponse(response.RespCode)

	return response, responseError
}

// QueryStatus is function to check the sending status of a SMS message.
func (queryStatusRequest *QueryStatusRequest) QueryStatus() (QueryStatusResponse, error) {
	queryStatusRequest.Token = Conf.token
	var err = validation.Struct(queryStatusRequest)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var apiEndpoint = Conf.baseURL + "/sms/queryStatus"
	var response = QueryStatusResponse{}

	SendRequest(apiEndpoint, queryStatusRequest, &response)

	var responseError = ValidateResponse(response.RespCode)

	return response, responseError
}

// Send is function to send a single SMS message to specific number.
func (sendSMSRequest *SendSMSRequest) Send() (SendSMSResponse, error) {
	sendSMSRequest.Token = Conf.token
	var err = validation.Struct(sendSMSRequest)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var apiEndpoint = Conf.baseURL + "/sms/send"
	var response = SendSMSResponse{}

	SendRequest(apiEndpoint, sendSMSRequest, &response)

	var responseError = ValidateResponse(response.RespCode)

	return response, responseError
}

// SetChannelKey sets the required Channel Key for API Call.
func (c *Config) SetChannelKey(channelKey string) {
	c.ChannelKey = channelKey
}

// SetToken sets the required Token for API Call.
func (c *Config) SetToken(token string) *Config {
	c.token = token

	return c
}
