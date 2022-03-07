package errors

type Error struct {
	Err  string `json:"err,omitempty"`
	Code string `json:"ECODE,omitempty"`
}
