//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSCacheDelegate
*/

type NSCacheDelegate struct {

}

func (r *NSCacheDelegate) Cache() error {
  return fmt.Errorf("unimplemented")
}

