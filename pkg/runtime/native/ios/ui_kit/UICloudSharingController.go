//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICloudSharingController : UIKit.UIViewController
*/

type UICloudSharingController struct {
  *UIViewController

}

func (r *UICloudSharingController) SetAvailablePermissions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) InitWithPreparationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) AvailablePermissions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) InitWithShare() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) ActivityItemSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICloudSharingController) Share() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

