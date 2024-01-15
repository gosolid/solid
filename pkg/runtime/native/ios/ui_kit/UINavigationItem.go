//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UINavigationItem : objc.NSObject
*/

type UINavigationItem struct {
  *objc.NSObject

}

func (r *UINavigationItem) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetLeftBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetDocumentProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) LeftItemsSupplementBackButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) LargeTitleDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetSearchController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetHidesSearchBarWhenScrolling() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetPreferredSearchBarPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) HidesBackButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetRenameDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) DocumentProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) LeadingItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetTitleView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) BackBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetBackBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) TitleMenuProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetTrailingItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) PinnedTrailingGroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetHidesBackButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetTitleMenuProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) LeftBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetBackButtonTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) RightBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) AdditionalOverflowItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetRightBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) Prompt() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetBackAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) LeftBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) RightBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) CustomizationIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetCompactScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetRightBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) BackAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetCustomizationIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetLargeTitleDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SearchBarPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetLeftBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetCenterItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) OverflowPresentationSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) StandardAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) BackButtonDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) PreferredSearchBarPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetStandardAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetPrompt() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) RenameDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) CenterItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) HidesSearchBarWhenScrolling() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) CompactScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) TitleView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) BackButtonTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetAdditionalOverflowItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetCompactAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetBackButtonDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) CompactAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) ScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetLeadingItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) InitWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) TrailingItemGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetPinnedTrailingGroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SearchController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationItem) SetLeftItemsSupplementBackButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

