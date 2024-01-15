//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol QuartzCore.CALayoutManager
*/

type CALayoutManager struct {

}

func (r *CALayoutManager) LayoutSublayersOfLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CALayoutManager) PreferredSizeOfLayer() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CALayoutManager) InvalidateLayoutOfLayer() error {
  return fmt.Errorf("unimplemented")
}

