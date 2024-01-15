//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSUserNotificationCenterDelegate
*/

type NSUserNotificationCenterDelegate struct {

}

func (r *NSUserNotificationCenterDelegate) UserNotificationCenter() error {
  return fmt.Errorf("unimplemented")
}

