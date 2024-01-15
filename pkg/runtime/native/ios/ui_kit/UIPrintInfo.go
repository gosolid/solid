//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrintInfo : objc.NSObject
*/

type UIPrintInfo struct {
  *objc.NSObject

}

func (r *UIPrintInfo) SetOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) PrintInfoWithDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) PrinterID() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) JobName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) DictionaryRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) PrintInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) SetPrinterID() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) SetJobName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) OutputType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) Duplex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) SetOutputType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) Orientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInfo) SetDuplex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

