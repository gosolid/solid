//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIEventAttribution : objc.NSObject
*/

type UIEventAttribution struct {
  *objc.NSObject

}

func (r *UIEventAttribution) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) SourceIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) DestinationURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) ReportEndpoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) SourceDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) Purchaser() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEventAttribution) InitWithSourceIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

