package effective



/* omitempty的用法*/
type Tips struct {
	name string
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Tips     *Tips  `json:"tips,omitempty"`
	// many more fields…
}




