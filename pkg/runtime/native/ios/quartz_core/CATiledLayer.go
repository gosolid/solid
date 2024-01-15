//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CATiledLayer : QuartzCore.CALayer
*/

type CATiledLayer struct {
  *CALayer

}

func (r *CATiledLayer) SetLevelsOfDetailBias() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) TileSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) SetTileSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) FadeDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) LevelsOfDetail() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) SetLevelsOfDetail() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATiledLayer) LevelsOfDetailBias() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

