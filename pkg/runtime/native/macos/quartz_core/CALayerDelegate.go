//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol QuartzCore.CALayerDelegate
*/

type CALayerDelegate struct {

}

func (r *CALayerDelegate) LayoutSublayersOfLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayerDelegate) ActionForLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CALayerDelegate) DisplayLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayerDelegate) DrawLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayerDelegate) LayerWillDraw() error {
  return fmt.Errorf("unimplemented")
}

