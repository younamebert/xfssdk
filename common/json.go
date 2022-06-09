package common

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func MarshalIndent(val interface{}) ([]byte, error) {
	return json.MarshalIndent(val, "", "    ")
}

func SortAndEncodeMap(data map[string]string) string {
	mapkeys := make([]string, 0)
	for k := range data {
		mapkeys = append(mapkeys, k)
	}
	sort.Strings(mapkeys)
	strbuf := ""
	for i, key := range mapkeys {
		val := data[key]
		if val == "" {
			continue
		}
		strbuf += fmt.Sprintf("%s=%s", key, val)
		if i < len(mapkeys)-1 {
			strbuf += "&"
		}
	}
	return strbuf
}

func Struct2Bytes(iter interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, iter)
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

func Bytes2Struct(buf []byte) unsafe.Pointer {
	return unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&buf)).Data,
	)
}

func Int2Byte(data int) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

func Byte2Int(data []byte) int {
	var ret int = 0
	var len int = len(data)
	var i uint = 0
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}
