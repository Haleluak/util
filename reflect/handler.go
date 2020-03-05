package reflect

import (
	"reflect"
)
type subscriber struct {
	rcvr       reflect.Value
	typ        reflect.Type
	handlers   []*handler
}

type handler struct {
	method  reflect.Value
	reqType reflect.Type
	ctxType reflect.Type
}

func newSubscriber(sub interface{})  *subscriber{
	var handlers []*handler
	if typ := reflect.TypeOf(sub); typ.Kind() == reflect.Func {
		h := &handler{
			method: reflect.ValueOf(sub),
		}
		switch typ.NumIn() {
		case 1:
			h.reqType = typ.In(0)
		case 2:
			h.ctxType = typ.In(0)
			h.reqType = typ.In(1)
		}

		handlers = append(handlers, h)
	}else {
		for m := 0; m < typ.NumMethod(); m++ {
			method := typ.Method(m)
			h := &handler{
				method: method.Func,
			}
			switch method.Type.NumIn() {
			case 2:
				h.reqType = method.Type.In(1)
			case 3:
				h.ctxType = method.Type.In(1)
				h.reqType = method.Type.In(2)
			}

			handlers = append(handlers, h)
		}
	}
	return &subscriber{
		rcvr:       reflect.ValueOf(sub),
		typ:        reflect.TypeOf(sub),
		handlers:   handlers,
	}
}

func createSubHandler(sb *subscriber) (err error){
	for i := 0; i < len(sb.handlers); i++ {
		handler := sb.handlers[i]
		var isVal bool
		var req reflect.Value

		if handler.reqType.Kind() == reflect.Ptr {
			req = reflect.New(handler.reqType.Elem())
		} else {
			req = reflect.New(handler.reqType)
			isVal = true
		}
		if isVal {
			req = req.Elem()
		}

		var vals []reflect.Value
		if sb.typ.Kind() != reflect.Func {
			vals = append(vals, sb.rcvr)
		}
		if handler.ctxType != nil {
			vals = append(vals, reflect.ValueOf("abc"))
		}
		vals = append(vals, reflect.ValueOf(req.Interface()))
		returnValues := handler.method.Call(vals)
		if rerr := returnValues[0].Interface(); rerr != nil {
			return rerr.(error)
		}
	}

	return nil
}