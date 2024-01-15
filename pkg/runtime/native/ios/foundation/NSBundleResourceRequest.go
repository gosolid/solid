//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSBundleResourceRequest : objc.NSObject
*/

type NSBundleResourceRequest struct {
  *objc.NSObject

}

func (r *NSBundleResourceRequest) Progress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) BeginAccessingResourcesWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) ConditionallyBeginAccessingResourcesWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) LoadingPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Bundle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) InitWithTags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) EndAccessingResources() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) SetLoadingPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Tags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

