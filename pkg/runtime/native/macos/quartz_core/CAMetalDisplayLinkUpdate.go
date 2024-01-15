//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CAMetalDisplayLinkUpdate : objc.NSObject
*/

type CAMetalDisplayLinkUpdate struct {
  *objc.NSObject

}

func (r *CAMetalDisplayLinkUpdate) Drawable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLinkUpdate) TargetTimestamp() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLinkUpdate) TargetPresentationTimestamp() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

