//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISimpleTextPrintFormatter : UIKit.UIPrintFormatter
*/

type UISimpleTextPrintFormatter struct {
  *UIPrintFormatter

}

func (r *UISimpleTextPrintFormatter) SetFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) Color() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) SetColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) SetTextAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) InitWithText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) InitWithAttributedText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) Text() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) SetAttributedText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) SetText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) AttributedText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) Font() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISimpleTextPrintFormatter) TextAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

