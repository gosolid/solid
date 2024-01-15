//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIPrinterPickerController : objc.NSObject
*/

type UIPrinterPickerController struct {
  *objc.NSObject

}

func (r *UIPrinterPickerController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) PrinterPickerControllerWithInitiallySelectedPrinter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) PresentAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) PresentFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) PresentFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) DismissAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrinterPickerController) SelectedPrinter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

