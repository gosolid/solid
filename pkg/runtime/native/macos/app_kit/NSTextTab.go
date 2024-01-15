//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextTab : objc.NSObject
*/

type NSTextTab struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
  *foundation.NSSecureCoding
}

func (r *NSTextTab) InitWithTextAlignment() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Location() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) Options() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTab) ColumnTerminatorsForLocale() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

