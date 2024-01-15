//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSPrintInfo : objc.NSObject
*/

type NSPrintInfo struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSPrintInfo) Printer() (*NSPrinter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetLeftMargin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetTopMargin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetHorizontallyCentered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetVerticalPagination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetRightMargin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) ImageablePageBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) DefaultPrinter() (*NSPrinter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PaperName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetPaperName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) ScalingFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) LeftMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) VerticalPagination() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) JobDisposition() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PMPageFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SharedPrintInfo() (*NSPrintInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PaperSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetVerticallyCentered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) RightMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetBottomMargin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetJobDisposition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetUpPrintOperationDefaultValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) UpdateFromPMPrintSettings() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) LocalizedPaperName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetSelectionOnly() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) UpdateFromPMPageFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) TakeSettingsFromPDFInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) Orientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) HorizontalPagination() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) IsHorizontallyCentered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) IsVerticallyCentered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) IsSelectionOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) InitWithDictionary() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetSharedPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetPaperSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) TopMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) BottomMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetHorizontalPagination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetPrinter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PrintSettings() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) Dictionary() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PMPrintSession() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) PMPrintSettings() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintInfo) SetScalingFactor() error {
  return fmt.Errorf("unimplemented")
}

