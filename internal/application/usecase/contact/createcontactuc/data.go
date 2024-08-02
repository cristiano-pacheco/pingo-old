package createcontactuc

type Input struct {
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
