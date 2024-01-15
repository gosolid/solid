//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSLocale : objc.NSObject
*/

type NSLocale struct {
  *objc.NSObject

}

func (r *NSLocale) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) DisplayNameForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) InitWithLocaleIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLocale) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

