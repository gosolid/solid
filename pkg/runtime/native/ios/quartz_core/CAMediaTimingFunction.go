//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CAMediaTimingFunction : objc.NSObject
*/

type CAMediaTimingFunction struct {
  *objc.NSObject

}

func (r *CAMediaTimingFunction) FunctionWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) FunctionWithControlPoints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) InitWithControlPoints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) GetControlPointAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

