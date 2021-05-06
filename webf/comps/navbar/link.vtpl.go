package navbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Link(url, name string) vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.Attribute("href", url),
			vecty.Class("bg-gray-900", "text-white", "px-3", "py-2", "rounded-md", "text-sm", "font-medium"),
		),
		vecty.Text(name),
	)
}
