//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLDrawable
*/

type MTLDrawable struct {

}

func (r *MTLDrawable) PresentedTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDrawable) DrawableID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDrawable) Present() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDrawable) PresentAtTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDrawable) PresentAfterMinimumDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDrawable) AddPresentedHandler() error {
  return fmt.Errorf("unimplemented")
}

