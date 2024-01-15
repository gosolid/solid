//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSProcessInfo : objc.NSObject
*/

type NSProcessInfo struct {
  *objc.NSObject

}

func (r *NSProcessInfo) EnableAutomaticTermination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) Arguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) HostName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ActiveProcessorCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) IsOperatingSystemAtLeastVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) EnableSuddenTermination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) DisableAutomaticTermination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) SystemUptime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) DisableSuddenTermination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) SetProcessName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) GloballyUniqueString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessorCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) PhysicalMemory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) AutomaticTerminationSupportEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) ProcessIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) OperatingSystemVersionString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) SetAutomaticTerminationSupportEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProcessInfo) Environment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

