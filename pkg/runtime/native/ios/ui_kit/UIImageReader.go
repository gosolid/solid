//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIImageReader : objc.NSObject
*/

type UIImageReader struct {
  *objc.NSObject

}

func (r *UIImageReader) ReaderWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReader) ImageWithContentsOfFileURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReader) ImageWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReader) DefaultReader() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageReader) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

