//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFileAccessIntent : objc.NSObject
*/

type NSFileAccessIntent struct {
  *objc.NSObject

}

func (r *NSFileAccessIntent) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileAccessIntent) ReadingIntentWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileAccessIntent) WritingIntentWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

