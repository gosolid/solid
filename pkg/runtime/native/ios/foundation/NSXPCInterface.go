//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSXPCInterface : objc.NSObject
*/

type NSXPCInterface struct {
  *objc.NSObject

}

func (r *NSXPCInterface) InterfaceWithProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) ClassesForSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetInterface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) InterfaceForSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) Protocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

