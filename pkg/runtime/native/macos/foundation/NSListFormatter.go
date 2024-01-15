//js:package native/macos/foundation
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

func (r *NSListFormatter) StringFromItems() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) StringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) ItemFormatter() (*NSFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) SetItemFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSListFormatter) LocalizedStringByJoiningStrings() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

