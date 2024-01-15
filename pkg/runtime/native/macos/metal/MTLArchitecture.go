//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLArchitecture : objc.NSObject
*/

type MTLArchitecture struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLArchitecture) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

