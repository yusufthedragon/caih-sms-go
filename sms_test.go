package sms

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

// Initiate the test.
func init() {
	var err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	Conf.SetToken(os.Getenv("TOKEN")).SetChannelKey(os.Getenv("CHANNEL_KEY"))
}

// Test the Batch Send API.
func TestBatchSend(t *testing.T) {
	var test = BatchSendRequest{
		BatchMessage:  []string{"6282147218942", "6282147218943", "6282147218944"},
		BatchToNumber: []string{"Test Message 1", "Test Message 2", "Test Message 3"},
		RequestID:     strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
	}

	var batchSend, err = test.BatchSend()

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(batchSend)
}

// Test the Batch Query Status API.
func TestBatchQueryStatus(t *testing.T) {
	var test = BatchQueryStatusRequest{
		BatchMessageID: []string{"910471603446566431", "910471603446566432", "910471603446566433"},
		BatchToNumber:  []string{"6282147218942", "6282147218943", "6282147218944"},
		RequestID:      "160595797880071",
	}

	var batchQueryStatus, err = test.BatchQueryStatus()

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(batchQueryStatus)
}

// Test the Send SMS API.
func TestSend(t *testing.T) {
	var test = SendSMSRequest{
		ToNumber:  "6282147218942",
		Message:   "Test Message",
		RequestID: strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
		SendType:  "S0002",
	}

	var send, err = test.Send()

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(send)
}

// Test the Query Status API.
func TestQueryStatus(t *testing.T) {
	var test = QueryStatusRequest{
		ToNumber:  "6282147218942",
		MessageID: "1348644286813773824",
	}

	var queryStatus, err = test.QueryStatus()

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(queryStatus)
}
