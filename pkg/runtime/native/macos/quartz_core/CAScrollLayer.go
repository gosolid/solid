//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CAScrollLayer : QuartzCore.CALayer
*/

type CAScrollLayer struct {
  *CALayer

}

func (r *CAScrollLayer) ScrollToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) ScrollToRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) ScrollMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) SetScrollMode() error {
  return fmt.Errorf("unimplemented")
}

