package shared

import "time"

// Person represents the a person for our db.
type Person struct {
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Address   string     `json:"address"`
	Gender    string     `json:"gender"`
	Timestamp *time.Time `json:"timestamp"`
}

// EnsureTimeStampIsSet checks the timestamp. If it's null, sets it to current time
func (p *Person) EnsureTimeStampIsSet() {
	if p.Timestamp == nil {
		t := time.Now()
		p.Timestamp = &t
	}
}
