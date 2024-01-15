//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
interface AppKit.NSPDFImageRep : AppKit.NSImageRep
*/

type NSPDFImageRep struct {
  *NSImageRep

}

func (r *NSPDFImageRep) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) PDFRepresentation() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) CurrentPage() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) SetCurrentPage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) PageCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPDFImageRep) ImageRepWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

