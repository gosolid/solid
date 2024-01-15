//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSProgressReporting
*/

type NSProgressReporting struct {

}

func (r *NSProgressReporting) Progress() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

