//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSPersistentHistoryToken : objc.NSObject
*/

type NSPersistentHistoryToken struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

