// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"context"
	"encoding/binary"

	"github.com/liquidata-inc/ld/dolt/go/store/hash"
)

// Int is a Noms Value wrapper around the primitive int32 type.
type Int int64

// Value interface
func (v Int) Value(ctx context.Context) Value {
	return v
}

func (v Int) Equals(other Value) bool {
	return v == other
}

func (v Int) Less(nbf *NomsBinFormat, other LesserValuable) bool {
	if v2, ok := other.(Int); ok {
		return v < v2
	}
	return IntKind < other.Kind()
}

func (v Int) Hash(nbf *NomsBinFormat) hash.Hash {
	return getHash(v, nbf)
}

func (v Int) WalkValues(ctx context.Context, cb ValueCallback) {
}

func (v Int) WalkRefs(nbf *NomsBinFormat, cb RefCallback) {
}

func (v Int) typeOf() *Type {
	return IntType
}

func (v Int) Kind() NomsKind {
	return IntKind
}

func (v Int) valueReadWriter() ValueReadWriter {
	return nil
}

func (v Int) writeTo(w nomsWriter, nbf *NomsBinFormat) {
	IntKind.writeTo(w, nbf)
	w.writeInt(v)
}

func (v Int) valueBytes(nbf *NomsBinFormat) []byte {
	// We know the size of the buffer here so allocate it once.
	// IntKind, int (Varint), exp (Varint)
	buff := make([]byte, 1+2*binary.MaxVarintLen64)
	w := binaryNomsWriter{buff, 0}
	v.writeTo(&w, nbf)
	return buff[:w.offset]
}