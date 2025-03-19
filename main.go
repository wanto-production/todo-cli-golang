package main

import (
	"fmt"

	"wanto/cli_app/cmd"
	"wanto/cli_app/model"
)

func main() {
	err := model.Init()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	model.Database.AutoMigrate(&model.Todos{})
	cmd.Exec()
}
