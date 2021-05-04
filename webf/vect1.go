package webf

import (
	"encoding/json"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/mdev5000/vectest2/ajax"
	ajax2 "github.com/mdev5000/vectest2/webf/ajax"
	"github.com/mdev5000/vectest2/webf/comps"
	"github.com/mdev5000/vectest2/webf/router"
)

func Main() error {
	ajax := ajax2.NewAjaxApi("/api/first")
	comp := &PageView{
		ajax:   ajax,
		text:   "this is some text",
		navbar: comps.NewNavBar(),
	}
	compW := &PageWrapper{child: comp}

	vecty.AddStylesheet("https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css")

	r := router.New()
	r.Add("/main/first", compW, func() {
		comp.routeString = "first"
	})
	r.Add("/main/second", compW, func() {
		comp.routeString = "second"
	})

	comp.router = &r

	r.Listen()
	return nil
}

type PageWrapper struct {
	vecty.Core
	child vecty.Component
}

func (p *PageWrapper) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("first", "second"),
		),
		elem.Div(
			comps.NewNavBar(),
			elem.Div(
				vecty.Markup(
					vecty.Class("container", "max-w-7xl", "mx-auto", "py-6", "sm:px-6", "lg:px-8"),
				),
				p.child,
			),
		),
	)
}

type PageView struct {
	vecty.Core
	text        string
	ajax        ajax2.AjaxRequest
	navbar      *comps.Navbar
	router      *router.Router
	routeString string
}

func (p *PageView) DoAjax() {
	rq, err := json.Marshal([]byte("{}"))
	if err != nil {
		panic(err)
	}
	p.ajax.Request(rq, func(result string) {
		u := &ajax.User{}
		if err := json.Unmarshal([]byte(result), u); err != nil {
			panic(err)
		}
		p.text = u.Name
		vecty.Rerender(p)
	})
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
			elem.Div(
				vecty.Text("page: "),
				vecty.Text(p.routeString),
			),
			vecty.Text(p.text),
		),
		elem.Div(
			elem.Button(
				vecty.Markup(
					event.Click(func(v *vecty.Event) {
						p.router.Push("/main/second")
					}),
				),
				vecty.Text("stuff"),
			),
		),
		elem.Div(
			elem.Button(
				vecty.Markup(
					event.Click(func(v *vecty.Event) {
						p.DoAjax()
					}),
				),
				vecty.Text("click me"),
			),
		),
	)
}