package brave

import (
	"fmt"
	"reflect"
)

type map_ struct {
	m      map[interface{}]interface{}
	strict bool
}

func newMap(strict bool) *map_ {
	return &map_{
		m:      make(map[interface{}]interface{}),
		strict: strict,
	}
}

func (m *map_) get(k interface{}) interface{} {
	v, ok := m.m[k]
	if ok {
		return v
	}
	if m.strict {
		return nil
	}
	return k
}

func (m *map_) has(k interface{}) bool {
	_, ok := m.m[k]
	return ok
}

func (m *map_) set(k, v interface{}) {
	m.m[k] = v
}

func mapData(src interface{}, m *map_) (dst interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			dst, err = nil, fmt.Errorf("Error in mapping data %v", r)
		}
	}()
	sv := reflect.ValueOf(src)
	dv := mapValue(sv, m)
	return dv.Interface(), nil
}

func mapValue(src reflect.Value, m *map_) reflect.Value {
	switch src.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String, reflect.Interface:
		return reflect.ValueOf(m.get(src.Interface()))
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
