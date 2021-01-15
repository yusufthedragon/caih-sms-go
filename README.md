# CAIH SMS API Go Client

Unofficial library for send SMS messages with [China-ASEAN Information Harbor](http://www.caih.com) SMS API from applications written with Go language.

- [Installation](#installation)
  - [Install the Package](#install-the-package)
  - [Set the Token](#set-the-token)
  - [Set the Channel Key](#set-the-channel-key)
- [Usages](#usages)
  - [Send a SMS Message](#send-a-sms-message)
  - [Check the Status of SMS Message](#check-the-status-of-sms-message)
  - [Batch Send SMS Message](#batch-send-sms-message)
  - [Batch Check SMS Message](#batch-check-sms-message)
- [Test](#test)
- [Contributing](#contributing)

---

## Installation

### Install the Package

Install caih-sms-go by following command:

```sh
go get -u github.com/yusufthedragon/caih-sms-go
```

Then, you can import it using:

```go
import caihsms "github.com/yusufthedragon/caih-sms-go"
```

### Set the Token

Configure package with your token obtained from CAIH.

```go
caihsms.Conf.SetToken("TOKEN")
```

### Set the Channel Key

Configure package with your channel key obtained from CAIH.

```go
caihsms.Conf.SetChannelKey("CHANNEL_KEY")
// or chain it with setToken method
caihsms.Conf.SetToken("TOKEN").SetChannelKey("CHANNEL_KEY")
```

## Usages

### Send a SMS Message

Send a single SMS request to specific number.

```go
var request = caihsms.SendSMSRequest{
    ToNumber:  toNumber, // string
    Message:   message, // string
    RequestID: requestId, // string
    SendType:  sendType, // string
    From:      from, // string, optional
}

var send, err = request.Send()
```

Usage example:

```go
var test = caihsms.SendSMSRequest{
    ToNumber:  "6282147218942",
    Message:   "Test Message",
    RequestID: "X00000001",
    SendType:  "S0001",
}

var send, err = test.Send()
```

### Check the Status of SMS Message

Check the sending status of SMS message.

```go
var request = caihsms.QueryStatusRequest{
    ToNumber:    toNumber, // string
    MessageID:   messageId, // string
}

var queryStatus, err = request.QueryStatus()
```

Usage example:

```go
var request = caihsms.QueryStatusRequest{
    ToNumber:    "6282147218942",
    MessageID:   "1348644286813773824",
}

var queryStatus, err = request.QueryStatus()
```

### Batch Send SMS Messages

Send SMS messages in batches.

```go
var request = caihsms.BatchSendRequest{
    BatchMessage:  batchMessage, // []string
    BatchToNumber: batchToNumber, // []string
    RequestID:     requestId, // string
}

var batchSend, err = request.BatchSend()
```

Usage example:

```go
var request = caihsms.BatchSendRequest{
    BatchMessage:  []string{"6282147218942", "6282147218943", "6282147218944"},
    BatchToNumber: []string{"Test Message 1", "Test Message 2", "Test Message 3"},
    RequestID:     "X00000002",
}

var batchSend, err = request.BatchSend()
```

### Batch Check SMS Messages

Check the sending status of SMS messages in batches.

```go
var request = caihsms.BatchQueryStatusRequest{
    BatchMessageID: batchMessageId, // []string
    BatchToNumber:  batchToNumber, // []string
    RequestID:      requestId, // string
}

var batchQueryStatus, err = request.BatchQueryStatus()
```

Usage example:

```go
var request = caihsms.BatchQueryStatusRequest{
    BatchMessageID: []string{"910471603446566431", "910471603446566432", "910471603446566433"},
    BatchToNumber:  []string{"6282147218942", "6282147218943", "6282147218944"},
    RequestID:      "X00000002",
}

var batchQueryStatus, err = request.BatchQueryStatus()
```

## Test

Rename the `.env.example` file to `.env` and set the Token and Channel Key with the one you obtained from CAIH. After that, you can run the test by following command:

```sh
go test -v
```

## Contributing

For any requests, bugs, or comments, please open an [issue](https://github.com/yusufthedragon/caih-sms-go/issues) or [submit a pull request](https://github.com/yusufthedragon/caih-sms-go/pulls).