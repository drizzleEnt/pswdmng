package main

import "pswdmng/internal/app"

func main() {
	a := app.New(
		app.WhithStorePath(""),
	)
	a.Run()
}
