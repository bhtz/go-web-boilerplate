package models

/**
 * User structure.
 */
type User struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Github    string `json:"github"`
}
