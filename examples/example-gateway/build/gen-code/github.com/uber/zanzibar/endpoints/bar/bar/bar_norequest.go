// Code generated by thriftrw v1.0.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_NoRequest_Args struct{}

func (v *Bar_NoRequest_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_NoRequest_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *Bar_NoRequest_Args) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("Bar_NoRequest_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_NoRequest_Args) MethodName() string {
	return "noRequest"
}

func (v *Bar_NoRequest_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_NoRequest_Helper = struct {
	Args           func() *Bar_NoRequest_Args
	IsException    func(error) bool
	WrapResponse   func(*BarResponse, error) (*Bar_NoRequest_Result, error)
	UnwrapResponse func(*Bar_NoRequest_Result) (*BarResponse, error)
}{}

func init() {
	Bar_NoRequest_Helper.Args = func() *Bar_NoRequest_Args {
		return &Bar_NoRequest_Args{}
	}
	Bar_NoRequest_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}
	Bar_NoRequest_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_NoRequest_Result, error) {
		if err == nil {
			return &Bar_NoRequest_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_NoRequest_Result.BarException")
			}
			return &Bar_NoRequest_Result{BarException: e}, nil
		}
		return nil, err
	}
	Bar_NoRequest_Helper.UnwrapResponse = func(result *Bar_NoRequest_Result) (success *BarResponse, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Bar_NoRequest_Result struct {
	Success      *BarResponse  `json:"success,omitempty"`
	BarException *BarException `json:"barException,omitempty"`
}

func (v *Bar_NoRequest_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_NoRequest_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_NoRequest_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_NoRequest_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_NoRequest_Result) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}
	return fmt.Sprintf("Bar_NoRequest_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_NoRequest_Result) MethodName() string {
	return "noRequest"
}

func (v *Bar_NoRequest_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}