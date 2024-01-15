//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface CoreImage.CIBarcodeDescriptor : objc.NSObject
*/

type CIBarcodeDescriptor struct {
  *objc.NSObject

}

