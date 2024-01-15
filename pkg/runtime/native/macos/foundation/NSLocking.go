//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSLocking
*/

type NSLocking struct {

}

func (r *NSLocking) Lock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLocking) Unlock() error {
  return fmt.Errorf("unimplemented")
}

