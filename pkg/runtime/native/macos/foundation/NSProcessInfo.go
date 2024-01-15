//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSProcessInfo : objc.NSObject
*/

type NSProcessInfo struct {
  *objc.NSObject

}

func (r *NSProcessInfo) SystemUptime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) DisableSuddenTermination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessInfo() (*NSProcessInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) Arguments() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemVersion() (NSOperatingSystemVersion, error) {
  return NSOperatingSystemVersion{}, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessorCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ActiveProcessorCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) IsOperatingSystemAtLeastVersion() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) EnableSuddenTermination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) DisableAutomaticTermination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) Environment() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) SetProcessName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) PhysicalMemory() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) SetAutomaticTerminationSupportEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) GloballyUniqueString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemVersionString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) AutomaticTerminationSupportEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystem() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) EnableAutomaticTermination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) HostName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

