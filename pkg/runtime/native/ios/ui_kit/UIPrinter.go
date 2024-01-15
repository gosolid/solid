//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrinter : objc.NSObject
*/

type UIPrinter struct {
  *objc.NSObject

}

func (r *UIPrinter) DisplayLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) SupportedJobTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) SupportsDuplex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) PrinterWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) ContactPrinter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) MakeAndModel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) SupportsColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinter) DisplayName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

