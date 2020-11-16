package menu

import (
	"fmt"
	"runtime"

	"github.com/sharmarajdaksh/go-pwd/db"
)

type action interface {
	displayString() string
	execute() error
}

type quitAction struct{}

type listAllAction struct{}

type newPasswordAction struct{}

type passwordByAppAction struct{}

type passwordByEmailAction struct{}

type passwordByUnameAction struct{}

func (a *quitAction) displayString() string {
	return "Quit"
}

func (a *quitAction) execute() error {
	fmt.Println("Exiting program...")
	runtime.Goexit()
	return nil
}

func (a *listAllAction) displayString() string {
	return "List all password entries"
}

func (a *listAllAction) execute() error {

	ps, err := db.GetAllRecords()
	if err != nil {
		return fmt.Errorf("An error occurred while fetching database entries")
	}

	displayRecords(ps)

	return nil

}

func (a *newPasswordAction) displayString() string {
	return "Add a new password entry"
}

func (a *newPasswordAction) execute() error {
	app := getUserInput("the name of the app")
	if app == "" {
		return fmt.Errorf("App name cannot be blank")
	}

	url := getUserInput("the app url (optional)")
	username := getUserInput("the associated username (optional)")
	email := getUserInput("the associated email (optional)")
	if email == "" && username == "" {
		return fmt.Errorf("You must provide an email or a password for a password entry")
	}

	password := getUserInput("the password to save")
	if password == "" {
		return fmt.Errorf("Password cannot be blank")
	}

	p := db.NewPassword(
		app,
		url,
		username,
		email,
		password,
	)

	err := p.Save()
	if err != nil {
		return fmt.Errorf("An error occurred while creating the new password entry: %v", err)
	}

	fmt.Println("Password entry created successfully")

	return nil
}

func (a *passwordByAppAction) displayString() string {
	return "Find password entry by app"
}

func (a *passwordByAppAction) execute() error {
	app := getUserInput("app name")
	if app == "" {
		return fmt.Errorf("App name cannot be blank")
	}

	ps, err := db.GetRecordsByApp(app)
	if err != nil {
		return fmt.Errorf("An error occurred while fetching database entries")
	}

	displayRecords(ps)

	return nil
}

func (a *passwordByEmailAction) displayString() string {
	return "Find password entry by email"
}

func (a *passwordByEmailAction) execute() error {
	email := getUserInput("email")
	if email == "" {
		return fmt.Errorf("Email cannot be blank for this option")
	}

	ps, err := db.GetRecordsByEmail(email)
	if err != nil {
		return fmt.Errorf("An error occurred while fetching database entries")
	}

	displayRecords(ps)

	return nil
}

func (a *passwordByUnameAction) displayString() string {
	return "Find password entry by username"
}

func (a *passwordByUnameAction) execute() error {
	username := getUserInput("username")
	if username == "" {
		return fmt.Errorf("Username cannot be blank for this option")
	}

	ps, err := db.GetRecordsByUsername(username)
	if err != nil {
		return fmt.Errorf("An error occurred while fetching database entries")
	}

	displayRecords(ps)

	return nil
}
