//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSInputServer : objc.NSObject
*/

type NSInputServer struct {
  *objc.NSObject
  *NSInputServiceProvider
  *NSInputServerMouseTracker
}

func (r *NSInputServer) InitWithDelegate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

