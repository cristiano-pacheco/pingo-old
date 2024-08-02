package updatecontacthandler

type input struct {
	Name         string `json:"name"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsEnabled    bool   `json:"is_enabled"`
}
