// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package fs

import (
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("fs", "stream.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ReadStream", func (in isolates.FunctionArgs) (*ReadStreamBase, error) {
    return NewReadStream(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ReadStream", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

