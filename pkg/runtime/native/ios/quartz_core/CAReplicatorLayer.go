//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAReplicatorLayer : QuartzCore.CALayer
*/

type CAReplicatorLayer struct {
  *CALayer

}

func (r *CAReplicatorLayer) InstanceCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) PreservesDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceDelay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceDelay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceRedOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceBlueOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceRedOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceBlueOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceAlphaOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceGreenOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetPreservesDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) SetInstanceGreenOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAReplicatorLayer) InstanceAlphaOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

