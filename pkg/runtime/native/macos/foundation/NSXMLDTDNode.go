//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSXMLDTDNode : Foundation.NSXMLNode
*/

type NSXMLDTDNode struct {
  *NSXMLNode

}

func (r *NSXMLDTDNode) PublicID() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) SetSystemID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) NotationName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) DTDKind() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) SetDTDKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) SetPublicID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) SystemID() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) SetNotationName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) InitWithXMLString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) InitWithKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTDNode) IsExternal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

