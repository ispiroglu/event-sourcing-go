package command

type CreateCustomerCommand struct {
	BaseCommand
	FirstName string
	LastName  string
}
