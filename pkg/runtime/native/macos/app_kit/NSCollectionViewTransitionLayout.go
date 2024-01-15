//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSCollectionViewTransitionLayout : AppKit.NSCollectionViewLayout
*/

type NSCollectionViewTransitionLayout struct {
  *NSCollectionViewLayout

}

func (r *NSCollectionViewTransitionLayout) TransitionProgress() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) SetTransitionProgress() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) CurrentLayout() (*NSCollectionViewLayout, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) NextLayout() (*NSCollectionViewLayout, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) InitWithCurrentLayout() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) UpdateValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewTransitionLayout) ValueForAnimatedKey() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

