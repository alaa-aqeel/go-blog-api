package helpers

import "strings"

type Map map[string][]string

func (e Map) Add(key string, err string) {
	e[strings.ToLower(key)] = append(e[key], err)
}

func (e Map) Get(key string) []string {
	return e[key]
}

func (e Map) Has(key string) bool {
	_, ok := e[strings.ToLower(key)]
	return ok
}

func (e Map) Keys() []string {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}
	return keys
}

func (e Map) Values() []string {
	values := make([]string, 0, len(e))
	for _, v := range e {
		values = append(values, v...)
	}
	return values
}

func (e Map) Empty() bool {
	return len(e.Values()) <= 0
}
