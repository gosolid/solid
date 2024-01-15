//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionViewSectionHeaderView
*/

type NSCollectionViewSectionHeaderView struct {

}

func (r *NSCollectionViewSectionHeaderView) SectionCollapseButton() (*NSButton, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewSectionHeaderView) SetSectionCollapseButton() error {
  return fmt.Errorf("unimplemented")
}

