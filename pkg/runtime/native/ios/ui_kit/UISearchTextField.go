//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISearchTextField : UIKit.UITextField
*/

type UISearchTextField struct {
  *UITextField

}

func (r *UISearchTextField) SetAllowsDeletingTokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) TokenBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) SearchSuggestions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) SetSearchSuggestions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) Tokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) TokensInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) ReplaceTextualPortionOfRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) SetTokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) TextualRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) SetTokenBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) AllowsCopyingTokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) SetAllowsCopyingTokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) PositionOfTokenAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) RemoveTokenAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) AllowsDeletingTokens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchTextField) InsertToken() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

