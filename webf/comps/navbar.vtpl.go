package comps

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/mdev5000/vectest2/webf/comps/navbar"
)

type Navbar struct {
	vecty.Core
	active bool
}

func NewNavBar() *Navbar {
	return &Navbar{
		active: false,
	}
}

func (n *Navbar) toggleActive(e *vecty.Event) {
	n.active = !n.active
	vecty.Rerender(n)
}

func (n *Navbar) Render() vecty.ComponentOrHTML {
	dropdown := Dropdown(List(
		DropdownItemSimple("#", "Your profile"),
		DropdownItemSimple("#", "Settings"),
		DropdownItemSimple("#", "Sign out"),
	), n.active)

	return elem.Div(
		elem.Div(
			vecty.Markup(
				vecty.Class("bg-gray-800"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("max-w-7xl", "mx-auto", "px-4", "sm:px-6", "lg:px-8"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("flex", "items-center", "justify-between", "h-16"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("flex", "items-center"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("flex-shrink-0"),
							),
							elem.Image(
								vecty.Markup(
									vecty.Class("h-8", "w-8"),
									vecty.Attribute("src", "https://tailwindui.com/img/logos/workflow-mark-indigo-500.svg"),
									vecty.Attribute("alt", "Workflow"),
								),
							),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("hidden", "md:block"),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("ml-10", "flex", "items-baseline", "space-x-4"),
								),
								navbar.Link("#", "Dashboard"),
								elem.Anchor(
									vecty.Markup(
										vecty.Attribute("href", "#"),
										vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "px-3", "py-2", "rounded-md", "text-sm", "font-medium"),
									),
									vecty.Text("Team"),
								),
								elem.Anchor(
									vecty.Markup(
										vecty.Attribute("href", "#"),
										vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "px-3", "py-2", "rounded-md", "text-sm", "font-medium"),
									),
									vecty.Text("Projects"),
								),
								elem.Anchor(
									vecty.Markup(
										vecty.Attribute("href", "#"),
										vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "px-3", "py-2", "rounded-md", "text-sm", "font-medium"),
									),
									vecty.Text("Calendar"),
								),
								elem.Anchor(
									vecty.Markup(
										vecty.Attribute("href", "#"),
										vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "px-3", "py-2", "rounded-md", "text-sm", "font-medium"),
									),
									vecty.Text("Reports"),
								),
							),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("hidden", "md:block"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("ml-4", "flex", "items-center", "md:ml-6"),
							),
							elem.Button(
								vecty.Markup(
									vecty.Class("bg-gray-800", "p-1", "rounded-full", "text-gray-400", "hover:text-white", "focus:outline-none", "focus:ring-2", "focus:ring-offset-2", "focus:ring-offset-gray-800", "focus:ring-white"),
								),
								elem.Span(
									vecty.Markup(
										vecty.Class("sr-only"),
									),
									vecty.Text("View notifications"),
								),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("ml-3", "relative"),
								),
								elem.Div(
									elem.Button(
										vecty.Markup(
											vecty.Attribute("type", "button"),
											vecty.Class("max-w-xs", "bg-gray-800", "rounded-full", "flex", "items-center", "text-sm", "focus:outline-none", "focus:ring-2", "focus:ring-offset-2", "focus:ring-offset-gray-800", "focus:ring-white"),
											vecty.Attribute("id", "user-menu-button"),
											vecty.Attribute("aria-expanded", "false"),
											vecty.Attribute("aria-haspopup", "true"),
											event.Click(n.toggleActive),
										),
										elem.Span(
											vecty.Markup(
												vecty.Class("sr-only"),
											),
											vecty.Text("Open user menu"),
										),
										elem.Image(
											vecty.Markup(
												vecty.Class("h-8", "w-8", "rounded-full"),
												vecty.Attribute("src", "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"),
												vecty.Attribute("alt", ""),
											),
										),
									),
								),
								dropdown,
							),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("-mr-2", "flex", "md:hidden"),
						),
						elem.Button(
							vecty.Markup(
								vecty.Attribute("type", "button"),
								vecty.Class("bg-gray-800", "inline-flex", "items-center", "justify-center", "p-2", "rounded-md", "text-gray-400", "hover:text-white", "hover:bg-gray-700", "focus:outline-none", "focus:ring-2", "focus:ring-offset-2", "focus:ring-offset-gray-800", "focus:ring-white"),
								vecty.Attribute("aria-controls", "mobile-menu"),
								vecty.Attribute("aria-expanded", "false"),
							),
							elem.Span(
								vecty.Markup(
									vecty.Class("sr-only"),
								),
								vecty.Text("Open main menu"),
							),
						),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("md:hidden"),
					vecty.Attribute("id", "mobile-menu"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("px-2", "pt-2", "pb-3", "space-y-1", "sm:px-3"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Attribute("href", "#"),
							vecty.Class("bg-gray-900", "text-white", "block", "px-3", "py-2", "rounded-md", "text-base", "font-medium"),
						),
						vecty.Text("Dashboard"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Attribute("href", "#"),
							vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "block", "px-3", "py-2", "rounded-md", "text-base", "font-medium"),
						),
						vecty.Text("Team"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Attribute("href", "#"),
							vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "block", "px-3", "py-2", "rounded-md", "text-base", "font-medium"),
						),
						vecty.Text("Projects"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Attribute("href", "#"),
							vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "block", "px-3", "py-2", "rounded-md", "text-base", "font-medium"),
						),
						vecty.Text("Calendar"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Attribute("href", "#"),
							vecty.Class("text-gray-300", "hover:bg-gray-700", "hover:text-white", "block", "px-3", "py-2", "rounded-md", "text-base", "font-medium"),
						),
						vecty.Text("Reports"),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("pt-4", "pb-3", "border-t", "border-gray-700"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("flex", "items-center", "px-5"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("flex-shrink-0"),
							),
							elem.Image(
								vecty.Markup(
									vecty.Class("h-10", "w-10", "rounded-full"),
									vecty.Attribute("src", "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"),
									vecty.Attribute("alt", ""),
								),
							),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("ml-3"),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("text-base", "font-medium", "leading-none", "text-white"),
								),
								vecty.Text("Tom Cook"),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("text-sm", "font-medium", "leading-none", "text-gray-400"),
								),
								vecty.Text("tom@example.com"),
							),
						),
						elem.Button(
							vecty.Markup(
								vecty.Class("ml-auto", "bg-gray-800", "flex-shrink-0", "p-1", "rounded-full", "text-gray-400", "hover:text-white", "focus:outline-none", "focus:ring-2", "focus:ring-offset-2", "focus:ring-offset-gray-800", "focus:ring-white"),
							),
							elem.Span(
								vecty.Markup(
									vecty.Class("sr-only"),
								),
								vecty.Text("View notifications"),
							),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("mt-3", "px-2", "space-y-1"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Attribute("href", "#"),
								vecty.Class("block", "px-3", "py-2", "rounded-md", "text-base", "font-medium", "text-gray-400", "hover:text-white", "hover:bg-gray-700"),
							),
							vecty.Text("Your Profile"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Attribute("href", "#"),
								vecty.Class("block", "px-3", "py-2", "rounded-md", "text-base", "font-medium", "text-gray-400", "hover:text-white", "hover:bg-gray-700"),
							),
							vecty.Text("Settings"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Attribute("href", "#"),
								vecty.Class("block", "px-3", "py-2", "rounded-md", "text-base", "font-medium", "text-gray-400", "hover:text-white", "hover:bg-gray-700"),
							),
							vecty.Text("Sign out"),
						),
					),
				),
			),
		),
		elem.Header(
			vecty.Markup(
				vecty.Class("bg-white", "shadow"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("max-w-7xl", "mx-auto", "py-6", "px-4", "sm:px-6", "lg:px-8"),
				),
				elem.Heading1(
					vecty.Markup(
						vecty.Class("text-3xl", "font-bold", "text-gray-900"),
					),
					vecty.Text("Dashboard"),
				),
			),
		),
		elem.Main(
			elem.Div(
				vecty.Markup(
					vecty.Class("max-w-7xl", "mx-auto", "py-6", "sm:px-6", "lg:px-8"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("px-4", "py-6", "sm:px-0"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("border-4", "border-dashed", "border-gray-200", "rounded-lg", "h-96"),
						),
					),
				),
			),
		),
	)
}
