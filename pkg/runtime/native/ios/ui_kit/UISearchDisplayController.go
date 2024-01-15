//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISearchDisplayController : objc.NSObject
*/

type UISearchDisplayController struct {
  *objc.NSObject

}

func (r *UISearchDisplayController) IsActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchContentsController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchResultsTableView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchResultsDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetSearchResultsDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchResultsDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) DisplaysSearchBarInNavigationBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) NavigationItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) InitWithSearchBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetDisplaysSearchBarInNavigationBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetSearchResultsDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SearchResultsTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchDisplayController) SetSearchResultsTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

