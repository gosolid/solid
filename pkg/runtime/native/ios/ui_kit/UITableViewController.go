//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableViewController : UIKit.UIViewController
*/

type UITableViewController struct {
  *UIViewController

}

func (r *UITableViewController) RefreshControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) TableView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) ClearsSelectionOnViewWillAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) SetClearsSelectionOnViewWillAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) SetRefreshControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) InitWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewController) SetTableView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

