package jsutil

import "github.com/gopherjs/gopherjs/js"

func NewJSObject() *js.Object {
	return js.Global.Get("Object").New()
}
