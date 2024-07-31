package updatepasswordhandler

type input struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
