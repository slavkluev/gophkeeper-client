package main

import (
	"client/cmd"
	_ "client/cmd/account"
	_ "client/cmd/text"
)

func main() {
	cmd.Execute()
}
