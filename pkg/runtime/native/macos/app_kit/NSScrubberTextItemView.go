//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSScrubberTextItemView : AppKit.NSScrubberItemView
*/

type NSScrubberTextItemView struct {
  *NSScrubberItemView

}

func (r *NSScrubberTextItemView) TextField() (*NSTextField, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberTextItemView) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberTextItemView) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

