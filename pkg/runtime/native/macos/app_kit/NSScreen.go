//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSScreen : objc.NSObject
*/

type NSScreen struct {
  *objc.NSObject

}

func (r *NSScreen) AuxiliaryTopRightArea() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) BackingAlignedRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) DeepestScreen() (*NSScreen, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) DeviceDescription() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) SupportedWindowDepths() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScreen) BackingScaleFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScreen) CanRepresentDisplayGamut() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScreen) ConvertRectToBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) Screens() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) VisibleFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) ConvertRectFromBacking() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) SafeAreaInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSScreen) MainScreen() (*NSScreen, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) ScreensHaveSeparateSpaces() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScreen) Depth() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScreen) LocalizedName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScreen) AuxiliaryTopLeftArea() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

