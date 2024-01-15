//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSPrintOperation : objc.NSObject
*/

type NSPrintOperation struct {
  *objc.NSObject

}

func (r *NSPrintOperation) PrintOperationWithView() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) RunOperationModalForWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) RunOperation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) CleanUpOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) CreateContext() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetPDFPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) ShowsPrintPanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PrintPanel() (*NSPrintPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetPrintPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) Context() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) EPSOperationWithView() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) CurrentOperation() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetJobTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PageRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) CurrentPage() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetCanSpawnSeparateThread() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetCurrentOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PreferredRenderingQuality() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetShowsPrintPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) ShowsProgressPanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PrintInfo() (*NSPrintInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetShowsProgressPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) CanSpawnSeparateThread() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) SetPageOrder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PDFOperationWithView() (*NSPrintOperation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) DestroyContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) DeliverResult() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) IsCopyingOperation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) JobTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PDFPanel() (*NSPDFPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintOperation) PageOrder() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

