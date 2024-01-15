//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSXMLDTD : Foundation.NSXMLNode
*/

type NSXMLDTD struct {
  *NSXMLNode

}

func (r *NSXMLDTD) ElementDeclarationForName() (*NSXMLDTDNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) SetPublicID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) SetSystemID() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) ReplaceChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) NotationDeclarationForName() (*NSXMLDTDNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) InsertChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) SetChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) PublicID() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) RemoveChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) EntityDeclarationForName() (*NSXMLDTDNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) InsertChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) AddChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) AttributeDeclarationForName() (*NSXMLDTDNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) PredefinedEntityDeclarationForName() (*NSXMLDTDNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) SystemID() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) InitWithKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDTD) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

