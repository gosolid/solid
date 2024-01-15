//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAnimationDelegate
*/

type NSAnimationDelegate struct {

}

func (r *NSAnimationDelegate) AnimationDidStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationDelegate) AnimationDidEnd() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationDelegate) Animation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimationDelegate) AnimationShouldStart() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

