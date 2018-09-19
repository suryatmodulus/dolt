package types

import (
	"github.com/attic-labs/noms/go/hash"
	"github.com/google/uuid"
)

const (
	uuidNumBytes = 16
)

type UUID uuid.UUID

func (v UUID) Value() Value {
	return v
}

func (v UUID) Equals(other Value) bool {
	return v == other
}

func (v UUID) Less(other Value) bool {
	if v2, ok := other.(UUID); ok {
		for i := 0; i < uuidNumBytes; i++ {
			b1 := v[i]
			b2 := v2[i]

			if b1 < b2 {
				return true
			}
		}

		return false
	}
	return UUIDKind < other.Kind()
}

func (v UUID) Hash() hash.Hash {
	return getHash(v)
}

func (v UUID) WalkValues(cb ValueCallback) {
}

func (v UUID) WalkRefs(cb RefCallback) {
}

func (v UUID) typeOf() *Type {
	return UUIDType
}

func (v UUID) Kind() NomsKind {
	return UUIDKind
}

func (v UUID) valueReadWriter() ValueReadWriter {
	return nil
}

func (v UUID) writeTo(w nomsWriter) {
	id := UUID(v)
	byteSl := id[:]
	UUIDKind.writeTo(w)
	w.writeBytes(byteSl)
}

func (v UUID) valueBytes() []byte {
	return v[:]
}