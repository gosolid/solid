//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSCollectionViewCompositionalLayout : AppKit.NSCollectionViewLayout
*/

type NSCollectionViewCompositionalLayout struct {
  *NSCollectionViewLayout

}

func (r *NSCollectionViewCompositionalLayout) InitWithSection() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayout) InitWithSectionProvider() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayout) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayout) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayout) Configuration() (*NSCollectionViewCompositionalLayoutConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayout) SetConfiguration() error {
  return fmt.Errorf("unimplemented")
}

