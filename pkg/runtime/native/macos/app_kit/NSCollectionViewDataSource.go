//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionViewDataSource
*/

type NSCollectionViewDataSource struct {

}

func (r *NSCollectionViewDataSource) CollectionView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDataSource) NumberOfSectionsInCollectionView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

