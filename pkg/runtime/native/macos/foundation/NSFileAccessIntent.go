//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileAccessIntent : objc.NSObject
*/

type NSFileAccessIntent struct {
  *objc.NSObject

}

func (r *NSFileAccessIntent) WritingIntentWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileAccessIntent) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileAccessIntent) ReadingIntentWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

