//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionDownloadDelegate
*/

type NSURLSessionDownloadDelegate struct {

}

func (r *NSURLSessionDownloadDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

