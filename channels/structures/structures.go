package structures

type EventData struct {
	Id        int
	Operation string
}

type Event struct {
	Type      string
	EventData EventData
}