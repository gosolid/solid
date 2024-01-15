//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextTab : objc.NSObject
*/

type NSTextTab struct {
  *objc.NSObject

}

func (r *NSTextTab) ColumnTerminatorsForLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) InitWithTextAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Alignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

