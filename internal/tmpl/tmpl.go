// in package tmpl (or wherever you init templates)
package tmpl

import "html/template"

var (
    Base *template.Template
    Home *template.Template
    Deck *template.Template
)

func MustInit() {
    // 1) Load base + partials once
    Base = template.Must(template.New("base").ParseFiles(
        "web/templates/base.html",
    ))
    template.Must(Base.ParseGlob("web/templates/partials/*.html"))

    // 2) Clone for HOME and add home page
    Home = template.Must(Base.Clone())
    template.Must(Home.ParseFiles("web/templates/home.html"))

    // 3) Clone for DECK and add deck page
    Deck = template.Must(Base.Clone())
    template.Must(Deck.ParseFiles("web/templates/deck.html"))
}
