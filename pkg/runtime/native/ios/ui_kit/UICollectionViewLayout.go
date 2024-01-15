//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UICollectionViewLayout : objc.NSObject
*/

type UICollectionViewLayout struct {
  *objc.NSObject

}

func (r *UICollectionViewLayout) RegisterClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) RegisterNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) CollectionView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) InvalidateLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayout) InvalidateLayoutWithContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

