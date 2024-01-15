//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewTransitionLayout : UIKit.UICollectionViewLayout
*/

type UICollectionViewTransitionLayout struct {
  *UICollectionViewLayout

}

func (r *UICollectionViewTransitionLayout) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) UpdateValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) TransitionProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) SetTransitionProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) NextLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) InitWithCurrentLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) ValueForAnimatedKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewTransitionLayout) CurrentLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

