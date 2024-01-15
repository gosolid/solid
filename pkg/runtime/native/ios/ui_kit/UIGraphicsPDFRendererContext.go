//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsPDFRendererContext : UIKit.UIGraphicsRendererContext
*/

type UIGraphicsPDFRendererContext struct {
  *UIGraphicsRendererContext

}

func (r *UIGraphicsPDFRendererContext) BeginPageWithBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererContext) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererContext) AddDestinationWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererContext) SetDestinationWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererContext) PdfContextBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsPDFRendererContext) BeginPage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

