//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSDiffableDataSourceSnapshot : objc.NSObject
*/

type NSDiffableDataSourceSnapshot struct {
  *objc.NSObject

}

func (r *NSDiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) IndexOfSectionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) AppendItemsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) DeleteItemsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) DeleteAllItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) AppendSectionsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) SectionIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ItemIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) InsertItemsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadItemsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReconfigureItemsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) InsertSectionsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) MoveSectionWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfItemsInSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) MoveItemWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadedSectionIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadedItemIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReconfiguredItemIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) IndexOfItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

