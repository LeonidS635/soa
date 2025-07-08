package events

type Event interface {
	Marshal() ([]byte, error)
	Unmarshal(b []byte) error
}
