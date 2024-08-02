package updatecontactuc

type Input struct {
	ID           string
	UserID       string
	Name         string
	ContactType  string
	ContactValue string
	IsEnabled    bool
}

type Output struct {
	ID           string
	UserID       string
	Name         string
	ContactType  string
	ContactValue string
	IsEnabled    bool
}
