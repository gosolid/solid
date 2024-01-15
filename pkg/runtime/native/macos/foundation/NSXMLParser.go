//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSXMLParser : objc.NSObject
*/

type NSXMLParser struct {
  *objc.NSObject

}

func (r *NSXMLParser) InitWithStream() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ExternalEntityResolvingPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ParserError() (*NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetShouldResolveExternalEntities() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) Parse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) AbortParsing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldReportNamespacePrefixes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldResolveExternalEntities() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetShouldProcessNamespaces() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetShouldReportNamespacePrefixes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetExternalEntityResolvingPolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) AllowedExternalEntityURLs() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetAllowedExternalEntityURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldProcessNamespaces() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

