//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSWorkspace : objc.NSObject
*/

type NSWorkspace struct {
  *objc.NSObject

}

func (r *NSWorkspace) ActivateFileViewerSelectingURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) SharedWorkspace() (*NSWorkspace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) ExtendPowerOffBy() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) HideOtherApplications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLsForApplicationsToOpenURL() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) SetDefaultApplicationAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) UnmountAndEjectDeviceAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLForApplicationToOpenURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLForApplicationToOpenContentType() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) FileLabels() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) OpenURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) OpenApplicationAtURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) ShowSearchResultsForQueryString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) DuplicateURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) FrontmostApplication() (*NSRunningApplication, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) IsFilePackageAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) GetFileSystemInfoForPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) UnmountAndEjectDeviceAtPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLForApplicationWithBundleIdentifier() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLsForApplicationsWithBundleIdentifier() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) SelectFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) IconForFiles() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) RecycleURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) OpenURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) NoteFileSystemChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) IconForFile() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) SetIcon() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) FileLabelColors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) IconForContentType() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) URLsForApplicationsToOpenContentType() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) NotificationCenter() (*foundation.NSNotificationCenter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspace) MenuBarOwningApplication() (*NSRunningApplication, error) {
  return nil, fmt.Errorf("unimplemented")
}

