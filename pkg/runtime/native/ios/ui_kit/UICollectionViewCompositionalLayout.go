//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewCompositionalLayout : UIKit.UICollectionViewLayout
*/

type UICollectionViewCompositionalLayout struct {
  *UICollectionViewLayout

}

func (r *UICollectionViewCompositionalLayout) SetConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayout) InitWithSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayout) InitWithSectionProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayout) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayout) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayout) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

