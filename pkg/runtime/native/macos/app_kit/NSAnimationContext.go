//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSAnimationContext : objc.NSObject
*/

type NSAnimationContext struct {
  *objc.NSObject

}

func (r *NSAnimationContext) EndGrouping() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) SetTimingFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) SetCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) AllowsImplicitAnimation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) SetAllowsImplicitAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) RunAnimationGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) BeginGrouping() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) CurrentContext() (*NSAnimationContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) SetDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) TimingFunction() (*CAMediaTimingFunction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimationContext) CompletionHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

