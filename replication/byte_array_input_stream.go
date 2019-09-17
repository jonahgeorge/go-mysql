package replication

import "bytes"

type ByteArrayInputStream struct {
	bytes.Buffer
}

func (s *ByteArrayInputStream) ReadPackedNumber() (int, error)           { return 0, nil }
func (s *ByteArrayInputStream) ReadIntegerPairs()                        {}
func (s *ByteArrayInputStream) ReadLengthEncodedString() (string, error) { return "", nil }
func (s *ByteArrayInputStream) ReadPackedInteger() (int, error)          { return 0, nil }
func (s *ByteArrayInputStream) ReadInteger() (int, error)                { return 0, nil }
func (s *ByteArrayInputStream) ReadString(length int) (string, error)    { return "", nil }
