package findcontactlistuc

type Input struct {
	UserID string
}

type Output struct {
	Items []*contact
}

type contact struct {
	ID           string
	UserID       string
	Name         string
	ContactType  string
	ContactValue string
	IsEnabled    bool
}
