//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionStreamDelegate
*/

type NSURLSessionStreamDelegate struct {

}

func (r *NSURLSessionStreamDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

