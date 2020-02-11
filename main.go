package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func main() {
	js.Global.Get("console").Call("log", "Hey")

	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	portraitContainer := &Portraits{}
	return elem.Body(
		&Button{
			onClickHandler: portraitContainer.add,
		},
		portraitContainer,
	)
}

type Portraits struct {
	vecty.Core

	portraitURLs []string
}

func (p *Portraits) add() {
	fmt.Println("Adding a new portrait")
	newURL := "https://www.thispersondoesnotexist.com/image?" + strconv.Itoa(int(time.Now().Unix()))
	p.portraitURLs = append(p.portraitURLs, newURL)
	vecty.Rerender(p)
}

func (p *Portraits) Render() vecty.ComponentOrHTML {
	var portraitIMGs []vecty.MarkupOrChild
	for _, portraitURL := range p.portraitURLs {
		portraitIMGs = append(portraitIMGs, elem.Image(
			vecty.Markup(vecty.Attribute("src", portraitURL)),
			vecty.Markup(vecty.Attribute("width", 250)),
		))
	}
	return elem.Div(portraitIMGs...)
}

type Button struct {
	vecty.Core

	onClickHandler func()
}

func (b *Button) onClick(event *vecty.Event) {
	fmt.Println("HERE")
	b.onClickHandler()
}

func (b *Button) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Text("Click it"),
		vecty.Markup(event.Click(b.onClick).PreventDefault()),
	)
}
