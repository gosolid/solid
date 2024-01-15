//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrintServiceExtension : objc.NSObject
*/

type UIPrintServiceExtension struct {
  *objc.NSObject

}

func (r *UIPrintServiceExtension) PrinterDestinationsForPrintInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

