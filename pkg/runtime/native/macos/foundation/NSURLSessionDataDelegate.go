//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionDataDelegate
*/

type NSURLSessionDataDelegate struct {

}

func (r *NSURLSessionDataDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

