//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFormatter : objc.NSObject
*/

type NSFormatter struct {
  *objc.NSObject
  *NSCopying
  *NSCoding
}

func (r *NSFormatter) StringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) AttributedStringForObjectValue() (*NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) EditingStringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) IsPartialStringValid() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

