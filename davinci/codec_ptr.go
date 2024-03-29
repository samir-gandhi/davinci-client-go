package davinci

import (
	"fmt"
	"reflect"
)

var _ ValueEncoder = &PtrCodec{}
var _ ValueDecoder = &PtrCodec{}

type PtrCodec struct {
	dCtx *DecoderContext
	eCtx *EncoderContext
}

func NewPtrDecoder(dCtx *DecoderContext) ValueDecoder {
	return PtrCodec{
		dCtx: dCtx,
	}
}

func NewPtrEncoder(eCtx *EncoderContext) ValueEncoder {
	return PtrCodec{
		eCtx: eCtx,
	}
}

func (PtrCodec) String() string {
	return "davinci.PtrCodec"
}

func (d PtrCodec) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Ptr {
		return fmt.Errorf("invalid pointer value to decode")
	}

	typ := v.Type()

	if v.IsNil() {
		v.Set(reflect.New(typ.Elem()))
	}

	// Decode the value into the struct field
	err := d.dCtx.Decode(data, v.Elem().Addr().Interface())
	if err != nil {
		return err
	}

	val := v.Elem()

	if val.Kind() == reflect.Struct {
		// Check if the struct is zero valued, if so, set the pointer to nil
		if reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface()) {
			v.Set(reflect.Zero(typ))
		}
	}

	return nil
}

func (d PtrCodec) EncodeValue(v reflect.Value) ([]byte, error) {
	if !v.IsValid() || v.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("invalid pointer value to encode")
	}

	val := v.Elem()

	if !val.IsValid() {
		return []byte{}, nil
	}

	// Encode the value
	return d.eCtx.Encode(val.Interface())
}
