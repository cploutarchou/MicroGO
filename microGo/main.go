package main

import "github.com/cploutarchou/rapiditas"

type application struct {
	App *rapiditas.Rapiditas
}

func main() {
	r := initApplication()
	r.App.ListenEndServe()
}
