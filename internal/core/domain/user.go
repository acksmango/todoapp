package domain

type User struct {
	ID      int
	version int

	FullName    string
	PhoneNumber *string
}
