package main

import (
	"collab/microscope/components"
	"collab/microscope/models"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jeminstall/gopherjs-vue"
)

type controller struct {
	*js.Object
}

func main() {
	d := models.MockDeck()

	components.RegisterAll()

	// create the VueJS viewModel using a struct pointer
	vue.New("#app", d)
}
