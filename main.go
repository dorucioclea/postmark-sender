package main

import "github.com/paysuper/postmark-sender/internal"

func main() {
	app := internal.NewApplication()

	defer app.Stop()
	app.Run()
}
