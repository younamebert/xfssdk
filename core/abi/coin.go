package abi

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/younamebert/xfssdk/common"
)

type CTypeUint8 [1]byte
type CTypeBool [1]byte
type CTypeUint16 [2]byte
type CTypeUint32 [4]byte
type CTypeUint64 [8]byte
type CTypeUint256 [32]byte
type CTypeString []byte
type CTypeAddress [25]byte
type CTypeTuple []byte

var (
	CBoolTrue  = CTypeBool{1}
	CBoolFalse = CTypeBool{0}
)

func (t CTypeTuple) String() string {
	return string(t)
}

func (t CTypeTuple) Map() map[string]interface{} {
	result := make(map[string]interface{})
	if err := json.Unmarshal(t, &result); err != nil {
		return nil
	}
	return result
}

func (t CTypeUint8) Uint8() uint8 {
	return t[0]
}

func (t CTypeUint8) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint8) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}

func (t CTypeUint16) Uint16() uint16 {
	return binary.LittleEndian.Uint16(t[:])
}

func (t CTypeUint16) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint16) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint32) Uint32() uint32 {
	return binary.LittleEndian.Uint32(t[:])
}
func (t CTypeUint32) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint32) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint64) Uint64() uint64 {
	return binary.LittleEndian.Uint64(t[:])
}
func (t CTypeUint64) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint64) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint256) BigInt() *big.Int {
	return new(big.Int).SetBytes(t[:])
}
func (t CTypeUint256) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint256) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}

func (t CTypeString) String() string {
	return string(t[:])
}

func (t CTypeString) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeString) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	*t = make([]byte, len(bs))
	copy(*t, bs)
	return
}
func (t CTypeAddress) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}

func (t *CTypeAddress) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeAddress) Address() common.Address {
	return common.Bytes2Address(t[:])
}

func (t CTypeBool) Bool() bool {
	if t[0] == 1 {
		return true
	}
	return false
}

func NewUint8(n uint8) CTypeUint8 {
	return CTypeUint8{n}
}

func NewUint16(n uint16) (m CTypeUint16) {
	binary.LittleEndian.PutUint16(m[:], n)
	return
}

func NewUint32(n uint32) (m CTypeUint32) {
	binary.LittleEndian.PutUint32(m[:], n)
	return
}
func NewUint64(n uint64) (m CTypeUint64) {
	binary.LittleEndian.PutUint64(m[:], n)
	return
}

func NewUint256(n *big.Int) (m CTypeUint256) {
	bs := n.Bytes()
	r := len(m) - len(bs)
	copy(m[r:], bs)
	return
}

func NewAddress(n common.Address) (m CTypeAddress) {
	copy(m[:], n[:])
	return
}

func NewTuple(n map[string]interface{}) (m CTypeTuple) {
	bs, err := json.Marshal(n)
	if err != nil {
		return nil
	}
	copy(m[:], bs[:])
	return
}

type Heap struct {
	Address  string         `json:"address"`
	TokenId  string         `json:"tokenid"`
	Children []HeapChildren `json:"children"`
}
type HeapChildren struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func DncodeCTypeTuple(n string) ([]Heap, error) {
	n = strings.TrimPrefix(n, "0x")
	ast, err := hex.DecodeString(n)
	if err != nil {
		return nil, err
	}
	pack := make(map[string]interface{}, 0)
	if err := json.Unmarshal(ast, &pack); err != nil {
		return nil, err
	}
	result := make([]Heap, 0)
	for key, val := range pack {
		var heapResp Heap
		resultChildren := make([]HeapChildren, 0)
		keyBatch := strings.Split(key, ",")
		heapResp.Address = keyBatch[0]
		heapResp.TokenId = keyBatch[1]
		if len(keyBatch) > 0 {
			heap := val.(map[string]interface{})
			for keys, vals := range heap {
				var heapChildrenResp HeapChildren
				heapChildren := vals.(map[string]interface{})
				heapChildrenkey := heapChildren["type"].(string)
				heapChildrenVal := heapChildren["value"].(string)
				ens, err := Decode(heapChildrenkey, heapChildrenVal)
				if err != nil {
					return nil, err
				}
				heapChildrenResp.Name = keys
				heapChildrenResp.Value = ens
				resultChildren = append(resultChildren, heapChildrenResp)
			}
			heapResp.Children = resultChildren
		}
		result = append(result, heapResp)
	}
	return result, nil
}

func Decode(heapkey, heapVal string) (interface{}, error) {
	var result interface{}
	switch heapkey {
	case "CTypeString":
		para := CTypeString(heapVal)
		pars, err := hex.DecodeString(para.String())
		if err != nil {
			return "", fmt.Errorf("ctypestring to string:%v", heapVal)
		}
		result = string(pars)
	case "CTypeUint8":
		big8, ok := big.NewInt(0).SetString(heapVal, 16)
		if !ok {
			return "", fmt.Errorf("ctypeuint256 to uint256:%v", heapVal)
		}
		para := NewUint8(uint8(big8.Uint64()))
		result = para.Uint8()
	case "CTypeUint256":
		big256, ok := big.NewInt(0).SetString(heapVal, 16)
		if !ok {
			return "", fmt.Errorf("ctypeuint256 to uint256:%v", heapVal)
		}
		para := NewUint256(big256)
		result = para.BigInt().String()
	case "CTypeAddress":
		addrBytes, err := hex.DecodeString(heapVal)
		if err != nil {
			return "", fmt.Errorf("ctypeaddress to address:%v", err)
		}
		para := NewAddress(common.Bytes2Address(addrBytes))
		address := para.Address()
		result = address.B58String()
	case "CTypeBool":
		// para := heapVal.(CTypeBool)
		// result = para.Bool()
	default:
		return "", fmt.Errorf("type check err")
	}
	return result, nil
}
