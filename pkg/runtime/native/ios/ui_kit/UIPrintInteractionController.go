//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIPrintInteractionController : objc.NSObject
*/

type UIPrintInteractionController struct {
  *objc.NSObject

}

func (r *UIPrintInteractionController) PrintPageRenderer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PresentFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) ShowsNumberOfCopies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) ShowsPaperOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintingItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetPrintingItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintingItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) CanPrintURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PresentFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetPrintFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SharedPrintController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetPrintInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetShowsPaperSelectionForLoadedPapers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintPaper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) IsPrintingAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) CanPrintData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PresentAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintToPrinter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) DismissAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetShowsNumberOfCopies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) ShowsPaperSelectionForLoadedPapers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) PrintableUTIs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) ShowsPageRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetShowsPageRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetPrintPageRenderer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetPrintingItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintInteractionController) SetShowsPaperOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

