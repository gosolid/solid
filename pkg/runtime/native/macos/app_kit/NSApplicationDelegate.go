//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSApplicationDelegate
*/

type NSApplicationDelegate struct {

}

func (r *NSApplicationDelegate) Application() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationShouldHandleReopen() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidResignActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationProtectedDataDidBecomeAvailable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationShouldTerminate() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillHide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationShouldAutomaticallyLocalizeKeyEquivalents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidBecomeActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillTerminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidChangeOcclusionState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationProtectedDataWillBecomeUnavailable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillFinishLaunching() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationShouldOpenUntitledFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationSupportsSecureRestorableState() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidHide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillUnhide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidChangeScreenParameters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDockMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidFinishLaunching() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillBecomeActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillResignActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationOpenUntitledFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidUnhide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationWillUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplicationDelegate) ApplicationDidUpdate() error {
  return fmt.Errorf("unimplemented")
}

