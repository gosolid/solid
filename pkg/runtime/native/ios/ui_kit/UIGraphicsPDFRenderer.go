//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsPDFRenderer : UIKit.UIGraphicsRenderer
*/

type UIGraphicsPDFRenderer struct {
  *UIGraphicsRenderer

}

func (r *UIGraphicsPDFRenderer) InitWithBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRenderer) WritePDFToURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRenderer) PDFDataWithActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

