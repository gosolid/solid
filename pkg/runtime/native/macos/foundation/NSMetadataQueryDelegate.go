//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSMetadataQueryDelegate
*/

type NSMetadataQueryDelegate struct {

}

func (r *NSMetadataQueryDelegate) MetadataQuery() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

