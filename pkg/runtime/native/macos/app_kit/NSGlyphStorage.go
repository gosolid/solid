//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSGlyphStorage
*/

type NSGlyphStorage struct {

}

func (r *NSGlyphStorage) SetIntAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGlyphStorage) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGlyphStorage) LayoutOptions() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGlyphStorage) InsertGlyphs() error {
  return fmt.Errorf("unimplemented")
}

