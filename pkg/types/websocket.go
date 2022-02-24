package types

import "encoding/json"

// WebsocketRequest defines a request sent over the websocket connection
type WebsocketRequest struct {
	Command     string      `json:"command"`
	Ack         bool        `json:"ack"`
	Origin      string      `json:"origin"`
	Destination string      `json:"destination"`
	RequestID   string      `json:"request_id"`
	Data        interface{} `json:"data"`
}

// WebsocketResponse defines the response structure received over the websocket connection
type WebsocketResponse struct {
	Command     string          `json:"command"`
	Ack         bool            `json:"ack"`
	Origin      string          `json:"origin"`
	Destination string          `json:"destination"`
	RequestID   string          `json:"request_id"`
	Data        json.RawMessage `json:"data"`
}

// WebsocketSubscription is the Data for a new subscribe request
type WebsocketSubscription struct {
	Service string `json:"service"`
}
