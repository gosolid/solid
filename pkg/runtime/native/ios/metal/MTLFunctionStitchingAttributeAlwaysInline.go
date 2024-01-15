//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLFunctionStitchingAttributeAlwaysInline : objc.NSObject
*/

type MTLFunctionStitchingAttributeAlwaysInline struct {
  *objc.NSObject

}

