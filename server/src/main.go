package main

import "laynas/server/src/server"

func main() {
	svc := server.Create()

	err := svc.Run(":20000")
	if err != nil {
		panic(err.Error())
	}
}
