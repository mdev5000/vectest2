package ajax

import (
	"syscall/js"
)

type AjaxApi struct {
	url string
}

type AjaxRequest interface {
	Request([]byte, func(string))
}

func NewAjaxApi(url string) *AjaxApi {
	return &AjaxApi{
		url: url,
	}
}

func (a *AjaxApi) Request(json []byte, fn func(result string)) {
	xHttp := js.Global().Get("XMLHttpRequest").New()
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if this.Get("readyState").Int() != 4 {
			return js.Null()
		}
		s := this.Get("responseText").String()
		fn(s)
		return js.Null()
	})
	xHttp.Set("onreadystatechange", f)
	xHttp.Call("open", "POST", a.url)
	request := js.ValueOf(string(json))
	xHttp.Call("send", request)
}
