//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSWorkspaceAuthorization : objc.NSObject
*/

type NSWorkspaceAuthorization struct {
  *objc.NSObject

}

