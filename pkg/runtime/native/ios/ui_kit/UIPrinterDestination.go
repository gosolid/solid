//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrinterDestination : objc.NSObject
*/

type UIPrinterDestination struct {
  *objc.NSObject

}

func (r *UIPrinterDestination) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) DisplayName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) SetDisplayName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) TxtRecord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) SetTxtRecord() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterDestination) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

