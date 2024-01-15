//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionTaskDelegate
*/

type NSURLSessionTaskDelegate struct {

}

func (r *NSURLSessionTaskDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

