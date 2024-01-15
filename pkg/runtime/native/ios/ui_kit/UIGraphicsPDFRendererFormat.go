//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsPDFRendererFormat : UIKit.UIGraphicsRendererFormat
*/

type UIGraphicsPDFRendererFormat struct {
  *UIGraphicsRendererFormat

}

func (r *UIGraphicsPDFRendererFormat) DocumentInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererFormat) SetDocumentInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

