//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIFilterGenerator : objc.NSObject
*/

type CIFilterGenerator struct {
  *objc.NSObject

}

func (r *CIFilterGenerator) FilterGenerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) FilterGeneratorWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) InitWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ExportedKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ExportKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) RemoveExportedKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ClassAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) DisconnectObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) RegisterFilterName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) ConnectObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) Filter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterGenerator) SetClassAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

