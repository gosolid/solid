//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreImage.CIFilterGenerator : objc.NSObject
*/

type CIFilterGenerator struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
  *CIFilterConstructor
}

func (r *CIFilterGenerator) ConnectObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) DisconnectObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ExportKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) FilterGenerator() (*CIFilterGenerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) Filter() (*CIFilter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ExportedKeys() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) FilterGeneratorWithContentsOfURL() (*CIFilterGenerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) RemoveExportedKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) SetAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) RegisterFilterName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ClassAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) SetClassAttributes() error {
  return fmt.Errorf("unimplemented")
}

