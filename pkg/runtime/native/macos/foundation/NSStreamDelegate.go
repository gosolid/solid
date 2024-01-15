//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSStreamDelegate
*/

type NSStreamDelegate struct {

}

func (r *NSStreamDelegate) Stream() error {
  return fmt.Errorf("unimplemented")
}

