package entities

type User struct {
	Id      int64  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Pwd     string `json:"pwd,omitempty"`
	Parents []User `json:"parents,omitempty"`
}
