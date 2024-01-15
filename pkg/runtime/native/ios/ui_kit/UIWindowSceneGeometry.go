//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIWindowSceneGeometry : objc.NSObject
*/

type UIWindowSceneGeometry struct {
  *objc.NSObject

}

func (r *UIWindowSceneGeometry) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) SystemFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) InterfaceOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) MinimumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) MaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometry) ResizingRestrictions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

