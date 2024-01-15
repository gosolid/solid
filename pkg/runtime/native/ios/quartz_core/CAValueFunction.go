//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CAValueFunction : objc.NSObject
*/

type CAValueFunction struct {
  *objc.NSObject

}

func (r *CAValueFunction) FunctionWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAValueFunction) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

