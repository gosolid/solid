//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSUserActivityDelegate
*/

type NSUserActivityDelegate struct {

}

func (r *NSUserActivityDelegate) UserActivity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivityDelegate) UserActivityWillSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserActivityDelegate) UserActivityWasContinued() error {
  return fmt.Errorf("unimplemented")
}

