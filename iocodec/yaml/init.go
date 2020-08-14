package yaml

import (
	"io"

	"github.com/fatih/structs"
	"gopkg.in/yaml.v2"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

func init() {
	client.RegisterInputDecoder("yaml", decoderMaker)
	client.RegisterOutputEncoder("yaml", encoderMaker)
}

func decoderMaker(r io.Reader) iocodec.Decoder {
	return yaml.NewDecoder(r).Decode
}

func encoderMaker(w io.Writer) iocodec.Encoder {
	return func(v interface{}) error {
		// workaround: yaml encoder doesn't honor json tags so pre-process with structs first
		s := structs.New(v)
		s.TagName = "json"
		v = s.Map()

		e := yaml.NewEncoder(w)
		defer e.Close()
		return e.Encode(v)
	}
}
