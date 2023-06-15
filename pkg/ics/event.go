package ics

const (
	// EventRegistered ...
	EventRegistered = "registered"
	// EventUnregistered ...
	EventUnregistered = "unregistered"
)

// Event is a message for rabbitMQ
type Event struct {
	GameID int32  `json:"game_id,omitempty"`
	Event  string `json:"event,omitempty"`
}
