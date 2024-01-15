//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureTriangleGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureTriangleGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) IndexBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) IndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) VertexBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) VertexFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) TransformationMatrixBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) VertexStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetTriangleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetVertexBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) IndexBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) TriangleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureTriangleGeometryDescriptor) VertexBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

