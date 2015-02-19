package brave

import (
	"reflect"
)

type map_ map[interface{}]interface{}

func mapData(src interface{}, m map_) interface{} {
	sv := reflect.ValueOf(src)
	dv := mapValue(sv, m)
	return dv.Interface()
}

func mapValue(src reflect.Value, m map_) reflect.Value {
	switch src.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:
		return reflect.ValueOf(m[src.Interface()])
	case reflect.Ptr:
		if src.IsNil() {
			return src
		}
		return mapValue(src.Elem(), m).Addr()
	case reflect.Slice:
		if src.IsNil() {
			return src
		}
		l := src.Len()
		d := reflect.MakeSlice(src.Type(), 0, l)
		for i := 0; i < l; i++ {
			d = reflect.Append(d, mapValue(src.Index(i), m))
		}
		return d
	case reflect.Map:
		if src.IsNil() {
			return src
		}
		d := reflect.MakeMap(src.Type())
		for _, key := range src.MapKeys() {
			d.SetMapIndex(key, mapValue(src.MapIndex(key), m))
		}
		return d
	case reflect.Struct:
		l := src.NumField()
		d := reflect.New(src.Type())
		for i := 0; i < l; i++ {
			if len(src.Type().Field(i).PkgPath) != 0 {
				continue
			}
			d.Elem().Field(i).Set(mapValue(src.Field(i), m))
		}
		return d.Elem()
	default:
		panic("Unsupported type")
	}
}
