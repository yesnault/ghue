package common

// Connection represents a connection on a bridge
type Connection struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Verbose  bool   `json:"verbose"`
}

// ErrorHUE represents an error from API
// http://www.developers.meethue.com/documentation/error-messages
type ErrorHUE struct {
	Error struct {
		Address     string `json:"address"`
		Description string `json:"description"`
		Type        int    `json:"type"`
	} `json:"error"`
}
