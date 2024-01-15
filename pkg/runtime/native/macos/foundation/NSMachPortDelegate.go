//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSMachPortDelegate
*/

type NSMachPortDelegate struct {

}

func (r *NSMachPortDelegate) HandleMachMessage() error {
  return fmt.Errorf("unimplemented")
}

