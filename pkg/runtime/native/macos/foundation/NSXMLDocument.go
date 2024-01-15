//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSXMLDocument : Foundation.NSXMLNode
*/

type NSXMLDocument struct {
  *NSXMLNode

}

func (r *NSXMLDocument) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InsertChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ValidateAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) IsStandalone() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) MIMEType() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ReplacementClassForClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetRootElement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InsertChildren() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetCharacterEncoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetStandalone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) DocumentContentKind() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InitWithXMLString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) RemoveChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) RootElement() (*NSXMLElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ReplaceChildAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) DTD() (*NSXMLDTD, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetDTD() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) AddChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) XMLDataWithOptions() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ObjectByApplyingXSLTString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ObjectByApplyingXSLTAtURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) CharacterEncoding() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) Version() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetMIMEType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) InitWithRootElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) ObjectByApplyingXSLT() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) SetDocumentContentKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLDocument) XMLData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

