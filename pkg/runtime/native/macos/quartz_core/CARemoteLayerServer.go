//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CARemoteLayerServer : objc.NSObject
*/

type CARemoteLayerServer struct {
  *objc.NSObject

}

func (r *CARemoteLayerServer) SharedServer() (*CARemoteLayerServer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARemoteLayerServer) ServerPort() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

