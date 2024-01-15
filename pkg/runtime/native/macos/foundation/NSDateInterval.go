//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDateInterval : objc.NSObject
*/

type NSDateInterval struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSDateInterval) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IsEqualToDateInterval() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IntersectsDateInterval() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) StartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) ContainsDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) EndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) InitWithStartDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IntersectionWithDateInterval() (*NSDateInterval, error) {
  return nil, fmt.Errorf("unimplemented")
}

