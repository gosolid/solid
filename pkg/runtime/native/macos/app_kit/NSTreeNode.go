//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSTreeNode : objc.NSObject
*/

type NSTreeNode struct {
  *objc.NSObject

}

func (r *NSTreeNode) ChildNodes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) MutableChildNodes() (*foundation.NSMutableArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) TreeNodeWithRepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) DescendantNodeAtIndexPath() (*NSTreeNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) RepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) IndexPath() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) IsLeaf() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) InitWithRepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) SortWithSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeNode) ParentNode() (*NSTreeNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

