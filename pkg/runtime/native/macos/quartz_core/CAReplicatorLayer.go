//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "fmt"
)

/*
interface QuartzCore.CAReplicatorLayer : QuartzCore.CALayer
*/

type CAReplicatorLayer struct {
  *CALayer

}

func (r *CAReplicatorLayer) InstanceDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceTransform() (CATransform3D, error) {
  return CATransform3D{}, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceGreenOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) PreservesDepth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceGreenOffset() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceAlphaOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetPreservesDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceRedOffset() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceAlphaOffset() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceRedOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceBlueOffset() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceBlueOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceCount() error {
  return fmt.Errorf("unimplemented")
}

