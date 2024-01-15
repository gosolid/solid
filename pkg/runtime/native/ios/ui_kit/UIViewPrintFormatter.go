//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIViewPrintFormatter : UIKit.UIPrintFormatter
*/

type UIViewPrintFormatter struct {
  *UIPrintFormatter

}

func (r *UIViewPrintFormatter) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

