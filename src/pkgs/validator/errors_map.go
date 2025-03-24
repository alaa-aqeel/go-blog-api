package validator

import "strings"

type Errors map[string][]string

func (e Errors) Add(key string, err string) {
	e[strings.ToLower(key)] = append(e[key], err)
}

func (e Errors) Get(key string) []string {
	return e[key]
}

func (e Errors) Keys() []string {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}
	return keys
}

func (e Errors) Values() []string {
	values := make([]string, 0, len(e))
	for _, v := range e {
		values = append(values, v...)
	}
	return values
}

func (e Errors) HasErrors() bool {
	return len(e.Values()) > 0
}
