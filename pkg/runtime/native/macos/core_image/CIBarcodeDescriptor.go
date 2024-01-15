//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreImage.CIBarcodeDescriptor : objc.NSObject
*/

type CIBarcodeDescriptor struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

