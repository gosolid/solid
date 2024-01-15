//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISearchBar : UIKit.UIView
*/

type UISearchBar struct {
  *UIView

}

func (r *UISearchBar) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SearchBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetScopeButtonTitles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) InputAssistantItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetBarTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SearchFieldBackgroundPositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) Placeholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ScopeButtonTitles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSelectedScopeButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) BackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetScopeBarButtonDividerImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetPrompt() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetShowsSearchResultsButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SearchTextField() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) BarTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) IsTranslucent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ImageForSearchBarIcon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ShowsCancelButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) InputAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetScopeBarButtonBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetScopeBarButtonTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SearchTextPositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetShowsCancelButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetTranslucent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) BarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) Text() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ShowsScopeBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetLookToDictateEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ScopeBarButtonTitleTextAttributesForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetPositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ScopeBarButtonBackgroundImageForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetShowsBookmarkButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ShowsSearchResultsButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSearchResultsButtonSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) IsLookToDictateEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) BackgroundImageForBarPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) PositionAdjustmentForSearchBarIcon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetInputAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ScopeBarBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetScopeBarBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ScopeBarButtonDividerImageForLeftSegmentState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSearchFieldBackgroundPositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSearchFieldBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SearchFieldBackgroundImageForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) ShowsBookmarkButton() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) IsSearchResultsButtonSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSearchBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SelectedScopeButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetSearchTextPositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) SetShowsScopeBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchBar) Prompt() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

