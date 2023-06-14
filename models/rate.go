package models

type Rate struct {
	ClientID       string `json:"clientID"`
	Tokens         int    `json:"tokens"`
	Requests       int    `json:"requests"`
	ApiKey         int    `json:"apiKey"`
	TargetEndpoint string `json:"targetEndpoint"`
}

// * clientID: A unique identifier for each client.
// * tokens: The number of tokens the client wishes to reserve.
// * requests: The number of requests the client intends to make.
// * apiKey: The unique API key for the client.
// * targetEndpoint: The targeted OpenAI API endpoint of the request (specified in the configuration file).
