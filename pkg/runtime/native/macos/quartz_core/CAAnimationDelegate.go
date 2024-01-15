//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol QuartzCore.CAAnimationDelegate
*/

type CAAnimationDelegate struct {

}

func (r *CAAnimationDelegate) AnimationDidStart() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAAnimationDelegate) AnimationDidStop() error {
  return fmt.Errorf("unimplemented")
}

