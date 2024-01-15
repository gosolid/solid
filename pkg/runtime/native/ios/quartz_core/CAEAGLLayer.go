//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAEAGLLayer : QuartzCore.CALayer
*/

type CAEAGLLayer struct {
  *CALayer

}

func (r *CAEAGLLayer) PresentsWithTransaction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEAGLLayer) SetPresentsWithTransaction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

