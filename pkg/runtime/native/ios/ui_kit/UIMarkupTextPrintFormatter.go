//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMarkupTextPrintFormatter : UIKit.UIPrintFormatter
*/

type UIMarkupTextPrintFormatter struct {
  *UIPrintFormatter

}

func (r *UIMarkupTextPrintFormatter) SetMarkupText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMarkupTextPrintFormatter) InitWithMarkupText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMarkupTextPrintFormatter) MarkupText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

