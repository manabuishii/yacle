package cwl

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

// ProvidedInputs ...
type ProvidedInputs map[string]ProvidedInput

// NewInputs ...
func NewInputs() ProvidedInputs {
	return ProvidedInputs{}
}

// ProvidedInput ...
type ProvidedInput struct {
	Self  interface{}
	Type  reflect.Kind
	Class string `yaml:"class"`
	Path  string `yaml:"path"`
}

// Arg ...
func (provided ProvidedInput) Arg() string {
	switch provided.Type {
	case reflect.String:
		return provided.Self.(string)
	case reflect.Map:
		return provided.Path
	}
	return ""
}

// Decode ...
func (inputs ProvidedInputs) Decode(r io.Reader) error {
	dict := map[string]interface{}{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(b, &dict); err != nil {
		return err
	}
	for key, val := range dict {
		switch reflect.ValueOf(val).Kind() {
		case reflect.String:
			inputs[key] = ProvidedInput{
				Self: val,
				Type: reflect.String,
			}
		case reflect.Map:
			inputs[key] = ProvidedInput{
				Self:  val,
				Type:  reflect.Map,
				Class: reflect.ValueOf(val).MapIndex(reflect.ValueOf("class")).Interface().(string),
				Path:  reflect.ValueOf(val).MapIndex(reflect.ValueOf("path")).Interface().(string),
			}
		default:
			fmt.Println(key, reflect.TypeOf(val).Kind())
		}
	}
	return nil
}
