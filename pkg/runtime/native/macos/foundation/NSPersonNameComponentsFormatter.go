//js:package native/macos/foundation
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

func (r *NSPersonNameComponentsFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) LocalizedStringFromPersonNameComponents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) StringFromPersonNameComponents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) IsPhonetic() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) Style() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) SetPhonetic() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) PersonNameComponentsFromString() (*NSPersonNameComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponentsFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

