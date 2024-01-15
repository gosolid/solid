//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol ApplicationServices.PDEPanel
*/

type PDEPanel struct {

}

func (r *PDEPanel) SummaryInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) SaveValuesAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) PanelView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) PanelName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) PanelKind() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) ShouldPrint() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) WillShow() error {
  return fmt.Errorf("unimplemented")
}

func (r *PDEPanel) ShouldHide() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) RestoreValuesAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) SupportedPPDOptionKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) ShouldShowHelp() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *PDEPanel) PrintWindowWillClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *PDEPanel) PPDOptionKeyValueDidChange() error {
  return fmt.Errorf("unimplemented")
}

