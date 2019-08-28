// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package base

import (
	errors "errors"
	fmt "fmt"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type Message struct {
	Body string `json:"body,required"`
}

// ToWire translates a Message struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Message) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Body), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Message struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Message struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Message
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Message) FromWire(w wire.Value) error {
	var err error

	bodyIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Body, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				bodyIsSet = true
			}
		}
	}

	if !bodyIsSet {
		return errors.New("field Body of Message is required")
	}

	return nil
}

// String returns a readable string representation of a Message
// struct.
func (v *Message) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Body: %v", v.Body)
	i++

	return fmt.Sprintf("Message{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Message match the
// provided Message.
//
// This function performs a deep comparison.
func (v *Message) Equals(rhs *Message) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Body == rhs.Body) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Message.
func (v *Message) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("body", v.Body)
	return err
}

// GetBody returns the value of Body if it is set or its
// zero value if it is unset.
func (v *Message) GetBody() (o string) {
	if v != nil {
		o = v.Body
	}
	return
}
