package domain

type Customer struct {
	id          string
	firstName   string
	lastName    string
	phoneNumber string
	email       string
	restricted  bool
	birthYear   int
	gender      Gender
}
