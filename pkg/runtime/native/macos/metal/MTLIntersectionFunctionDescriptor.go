//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLIntersectionFunctionDescriptor : Metal.MTLFunctionDescriptor
*/

type MTLIntersectionFunctionDescriptor struct {
  *MTLFunctionDescriptor
  *foundation.NSCopying
}

