//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionViewElement
*/

type NSCollectionViewElement struct {

}

func (r *NSCollectionViewElement) PrepareForReuse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewElement) ApplyLayoutAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewElement) WillTransitionFromLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewElement) DidTransitionFromLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewElement) PreferredLayoutAttributesFittingAttributes() (*NSCollectionViewLayoutAttributes, error) {
  return nil, fmt.Errorf("unimplemented")
}

