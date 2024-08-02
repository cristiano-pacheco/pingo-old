package findcontactlisthandler

type output struct {
	Items []*contact `json:"items"`
}

type contact struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsEnabled    bool   `json:"is_enabled"`
}
