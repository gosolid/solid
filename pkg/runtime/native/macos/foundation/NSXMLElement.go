//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSXMLElement : Foundation.NSXMLNode
*/

type NSXMLElement struct {
  *NSXMLNode

}

func (r *NSXMLElement) InitWithKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) NamespaceForPrefix() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) InsertChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) ElementsForLocalName() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) ResolveNamespaceForName() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) ElementsForName() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) InsertChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) SetAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) SetAttributesWithDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) AttributeForName() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) AttributeForLocalName() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) ResolvePrefixForNamespaceURI() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) SetNamespaces() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) InitWithXMLString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) AddAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) RemoveNamespaceForPrefix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) RemoveAttributeForName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) AddChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) Namespaces() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) Attributes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) AddNamespace() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) RemoveChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) SetChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) ReplaceChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLElement) NormalizeAdjacentTextNodesPreservingCDATA() error {
  return fmt.Errorf("unimplemented")
}

