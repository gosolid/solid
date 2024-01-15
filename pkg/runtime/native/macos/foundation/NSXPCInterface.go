//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSXPCInterface : objc.NSObject
*/

type NSXPCInterface struct {
  *objc.NSObject

}

func (r *NSXPCInterface) SetClasses() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) ClassesForSelector() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetInterface() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) InterfaceForSelector() (*NSXPCInterface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) XPCTypeForSelector() (*objc.XpcTypeS, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) Protocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) InterfaceWithProtocol() (*NSXPCInterface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetXPCType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCInterface) SetProtocol() error {
  return fmt.Errorf("unimplemented")
}

