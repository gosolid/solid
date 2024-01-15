//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSViewControllerPresentationAnimator
*/

type NSViewControllerPresentationAnimator struct {

}

func (r *NSViewControllerPresentationAnimator) AnimatePresentationOfViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewControllerPresentationAnimator) AnimateDismissalOfViewController() error {
  return fmt.Errorf("unimplemented")
}

