package common

type EventType int64

const (
	Press EventType = iota
	Release
)

type BrowserEvent struct {
	Origin string    // ip
	Key    string    // which key?
	Typ    EventType // press or release?
}

type GameEvent struct {
	Key string
	Typ EventType
}
