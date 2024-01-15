//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct CoreGraphics.CGPathElement
*/

type CGPathElement struct {
  Type int `v8:"type"`
  Points core_foundation.CGPoint `v8:"points"`
}
