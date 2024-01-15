//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface QuartzCore.CAMediaTimingFunction : objc.NSObject
*/

type CAMediaTimingFunction struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CAMediaTimingFunction) GetControlPointAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) FunctionWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) FunctionWithControlPoints() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTimingFunction) InitWithControlPoints() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

