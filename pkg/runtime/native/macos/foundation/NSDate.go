//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDate : objc.NSObject
*/

type NSDate struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSDate) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDate) InitWithTimeIntervalSinceReferenceDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDate) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDate) TimeIntervalSinceReferenceDate() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

