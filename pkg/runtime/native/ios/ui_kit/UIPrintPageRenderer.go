//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrintPageRenderer : objc.NSObject
*/

type UIPrintPageRenderer struct {
  *objc.NSObject

}

func (r *UIPrintPageRenderer) DrawHeaderForPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) HeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) PaperRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) AddPrintFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) CurrentRenderingQualityForRequestedRenderingQuality() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) DrawPrintFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) NumberOfPages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) PrepareForDrawingPages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) DrawFooterForPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) SetHeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) SetPrintFormatters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) PrintFormattersForPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) FooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) PrintableRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) PrintFormatters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) DrawPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) DrawContentForPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintPageRenderer) SetFooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

