//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSQueryGenerationToken : objc.NSObject
*/

type NSQueryGenerationToken struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSQueryGenerationToken) CurrentQueryGenerationToken() (*NSQueryGenerationToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

