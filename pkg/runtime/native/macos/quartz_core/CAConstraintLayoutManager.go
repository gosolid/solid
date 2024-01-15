//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CAConstraintLayoutManager : objc.NSObject
*/

type CAConstraintLayoutManager struct {
  *objc.NSObject
  *CALayoutManager
}

func (r *CAConstraintLayoutManager) LayoutManager() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

