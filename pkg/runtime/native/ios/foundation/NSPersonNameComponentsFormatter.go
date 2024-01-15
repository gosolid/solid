//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSPersonNameComponentsFormatter : Foundation.NSFormatter
*/

type NSPersonNameComponentsFormatter struct {
  *NSFormatter

}

func (r *NSPersonNameComponentsFormatter) SetPhonetic() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) LocalizedStringFromPersonNameComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) PersonNameComponentsFromString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) SetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) IsPhonetic() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) StringFromPersonNameComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

