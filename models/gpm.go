package models

type GPM struct {
	ID       int    `json:"id"`
	Website  string `json:"website"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Notes    string `json:"notes"`
}

type GPMs []GPM

type AllRows struct {
	Status string `json:"status"`
	Rows   GPMs   `json:"rows"`
}

type ReqInsertGPM struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Notes    string `json:"notes"`
}
