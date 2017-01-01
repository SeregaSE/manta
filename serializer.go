package manta

import (
	"fmt"
	"strings"
)

type serializer struct {
	name    string
	version int32
	fields  []*field
}

func (s *serializer) id() string {
	return serializerId(s.name, s.version)
}

func (s *serializer) getNameForFieldPath(fp *fieldPath, pos int) []string {
	return s.fields[fp.path[pos]].getNameForFieldPath(fp, pos+1)
}

func (s *serializer) getTypeForFieldPath(fp *fieldPath, pos int) *fieldType {
	return s.fields[fp.path[pos]].getTypeForFieldPath(fp, pos+1)
}

func (s *serializer) getDecoderForFieldPath(fp *fieldPath, pos int) fieldDecoder {
	return s.fields[fp.path[pos]].getDecoderForFieldPath(fp, pos+1)
}

func (s *serializer) getFieldForFieldPath(fp *fieldPath, pos int) *field {
	return s.fields[fp.path[pos]].getFieldForFieldPath(fp, pos+1)
}

func (s *serializer) getFieldPathForName(fp *fieldPath, name string) bool {
	for i, f := range s.fields {
		if name == f.varName {
			fp.path[fp.last] = i
			return true
		}

		if strings.HasPrefix(name, f.varName+".") {
			fp.path[fp.last] = i
			fp.last++
			return f.getFieldPathForName(fp, name[len(f.varName)+1:])
		}
	}

	return false
}

func serializerId(name string, version int32) string {
	return fmt.Sprintf("%s(%d)", name, version)
}