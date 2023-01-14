package book

type Book struct {
	UUID   string `json:"uuid,omitempty"`
	Name   string `json:"name,omitempty"`
	Year   int    `json:"year,omitempty"`
	Author string `json:"author,omitempty"`
	Busy   bool   `json:"busy,omitempty"`
	Owner  string `json:"owner,omitempty"`
}