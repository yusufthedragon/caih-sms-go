package caihsms

// Config struct for API Call configuration.
type Config struct {
	token      string
	ChannelKey string
	baseURL    string
}

// BatchSendRequest is struct for batch send request.
type BatchSendRequest struct {
	BatchMessage  []string `json:"batchMessage" validate:"required"`
	BatchToNumber []string `json:"batchToNumber" validate:"required"`
	RequestID     string   `json:"requestId" validate:"required"`
	Token         string   `json:"token" validate:"required"`
}

// BatchSendResponse is struct for batch send response.
type BatchSendResponse struct {
	Body     interface{} `json:"body"`
	Desc     string      `json:"desc"`
	RespCode string      `json:"respCode"`
	Success  bool        `json:"success"`
}

// BatchQueryStatusRequest is struct for batch query status request.
type BatchQueryStatusRequest struct {
	BatchMessageID []string `json:"batchMessageId" validate:"required"`
	BatchToNumber  []string `json:"batchToNumber" validate:"required"`
	RequestID      string   `json:"requestId" validate:"required"`
	Token          string   `json:"token" validate:"required"`
}

// BatchQueryStatusResponse is struct for batch query status response.
type BatchQueryStatusResponse struct {
	Body     interface{} `json:"body"`
	Desc     string      `json:"desc"`
	RespCode string      `json:"respCode"`
	Success  bool        `json:"success"`
}

// QueryStatusRequest is struct for query status request.
type QueryStatusRequest struct {
	MessageID string `json:"messageId" validate:"required"`
	Token     string `json:"token" validate:"required"`
	ToNumber  string `json:"toNumber" validate:"required"`
}

// QueryStatusResponse is struct for query status response.
type QueryStatusResponse struct {
	Desc      string `json:"desc"`
	MessageID string `json:"messageId"`
	RespCode  string `json:"respCode"`
	Success   bool   `json:"success"`
}

// SendSMSRequest is struct for send SMS request.
type SendSMSRequest struct {
	From      string `json:"from"`
	Message   string `json:"message" validate:"required"`
	RequestID string `json:"requestId" validate:"required"`
	SendType  string `json:"sendType" validate:"required"`
	Token     string `json:"token" validate:"required"`
	ToNumber  string `json:"toNumber" validate:"required"`
}

// SendSMSResponse is struct for send SMS response.
type SendSMSResponse struct {
	Desc      string `json:"desc"`
	MessageID string `json:"messageId"`
	RespCode  string `json:"respCode"`
	Success   bool   `json:"success"`
}
