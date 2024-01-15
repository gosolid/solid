//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSFastEnumeration
*/

type NSFastEnumeration struct {

}

func (r *NSFastEnumeration) CountByEnumeratingWithState() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

