//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionDelegate
*/

type NSURLSessionDelegate struct {

}

func (r *NSURLSessionDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionDelegate) URLSessionDidFinishEventsForBackgroundURLSession() error {
  return fmt.Errorf("unimplemented")
}

