//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSXPCListenerDelegate
*/

type NSXPCListenerDelegate struct {

}

func (r *NSXPCListenerDelegate) Listener() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

