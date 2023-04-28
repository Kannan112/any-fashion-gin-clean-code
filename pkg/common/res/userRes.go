package res

type UserData struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
type UserDetails struct {
	Name              string
	Email             string
	Mobile            string
	IsBlocked         bool
	BlockedAt         string `json:",omitempty"`
	BlockedBy         uint   `json:",omitempty"`
	ReasonForBlocking string `json:",omitempty"`
}
