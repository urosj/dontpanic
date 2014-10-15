package dontpanic

import (
	"math/rand"
	"reflect"

	"gopkg.in/errgo.v1"
)

func Fuzz(iface interface{}) (interface{}, error) {
	source := reflect.ValueOf(iface)
	fuzzed := reflect.New(source.Type()).Elem()
	err := fuzzWalker(fuzzed, source)
	return fuzzed.Interface(), err
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@##$^&*()=+-_1234567890")

func randStringSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func fuzzWalker(fuzzed, source reflect.Value) error {
	switch source.Kind() {

	case reflect.Interface:
		sourceValue := source.Elem()
		fuzzedValue := reflect.New(sourceValue.Type()).Elem()
		err := fuzzWalker(fuzzedValue, sourceValue)
		if err != nil {
			return errgo.Notef(err, "error in interface")
		}
		fuzzed.Set(fuzzedValue)

	case reflect.Struct:
		for i := 0; i < source.NumField(); i += 1 {
			err := fuzzWalker(fuzzed.Field(i), source.Field(i))
			if err != nil {
				return errgo.Notef(err, "error in struct")
			}
		}

	case reflect.Map:
		fuzzed.Set(reflect.MakeMap(source.Type()))
		for _, key := range source.MapKeys() {
			sourceValue := source.MapIndex(key)
			fuzzedValue := reflect.New(sourceValue.Type()).Elem()
			err := fuzzWalker(fuzzedValue, sourceValue)
			if err != nil {
				return errgo.Notef(err, "error in map")
			}
			fuzzed.SetMapIndex(key, fuzzedValue)
		}

	case reflect.Ptr:
		sourceValue := source.Elem()
		if !sourceValue.IsValid() {
			return errgo.New("source value is not valid")
		}
		fuzzed.Set(reflect.New(sourceValue.Type()))
		err := fuzzWalker(fuzzed.Elem(), sourceValue)
		if err != nil {
			return errgo.Notef(err, "error in pointer")
		}

	case reflect.Slice:
		fuzzed.Set(reflect.MakeSlice(source.Type(), source.Len(), source.Cap()))
		for i := 0; i < source.Len(); i += 1 {
			err := fuzzWalker(fuzzed.Index(i), source.Index(i))
			if err != nil {
				return errgo.Notef(err, "error in interface")
			}
		}

	case reflect.Uint8:
		newInt := uint8(rand.Intn(1 << 8))
		fuzzed.SetUint(uint64(newInt))
	case reflect.Uint16:
		newInt := uint16(rand.Intn(1 << 16))
		fuzzed.SetUint(uint64(newInt))
	case reflect.Uint32:
		newInt := uint32(rand.Int63n(1 << 32))
		fuzzed.SetUint(uint64(newInt))
	case reflect.Uint64:
		newInt := uint64((rand.Int63() << 1)) + uint64(rand.Intn(2))
		fuzzed.SetUint(uint64(newInt))

	case reflect.Int8:
		newInt := int8(rand.Int())
		fuzzed.SetInt(int64(newInt))
	case reflect.Int16:
		newInt := int16(rand.Int())
		fuzzed.SetInt(int64(newInt))
	case reflect.Int32:
		newInt := int32(rand.Int63())
		fuzzed.SetInt(int64(newInt))
	case reflect.Int64:
		newInt := int64(rand.Int63())
		fuzzed.SetInt(int64(newInt))

	case reflect.String:
		newInt := rand.Intn(1 << 8)
		translatedString := randStringSeq(newInt)
		fuzzed.SetString(translatedString)

	default:
		fuzzed.Set(source)
	}

	return nil
}
