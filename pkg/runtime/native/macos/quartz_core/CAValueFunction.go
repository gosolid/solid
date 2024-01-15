//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface QuartzCore.CAValueFunction : objc.NSObject
*/

type CAValueFunction struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CAValueFunction) FunctionWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAValueFunction) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

