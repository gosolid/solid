//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CAAnimationGroup : QuartzCore.CAAnimation
*/

type CAAnimationGroup struct {
  *CAAnimation

}

func (r *CAAnimationGroup) Animations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimationGroup) SetAnimations() error {
  return fmt.Errorf("unimplemented")
}

