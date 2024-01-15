//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSRunningApplication : objc.NSObject
*/

type NSRunningApplication struct {
  *objc.NSObject

}

func (r *NSRunningApplication) BundleURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) Unhide() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) RunningApplicationsWithBundleIdentifier() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) TerminateAutomaticallyTerminableApplications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) OwnsMenuBar() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ActivationPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) BundleIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) Icon() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) Terminate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ExecutableURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ProcessIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) LaunchDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ExecutableArchitecture() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) CurrentApplication() (*NSRunningApplication, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) Hide() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) RunningApplicationWithProcessIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) IsTerminated() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) LocalizedName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ActivateFromApplication() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ActivateWithOptions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) ForceTerminate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) IsFinishedLaunching() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRunningApplication) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

