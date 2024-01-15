//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLAccelerationStructureMotionTriangleGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureMotionTriangleGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTriangleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) TriangleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetTransformationMatrixBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) TransformationMatrixBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) VertexBuffers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) IndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetVertexStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionTriangleGeometryDescriptor) SetIndexType() error {
  return fmt.Errorf("unimplemented")
}

