//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSPortDelegate
*/

type NSPortDelegate struct {

}

func (r *NSPortDelegate) HandlePortMessage() error {
  return fmt.Errorf("unimplemented")
}

