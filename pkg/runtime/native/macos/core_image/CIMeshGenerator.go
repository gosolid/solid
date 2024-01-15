//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIMeshGenerator
*/

type CIMeshGenerator struct {

}

func (r *CIMeshGenerator) Mesh() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMeshGenerator) SetMesh() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMeshGenerator) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMeshGenerator) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMeshGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMeshGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

