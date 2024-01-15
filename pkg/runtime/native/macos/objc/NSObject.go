//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol objc.NSObject
*/

type NSObject struct {

}

func (r *NSObject) IsKindOfClass() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) Autorelease() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) DebugDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) Self() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) PerformSelector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) IsProxy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) IsMemberOfClass() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) RespondsToSelector() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) Zone() (*foundation.NSZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) Description() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) IsEqual() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) Retain() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) Hash() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSObject) Class() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObject) ConformsToProtocol() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObject) Release() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObject) RetainCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSObject) Superclass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

