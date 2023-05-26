//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: payload.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package proto

import "errors"
import "github/lwileczek/goBenchmarkSerialization/types/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding SubStruct field model
type FieldModelSubStruct struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Cat *fbe.FieldModelString
    Feeling *fbe.FieldModelString
}

// Create a new SubStruct field model
func NewFieldModelSubStruct(buffer *fbe.Buffer, offset int) *FieldModelSubStruct {
    fbeResult := FieldModelSubStruct{buffer: buffer, offset: offset}
    fbeResult.Cat = fbe.NewFieldModelString(buffer, 4 + 4)
    fbeResult.Feeling = fbe.NewFieldModelString(buffer, fbeResult.Cat.FBEOffset() + fbeResult.Cat.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelSubStruct) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelSubStruct) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Cat.FBESize() +
        fm.Feeling.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelSubStruct) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Cat.FBEExtra() +
        fm.Feeling.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelSubStruct) FBEType() int { return 1 }

// Get the field offset
func (fm *FieldModelSubStruct) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelSubStruct) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelSubStruct) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelSubStruct) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelSubStruct) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelSubStruct) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelSubStruct) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Cat.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Cat.Verify() {
        return false
    }
    fbeCurrentSize += fm.Cat.FBESize()

    if (fbeCurrentSize + fm.Feeling.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Feeling.Verify() {
        return false
    }
    fbeCurrentSize += fm.Feeling.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelSubStruct) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelSubStruct) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelSubStruct) Get() (*SubStruct, error) {
    fbeResult := NewSubStruct()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelSubStruct) GetValue(fbeValue *SubStruct) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelSubStruct) GetFields(fbeValue *SubStruct, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Cat.FBESize()) <= fbeStructSize {
        fbeValue.Cat, _ = fm.Cat.Get()
    } else {
        fbeValue.Cat = ""
    }
    fbeCurrentSize += fm.Cat.FBESize()

    if (fbeCurrentSize + fm.Feeling.FBESize()) <= fbeStructSize {
        fbeValue.Feeling, _ = fm.Feeling.Get()
    } else {
        fbeValue.Feeling = ""
    }
    fbeCurrentSize += fm.Feeling.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelSubStruct) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelSubStruct) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelSubStruct) Set(fbeValue *SubStruct) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelSubStruct) SetFields(fbeValue *SubStruct) error {
    var err error = nil

    if err = fm.Cat.Set(fbeValue.Cat); err != nil {
        return err
    }
    if err = fm.Feeling.Set(fbeValue.Feeling); err != nil {
        return err
    }
    return err
}
