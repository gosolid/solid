//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPathComponentCell : AppKit.NSTextFieldCell
*/

type NSPathComponentCell struct {
  *NSTextFieldCell

}

func (r *NSPathComponentCell) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathComponentCell) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathComponentCell) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathComponentCell) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

