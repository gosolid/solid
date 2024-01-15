//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKNotificationID : objc.NSObject
*/

type CKNotificationID struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

