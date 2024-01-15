//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSLocale : objc.NSObject
*/

type NSLocale struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSLocale) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) DisplayNameForKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) InitWithLocaleIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

