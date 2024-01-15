//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
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

func (r *CAMetalDisplayLinkUpdate) TargetTimestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLinkUpdate) TargetPresentationTimestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

