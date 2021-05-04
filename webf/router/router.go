package router

import (
	"fmt"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"syscall/js"
)

type routerComp struct {
	vecty.Core
	activeComp vecty.Component
}

func (p *routerComp) Render() vecty.ComponentOrHTML {
	return elem.Body(
		p.activeComp,
	)
}

type route struct {
	Comp vecty.Component
	Init func()
}

type Router struct {
	routes map[string]route
	routerComp *routerComp
}

func New() Router {
	return Router{
		routes: map[string]route{},
		routerComp: &routerComp{},
	}
}

func (r *Router) Add(path string, comp vecty.Component, fn func()) {
	r.routes[path] = route{
		Comp: comp,
		Init: fn,
	}
}

func (*Router) Path() string {
	loc := js.Global().Get("location")
	path := loc.Get("pathname")
	return path.String()
}

func (r *Router) updateState(rerender bool) {
	path := r.Path()
	route, found := r.routes[path]
	if !found {
		fmt.Printf("not found: %s\n", path)
	}
	if route.Init != nil {
		route.Init()
	}
	r.routerComp.activeComp = route.Comp
	//vecty.RenderBody(route.Comp)
	if rerender {
		vecty.Rerender(r.routerComp)
	}
}

func (r *Router) Push(uri string) {
	history := js.Global().Get("history")
	history.Call("pushState", uri, "Title", uri)
	r.updateState(true)
}

func (r *Router) Listen() {
	rerender := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		r.updateState(true)
		return js.Null()
	})
	js.Global().Call("addEventListener", "load", rerender)
	window := js.Global().Get("window")
	window.Set("onpopstate", rerender)
	r.updateState(false)
	vecty.RenderBody(r.routerComp)
}
