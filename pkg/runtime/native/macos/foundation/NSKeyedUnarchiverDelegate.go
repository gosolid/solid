//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSKeyedUnarchiverDelegate
*/

type NSKeyedUnarchiverDelegate struct {

}

func (r *NSKeyedUnarchiverDelegate) Unarchiver() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiverDelegate) UnarchiverWillFinish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiverDelegate) UnarchiverDidFinish() error {
  return fmt.Errorf("unimplemented")
}

