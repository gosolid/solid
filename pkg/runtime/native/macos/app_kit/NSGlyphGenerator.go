//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSGlyphGenerator : objc.NSObject
*/

type NSGlyphGenerator struct {
  *objc.NSObject

}

func (r *NSGlyphGenerator) GenerateGlyphsForGlyphStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGlyphGenerator) SharedGlyphGenerator() (*NSGlyphGenerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

