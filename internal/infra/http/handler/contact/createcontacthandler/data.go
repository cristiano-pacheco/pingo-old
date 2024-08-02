package createcontacthandler

type input struct {
	Name         string `json:"name"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsEnabled    bool   `json:"is_enabled"`
}

type output struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsEnabled    bool   `json:"is_enabled"`
}
