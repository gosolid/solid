//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UISceneDestructionRequestOptions : objc.NSObject
*/

type UISceneDestructionRequestOptions struct {
  *objc.NSObject

}

