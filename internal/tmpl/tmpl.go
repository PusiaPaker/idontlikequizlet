package tmpl

import (
	"html/template"
	"time"
	"fmt"
)

func dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: need even number of args")
	}
	m := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}

var (
	Base *template.Template
	Home *template.Template
	Deck *template.Template
	Edit *template.Template
)

func MustInit() {
	// create base template with FuncMap BEFORE parsing
	funcs := template.FuncMap{
		"now":  time.Now,
		"dict": func(values ...any) (map[string]any, error) { return dict(values...) },
	}

	// parse base + partials
	Base = template.New("base").Funcs(funcs)
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
