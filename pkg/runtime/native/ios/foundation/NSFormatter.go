//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFormatter : objc.NSObject
*/

type NSFormatter struct {
  *objc.NSObject

}

func (r *NSFormatter) StringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) AttributedStringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) EditingStringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormatter) IsPartialStringValid() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

