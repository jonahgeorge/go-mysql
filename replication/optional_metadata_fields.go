package replication

import (
	"math/big"

	"github.com/pingcap/errors"
)

type OptionalMetadataFieldType uint

// https://github.com/shyiko/mysql-binlog-connector-java/pull/251/files
const (
	METADATA_FIELD_TYPE_SIGNEDNESS              OptionalMetadataFieldType = 1 << iota // UNSIGNED flag of numeric columns
	METADATA_FIELD_TYPE_DEFAULT_CHARSET                                               // Default character set of string columns
	METADATA_FIELD_TYPE_COLUMN_CHARSET                                                // Character set of string columns
	METADATA_FIELD_TYPE_COLUMN_NAME                                                   //
	METADATA_FIELD_TYPE_SET_STR_VALUE                                                 // String value of SET columns
	METADATA_FIELD_TYPE_ENUM_STR_VALUE                                                // String value of ENUM columns
	METADATA_FIELD_TYPE_GEOMETRY_TYPE                                                 // Real type of geometry columns
	METADATA_FIELD_TYPE_SIMPLE_PRIMARY_KEY                                            // Primary key without prefix
	METADATA_FIELD_TYPE_PRIMARY_KEY_WITH_PREFIX                                       // Primary key with prefix
)

type OptionalMetadataFields struct {
	// // Content of DEFAULT_CHARSET field is converted into Default_charset.
	// Default_charset m_default_charset

	Signedness              []bool
	ColumnCharSet           []uint
	EnumAndSetColumnCharset []uint
	ColumnName              []string
	EnumStrValue            []string
	SetStrValue             []string
	GeometryType            []uint
	// /*
	//    The uint_pair means <column index, prefix length>.  Prefix length is 0 if
	//    whole column value is used.
	// */
	// m_primary_key []uint_pair

	IsValid bool
}

func (o *OptionalMetadataFields) Decode(buf []byte) error {
	if len(buf) == 0 {
		return nil
	}

	idx := 0
	for {
		if idx >= len(buf) {
			break
		}

		fieldType := getFieldType()
		fieldLength := getFieldLength()
		_ = fieldLength

		switch fieldType {
		//case SIGNEDNESS:
		//	o.signedness = parseSignedness()
		//case COLUMN_CHARSET:
		//	o.columnCharSet = parseColumnCharSet()
		case COLUMN_NAME:
		//	o.columnName = parseColumnName()
		default:
			return errors.Errorf("unknown OptionalMetadataFieldType: %v", fieldType)
		}

	}

	return nil
}

func getFieldType() OptionalMetadataFieldType {
	return COLUMN_NAME
}

func getFieldLength() uint {
	return 0
}

func ReadSignedness(stream *ByteArrayInputStream, length int) (*big.Int, error) {
	result := big.NewInt()
	// according to MySQL internals the amount of storage required for N columns is INT((N+7)/8) bytes
	buf := stream.Read((length + 7) >> 3)

	return result, nil
}

func readDefaultCharset(field []byte) {}

func ReadIntegers(stream *ByteArrayInputStream) ([]int, error) {
	var result []int
	for stream.Len() >= 0 {
		i, err := stream.ReadPackedInteger()
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func ReadColumnNames(stream *ByteArrayInputStream) ([]string, error) {
	var columnNames []string
	for stream.Len() >= 0 {
		columnName, err := stream.ReadLengthEncodedString()
		if err != nil {
			return nil, err
		}
		columnNames = append(columnNames, columnName)
	}
	return columnNames, nil
}

func ReadTypeValues(stream *ByteArrayInputStream) [][]string { return nil }

func ReadIntegerPairs(stream *ByteArrayInputStream) (map[int]int, error) {
	var result map[int]int
	for stream.Len() >= 0 {
		columnIndex, err := stream.ReadPackedInteger()
		if err != nil {
			return nil, err
		}
		columnCharset, err := stream.ReadPackedInteger()
		if err != nil {
			return nil, err
		}
		result[columnIndex] = columnCharset
	}
	return result, nil
}

