//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSGlyphInfo : objc.NSObject
*/

type NSGlyphInfo struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSGlyphInfo) GlyphInfoWithCGGlyph() (*NSGlyphInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGlyphInfo) GlyphID() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGlyphInfo) BaseString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

