//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAScrollLayer : QuartzCore.CALayer
*/

type CAScrollLayer struct {
  *CALayer

}

func (r *CAScrollLayer) ScrollMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) SetScrollMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) ScrollToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAScrollLayer) ScrollToRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

