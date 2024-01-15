//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSFontChanging
*/

type NSFontChanging struct {

}

func (r *NSFontChanging) ChangeFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontChanging) ValidModesForFontPanel() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

