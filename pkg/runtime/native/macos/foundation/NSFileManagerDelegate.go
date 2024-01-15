//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSFileManagerDelegate
*/

type NSFileManagerDelegate struct {

}

func (r *NSFileManagerDelegate) FileManager() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

