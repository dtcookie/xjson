package xjson

import (
	"encoding/json"
	"reflect"
)

type Properties map[string]json.RawMessage

func NewProperties(m map[string]json.RawMessage) Properties {
	props := Properties{}
	if len(m) > 0 {
		for k, v := range m {
			props[k] = v
		}
	}
	return props
}

func (p Properties) Unmarshal(key string, target interface{}) error {
	if v, found := p[key]; found {
		if err := json.Unmarshal(v, target); err != nil {
			return err
		}
		delete(p, key)
	}
	return nil
}

func (p Properties) Marshal(key string, v interface{}) error {
	isNil := (v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()))
	if isNil {
		return nil
	}
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Slice, reflect.Map:
		if reflect.ValueOf(v).Len() == 0 {
			return nil
		}
	default:
	}
	rawMessage, err := json.Marshal(v)
	if err != nil {
		return err
	}
	p[key] = rawMessage
	return nil
}
