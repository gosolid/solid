//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSXMLParser : objc.NSObject
*/

type NSXMLParser struct {
  *objc.NSObject

}

func (r *NSXMLParser) SetShouldResolveExternalEntities() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) Parse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) AbortParsing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ExternalEntityResolvingPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetAllowedExternalEntityURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) AllowedExternalEntityURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) InitWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) InitWithStream() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetExternalEntityResolvingPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldReportNamespacePrefixes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetShouldReportNamespacePrefixes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldProcessNamespaces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) SetShouldProcessNamespaces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ParserError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXMLParser) ShouldResolveExternalEntities() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

