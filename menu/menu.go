package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sharmarajdaksh/go-pwd/db"
)

var programActions = []action{
	&quitAction{},
	&newPasswordAction{},
	&listAllAction{},
	&passwordByAppAction{},
	&passwordByEmailAction{},
	&passwordByUnameAction{},
}

// RunProgram starts the program menu
func RunProgram() {

	for true {
		ac, err := selectAction()
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = ac.execute()
		if err != nil {
			fmt.Println("\n", err)
		}

		awaitEnterInput()
	}

}

func selectAction() (action, error) {
	displayMenu()

	ch, err := strconv.Atoi(getUserInput("your choice"))
	if err != nil || ch >= len(programActions) {
		return nil, fmt.Errorf("Invalid input, please provide a valid integer choice")
	}

	ac := programActions[ch]

	return ac, nil
}

func displayMenu() {
	fmt.Println()
	fmt.Println("--------------INPUT YOUR CHOICE--------------")
	for i, a := range programActions {
		fmt.Println(i, ". ", a.displayString())
	}
	fmt.Println()
}

func getUserInput(inputName string) string {
	fmt.Printf("Please input %s: ", inputName)
	var ip string
	fmt.Scanln(&ip)
	fmt.Println()

	ip = strings.TrimSpace(ip)
	return ip
}

func awaitEnterInput() {
	fmt.Println("\nPress enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func displayRecords(ps []db.Password) {
	for i, p := range ps {
		fmt.Println()
		fmt.Println(i, " -------------------------------------------")

		fields := []struct {
			key   string
			value string
		}{
			{"App", p.App},
			{"URL", p.URL},
			{"Username", p.Username},
			{"Email", p.Email},
			{"Password", p.Password},
		}

		for _, f := range fields {
			fmt.Println(f.key, ": ", stringOrNone(f.value))
		}

	}
}

func stringOrNone(s string) string {
	if s == "" {
		return "<NONE>"
	}

	return s
}
