package comps

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Dropdown(items vecty.List, active bool) vecty.ComponentOrHTML {
	//Dropdown menu, show/hide based on menu state.
	//
	//Entering: "transition ease-out duration-100"
	//From: "transform opacity-0 scale-95"
	//To: "transform opacity-100 scale-100"
	//Leaving: "transition ease-in duration-75"
	//From: "transform opacity-100 scale-100"
	//To: "transform opacity-0 scale-95"
	//
	if !active {
		return nil
	}
	return elem.Div(
		vecty.Markup(
			vecty.Class("origin-top-right", "absolute", "right-0", "mt-2", "w-48", "rounded-md", "shadow-lg", "py-1", "bg-white", "ring-1", "ring-black", "ring-opacity-5", "focus:outline-none"),
			vecty.Attribute("role", "menu"),
			vecty.Attribute("aria-orientation", "vertical"),
			vecty.Attribute("aria-labelledby", "user-menu-button"),
			vecty.Attribute("tabindex", "-1"),
		),
		items,
	)
}

// Active: "bg-gray-100", Not Active: ""
// <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-0">{s:"Your Profile"}</a>
// <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-1">{s:"Settings"}</a>
// <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1" id="user-menu-item-2">{s:"Sign out"}</a>
func DropdownItemSimple(link, text string) vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.Attribute("href", link),
			vecty.Class("block", "px-4", "py-2", "text-sm", "text-gray-700"),
			vecty.Attribute("role", "menuitem"),
			vecty.Attribute("tabindex", "-1"),
			vecty.Attribute("id", "user-menu-item-2"),
		),
		vecty.Text(text),
	)
}
