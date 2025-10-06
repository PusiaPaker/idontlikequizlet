package tmpl

import (
	"html/template"
	"time"
)

var (
	Base *template.Template
	Home *template.Template
	Deck *template.Template
	Edit *template.Template
)

func MustInit() {
	// create base template with FuncMap BEFORE parsing
	Base = template.New("base").Funcs(template.FuncMap{
		"now": time.Now, // adds the now() function
	})

	// parse base + partials
	template.Must(Base.ParseFiles("web/templates/base.html"))
	template.Must(Base.ParseGlob("web/templates/partials/*.html"))

	// clone for each page
	Home = template.Must(Base.Clone())
	template.Must(Home.ParseFiles("web/templates/home.html"))

	Deck = template.Must(Base.Clone())
	template.Must(Deck.ParseFiles("web/templates/deck.html"))

	Edit = template.Must(Base.Clone())
	template.Must(Edit.ParseFiles("web/templates/edit.html"))
}
