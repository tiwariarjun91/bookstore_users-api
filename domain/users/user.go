package users

type User struct{
	Id	int64	`json:"id"` // as this is int pass just a number in the postman request
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Email	string	`json:"email"`
	DateCreated	string	`json:"date_created"`

}