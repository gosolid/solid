//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedString : objc.NSObject
*/

type NSAttributedString struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSAttributedString) AttributesAtIndex() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedString) String() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

