package logging

import "github.com/rs/zerolog"

type AssociativeStringArray struct {
	value map[string][]string
}

func (h *AssociativeStringArray) MarshalZerologObject(e *zerolog.Event) {
	for key := range h.value {
		stringArray := zerolog.Arr()
		for i := range h.value[key] {
			stringArray.Str(h.value[key][i])
		}
		e.Array(key, stringArray)
	}
}
