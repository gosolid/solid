//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSXMLNode : objc.NSObject
*/

type NSXMLNode struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSXMLNode) DTDNodeWithXMLString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Parent() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) LocalName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) XMLStringWithOptions() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ObjectValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ChildCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) NextNode() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) CommentWithStringValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) TextWithStringValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Detach() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) PrefixForName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) XPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) URI() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) SetURI() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Description() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) XMLString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) AttributeWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ProcessingInstructionWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) PredefinedNamespaceForPrefix() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) CanonicalXMLStringPreservingComments() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) SetObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Children() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) PreviousSibling() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) PreviousNode() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) SetStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) NodesForXPath() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Prefix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) DocumentWithRootElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) LocalNameForName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) StringValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) NextSibling() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) RootDocument() (*NSXMLDocument, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) InitWithKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ChildAtIndex() (*NSXMLNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ObjectsForXQuery() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Kind() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) NamespaceWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Document() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) ElementWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLNode) Level() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

