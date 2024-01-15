//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLConnectionDownloadDelegate
*/

type NSURLConnectionDownloadDelegate struct {

}

func (r *NSURLConnectionDownloadDelegate) Connection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnectionDownloadDelegate) ConnectionDidResumeDownloading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnectionDownloadDelegate) ConnectionDidFinishDownloading() error {
  return fmt.Errorf("unimplemented")
}

