//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSBundleResourceRequest : objc.NSObject
*/

type NSBundleResourceRequest struct {
  *objc.NSObject
  *NSProgressReporting
}

func (r *NSBundleResourceRequest) Progress() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) InitWithTags() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) ConditionallyBeginAccessingResourcesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) EndAccessingResources() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) LoadingPriority() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Bundle() (*NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) BeginAccessingResourcesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) SetLoadingPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBundleResourceRequest) Tags() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

