//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSWindowDelegate
*/

type NSWindowDelegate struct {

}

func (r *NSWindowDelegate) WindowWillExitVersionBrowser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) CustomWindowsToEnterFullScreenForWindow() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidExpose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidChangeScreenProfile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidEnterFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidExitFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillUseStandardFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) Window() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidMiniaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidDeminiaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillBeginSheet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidEndSheet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidChangeBackingProperties() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidEndLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidChangeOcclusionState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillResize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidFailToExitFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidMove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidChangeScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillMiniaturize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillExitFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillEnterVersionBrowser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowShouldZoom() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidBecomeKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidResignKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidResignMain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidEnterVersionBrowser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowShouldClose() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) PreviewRepresentableActivityItemsForWindow() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidBecomeMain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillReturnFieldEditor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidFailToEnterFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillMove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowDidExitVersionBrowser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillReturnUndoManager() (*foundation.NSUndoManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) CustomWindowsToExitFullScreenForWindow() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillStartLiveResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowDelegate) WindowWillEnterFullScreen() error {
  return fmt.Errorf("unimplemented")
}

