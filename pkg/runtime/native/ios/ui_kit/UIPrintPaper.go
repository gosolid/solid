//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIPrintPaper : objc.NSObject
*/

type UIPrintPaper struct {
  *objc.NSObject

}

func (r *UIPrintPaper) BestPaperForPageSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPaper) PaperSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPaper) PrintableRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

