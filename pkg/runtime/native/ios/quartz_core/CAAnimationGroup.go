//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAAnimationGroup : QuartzCore.CAAnimation
*/

type CAAnimationGroup struct {
  *CAAnimation

}

func (r *CAAnimationGroup) Animations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimationGroup) SetAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

