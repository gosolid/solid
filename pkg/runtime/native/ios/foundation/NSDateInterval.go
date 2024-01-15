//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSDateInterval : objc.NSObject
*/

type NSDateInterval struct {
  *objc.NSObject

}

func (r *NSDateInterval) ContainsDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) Compare() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IsEqualToDateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IntersectsDateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) IntersectionWithDateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) EndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) Duration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) InitWithStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateInterval) StartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

