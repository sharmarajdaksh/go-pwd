package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sharmarajdaksh/go-pwd/menu"

	"github.com/sharmarajdaksh/go-pwd/config"
	"github.com/sharmarajdaksh/go-pwd/db"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("error: failed to load config: ", err)
		runtime.Goexit()
	}

	if err := db.Initialize(); err != nil {
		fmt.Println("error: failed to initialize database: ", err)
		runtime.Goexit()
	}
}

func main() {
	defer os.Exit(0)

	menu.RunProgram()
}
