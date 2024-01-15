//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSListFormatter : Foundation.NSFormatter
*/

type NSListFormatter struct {
  *NSFormatter

}

func (r *NSListFormatter) SetItemFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) LocalizedStringByJoiningStrings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) StringFromItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) StringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) ItemFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

