//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextList : objc.NSObject
*/

type NSTextList struct {
  *objc.NSObject

}

func (r *NSTextList) ListOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) StartingItemNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) SetStartingItemNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) IsOrdered() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) InitWithMarkerFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) MarkerForItemNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) MarkerFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

