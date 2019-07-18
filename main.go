package main

import "github.com/paysuper/postmark-sender/internal"

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Stop()
	app.Run()
}
