package rgbpp

import (
	"bytes"
	"encoding/binary"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

const HashTypeData2 types.ScriptHashType = "data2"

func PackCkbScript(script *types.Script) *Script {
	codeHash := Byte32FromSliceUnchecked(script.CodeHash.Bytes())
	HashType := script.HashType
	args := script.Args

	builder := NewScriptBuilder()
	builder.CodeHash(*codeHash)
	builder.HashType(*PackHashType(HashType))
	builder.Args(*PackBytes(args))
	b := builder.Build()

	return &b
}

func PackBytes(v []byte) *Bytes {
	builder := NewBytesBuilder()
	for _, vv := range v {
		builder.Push(*PackByte(vv))
	}
	b := builder.Build()
	return &b
}

func PackHashType(t types.ScriptHashType) *Byte {
	var b byte
	switch t {
	case types.HashTypeData:
		b = 0x00
	case types.HashTypeType:
		b = 0x01
	case types.HashTypeData1:
		b = 0x02
	case HashTypeData2:
		b = 0x04
	default:
		return nil
	}

	return PackByte(b)
}

func PackByte(v byte) *Byte {
	b := NewByte(v)

	return &b
}

func PackAfter(after int64) (Uint32, error) {
	val := int32(after)
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, val)
	if err != nil {
		return Uint32Default(), err
	}

	builder := NewUint32Builder()

	var b [4]Byte

	for idx, item := range buf.Bytes() {
		b[idx] = *PackByte(item)
	}

	builder.Set(b)

	return builder.Build(), nil
}
