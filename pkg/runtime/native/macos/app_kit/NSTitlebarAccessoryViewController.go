//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSTitlebarAccessoryViewController : AppKit.NSViewController
*/

type NSTitlebarAccessoryViewController struct {
  *NSViewController
  *NSAnimationDelegate
  *NSAnimatablePropertyContainer
}

func (r *NSTitlebarAccessoryViewController) SetAutomaticallyAdjustsSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) ViewWillAppear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) ViewDidDisappear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) LayoutAttribute() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) SetLayoutAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) SetFullScreenMinHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) AutomaticallyAdjustsSize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) ViewDidAppear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) FullScreenMinHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTitlebarAccessoryViewController) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

