package logger

import "time"

type (
	// Key context
	Key int
)

const (
	services = "Internal Service"

	logKey = Key(36)
)

// Data is data standard output
type Data struct {
	RequestID     string       `json:"RequestID"`
	TimeStart     time.Time    `json:"TimeStart"`
	UserCode      string       `json:"UserCode"`
	Device        string       `json:"Device"`
	Service       string       `json:"Service"`
	Host          string       `json:"Host"`
	Endpoint      string       `json:"Endpoint"`
	RequestHeader string       `json:"RequestHeader"`
	RequestBody   string       `json:"RequestBody"`
	StatusCode    int          `json:"StatusCode"`
	Response      string       `json:"Response"`
	ExecTime      float64      `json:"ExecutionTime"`
	Messages      []string     `json:"Messages"`
	ThirdParty    []ThirdParty `json:"3rdParty"`
}

// ThirdParty is data logging for any request to outside
type ThirdParty struct {
	URL           string  `json:"URL"`
	RequestHeader string  `json:"RequestHeader"`
	RequestBody   string  `json:"RequestBody"`
	Response      string  `json:"Response"`
	StatusCode    int     `json:"StatusCode"`
	ExecTime      float64 `json:"ExecTime"`
}
