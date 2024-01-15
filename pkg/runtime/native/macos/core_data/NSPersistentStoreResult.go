//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreData.NSPersistentStoreResult : objc.NSObject
*/

type NSPersistentStoreResult struct {
  *objc.NSObject

}

