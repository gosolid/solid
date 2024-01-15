//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface QuartzCore.CATiledLayer : QuartzCore.CALayer
*/

type CATiledLayer struct {
  *CALayer

}

func (r *CATiledLayer) TileSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) SetTileSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) FadeDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) LevelsOfDetail() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) SetLevelsOfDetail() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) LevelsOfDetailBias() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) SetLevelsOfDetailBias() error {
  return fmt.Errorf("unimplemented")
}

