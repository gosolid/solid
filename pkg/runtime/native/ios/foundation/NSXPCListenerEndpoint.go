//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSXPCListenerEndpoint : objc.NSObject
*/

type NSXPCListenerEndpoint struct {
  *objc.NSObject

}

