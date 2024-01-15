//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAnimatablePropertyContainer
*/

type NSAnimatablePropertyContainer struct {

}

func (r *NSAnimatablePropertyContainer) AnimationForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimatablePropertyContainer) DefaultAnimationForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimatablePropertyContainer) Animations() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimatablePropertyContainer) SetAnimations() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimatablePropertyContainer) Animator() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

