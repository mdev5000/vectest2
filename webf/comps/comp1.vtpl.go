package comps

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func (c *Comp1) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Text("first test and some mores"),
	)
}
