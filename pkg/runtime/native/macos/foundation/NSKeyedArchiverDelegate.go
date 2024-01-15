//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSKeyedArchiverDelegate
*/

type NSKeyedArchiverDelegate struct {

}

func (r *NSKeyedArchiverDelegate) Archiver() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiverDelegate) ArchiverWillFinish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiverDelegate) ArchiverDidFinish() error {
  return fmt.Errorf("unimplemented")
}

