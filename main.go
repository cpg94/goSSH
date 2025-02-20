package main

import (
	"fmt"
	"github.com/cpg94/goSSH/jsonutils"
)

func main() {
	sessions := jsonutils.Read()

	fmt.Println(sessions.Sessions[0].Name)
}
