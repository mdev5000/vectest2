package comps

import (
	"github.com/hexops/vecty"
)

func List(comps ...vecty.ComponentOrHTML) vecty.List {
	l := make(vecty.List, len(comps))
	for i, comp := range comps {
		l[i] = comp
	}
	return l
}

