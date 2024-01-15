//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSScrubberSelectionStyle : objc.NSObject
*/

type NSScrubberSelectionStyle struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSScrubberSelectionStyle) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberSelectionStyle) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberSelectionStyle) MakeSelectionView() (*NSScrubberSelectionView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberSelectionStyle) OutlineOverlayStyle() (*NSScrubberSelectionStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberSelectionStyle) RoundedBackgroundStyle() (*NSScrubberSelectionStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

