//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSApplication : AppKit.NSResponder
*/

type NSApplication struct {
  *NSResponder
  *NSUserInterfaceValidations
  *NSMenuItemValidation
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSApplication) SetActivationPolicy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplication) OcclusionState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) Run() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) ReportException() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetMainMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) UnhideWithoutActivation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) YieldActivationToApplication() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) BeginModalSessionForWindow() (*NSModalSession, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) HideOtherApplications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) DetachDrawingThread() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplication) Unhide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) Deactivate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) AbortModal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) Windows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) ApplicationIconImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) IsProtectedDataAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplication) WindowWithWindowNumber() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) FinishLaunching() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetWindowsNeedUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) UpdateWindows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) HelpMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) PresentationOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetPresentationOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) CurrentSystemPresentationOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) DockTile() (*NSDockTile, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) Hide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) ActivateIgnoringOtherApps() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) EnumerateWindowsWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) ActivationPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) MainMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) CancelUserAttentionRequest() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) MainWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) IsRunning() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSApplication) ModalWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) RunModalSession() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) RequestUserAttention() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) ReplyToApplicationShouldTerminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) ReplyToOpenOrPrint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetHelpMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) SetApplicationIconImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) OrderFrontCharacterPalette() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) Activate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) YieldActivationToApplicationWithBundleIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) UnhideAllApplications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) StopModal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) StopModalWithCode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) EndModalSession() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) Terminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) KeyWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) Stop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) PreventWindowOrdering() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSApplication) RunModalForWindow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSApplication) SharedApplication() (*NSApplication, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSApplication) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

