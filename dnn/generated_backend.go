package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// Backend is a representation of cudnnBackendDescriptor_t.
type Backend struct {
	internal C.cudnnBackendDescriptor_t

	attributeName   BackendAttributeName
	attributeType   BackendAttributeType
	elementCount    int64
	arrayOfElements Memory
}

// NewBackend creates a new Backend.
func NewBackend(attributeName BackendAttributeName, attributeType BackendAttributeType, elementCount int64, arrayOfElements Memory) (retVal *Backend, err error) {
	var internal C.cudnnBackendDescriptor_t
	if err := result(C.cudnnBackendCreateDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnBackendSetAttribute(internal, attributeName.C(), attributeType.C(), C.int64_t(elementCount), unsafe.Pointer(arrayOfElements.Uintptr()))); err != nil {
		return nil, err
	}

	retVal = &Backend{
		internal:        internal,
		attributeName:   attributeName,
		attributeType:   attributeType,
		elementCount:    elementCount,
		arrayOfElements: arrayOfElements,
	}
	runtime.SetFinalizer(retVal, destroyBackend)
	return retVal, nil
}

// C() returns the internal cgo representation
func (b *Backend) C() C.cudnnBackendDescriptor_t { return b.internal }

// AttributeName returns the internal attributeName.
func (b *Backend) AttributeName() BackendAttributeName { return b.attributeName }

// AttributeType returns the internal attributeType.
func (b *Backend) AttributeType() BackendAttributeType { return b.attributeType }

// ElementCount returns the internal elementCount parameter.
func (b *Backend) ElementCount() int64 { return b.elementCount }

// ArrayOfElements returns the internal arrayOfElements.
func (b *Backend) ArrayOfElements() Memory { return b.arrayOfElements }

func destroyBackend(obj *Backend) { C.cudnnBackendDestroyDescriptor(obj.internal) }
