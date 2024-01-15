//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CARemoteLayerClient : objc.NSObject
*/

type CARemoteLayerClient struct {
  *objc.NSObject

}

func (r *CARemoteLayerClient) InitWithServerPort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARemoteLayerClient) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARemoteLayerClient) ClientId() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CARemoteLayerClient) Layer() (*CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARemoteLayerClient) SetLayer() error {
  return fmt.Errorf("unimplemented")
}

