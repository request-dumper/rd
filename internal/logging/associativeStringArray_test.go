package logging

import (
	"bytes"
	"github.com/rs/zerolog"
	"testing"
)

func TestAssociativeStringArray_MarshalZerologObject(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	assoStrgArray := AssociativeStringArray{
		value: getTestData(),
	}

	buf := &bytes.Buffer{}
	log := zerolog.New(zerolog.ConsoleWriter{Out: buf, NoColor: true})

	log.Info().
		Object("test", &assoStrgArray).
		Msg("msg")

	want := "<nil> INF msg test={\"Firstkey\":[\"firstKey-firstValue\",\"firstKey-secondValue\"],\"Secondkey\":[\"secondKey-firstValue\"]}\n"

	got := buf.String()
	if want != got {
		t.Errorf("Unexpected output %q, want: %q", got, want)
	}
}
