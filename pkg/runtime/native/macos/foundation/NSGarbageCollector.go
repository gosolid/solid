//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSGarbageCollector : objc.NSObject
*/

type NSGarbageCollector struct {
  *objc.NSObject

}

func (r *NSGarbageCollector) DefaultCollector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) CollectExhaustively() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) DisableCollectorForPointer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) Zone() (*foundation.NSZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) IsCollecting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) Disable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) Enable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) CollectIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGarbageCollector) EnableCollectorForPointer() error {
  return fmt.Errorf("unimplemented")
}

