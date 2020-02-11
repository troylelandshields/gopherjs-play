package main

import (
	"strconv"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	js.Global.Get("console").Call("log", "Hey")
	document := dom.GetWindow().Document()

	btn := document.GetElementByID("btn")
	btn.AddEventListener("click", false, func(event dom.Event) {
		var x string
		img := document.CreateElement("img").(*dom.HTMLImageElement)
		img.Src = "https://www.thispersondoesnotexist.com/image?" + strconv.Itoa(int(time.Now().Unix()))
		img.Height = 250

		panic("Oops")

		document.GetElementByID("container").AppendChild(img)
	})
}
