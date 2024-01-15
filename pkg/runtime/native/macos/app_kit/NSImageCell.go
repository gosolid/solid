//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSImageCell : AppKit.NSCell
*/

type NSImageCell struct {
  *NSCell
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSImageCell) ImageAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageCell) SetImageAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageCell) ImageScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageCell) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageCell) ImageFrameStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageCell) SetImageFrameStyle() error {
  return fmt.Errorf("unimplemented")
}

