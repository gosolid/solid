//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedString : objc.NSObject
*/

type NSAttributedString struct {
  *objc.NSObject

}

func (r *NSAttributedString) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedString) AttributesAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

