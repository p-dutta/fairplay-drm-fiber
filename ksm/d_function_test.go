package ksm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHashValue struct {
	in  []byte
	out []byte
}

var testsData = []testHashValue{
	{
		[]byte{0xa4, 0xc4, 0x11, 0xc5, 0x4d, 0x94, 0x72, 0x71, 0x43, 0x50, 0x4a, 0xec, 0xe5, 0x61, 0x3d, 0xa8, 0xc6, 0xee, 0x6d, 0xd2},
		[]byte{0x3a, 0x24, 0xa0, 0x68, 0xc1, 0x16, 0x84, 0x2f, 0x10, 0xfa, 0xc5, 0xc8, 0x6c, 0xb1, 0xcb, 0x5c},
	},
	{
		[]byte{0xee, 0x81, 0x0d, 0xd5, 0x77, 0x9c, 0xe5, 0x0a, 0xa8, 0xdf, 0x38, 0x24, 0xb3, 0xf2, 0x59, 0x4c, 0xf9, 0xee, 0x68, 0xa9},
		[]byte{0x93, 0xc8, 0x2d, 0xc4, 0xd1, 0xc4, 0xa3, 0x38, 0x89, 0x6b, 0xfc, 0xe2, 0xe7, 0xdc, 0x12, 0xa0},
	},
	{
		[]byte{0x9a, 0x06, 0x18, 0xdb, 0x8a, 0x4e, 0xcc, 0x1c, 0x8f, 0x1f, 0xa8, 0xe9, 0x48, 0x2d, 0xbb, 0xe5, 0x66, 0x47, 0xaf, 0xe1},
		[]byte{0x73, 0xb8, 0x44, 0x5f, 0xd6, 0x39, 0x46, 0x21, 0x54, 0xf3, 0x50, 0x3c, 0x6e, 0xa3, 0x2c, 0x89},
	},
	{
		[]byte{0x30, 0x38, 0x38, 0xa2, 0xb4, 0xca, 0xb6, 0x6f, 0x25, 0xd5, 0x3b, 0xd3, 0x31, 0xb5, 0xee, 0x4c, 0x03, 0x5b, 0x6e, 0x12},
		[]byte{0xc7, 0x4f, 0x35, 0x70, 0x4e, 0xc5, 0x8c, 0x08, 0x82, 0xbc, 0x40, 0x14, 0xa7, 0x41, 0x52, 0xa8},
	},
	{
		[]byte{0xa4, 0xc4, 0x11, 0xc5, 0x4d, 0x94, 0x72, 0x71, 0x43, 0x50, 0x4a, 0xec, 0xe5, 0x61, 0x3d, 0xa8, 0xc6, 0xee, 0x6d, 0xd2, 0x6b, 0xa5, 0xd0, 0x7e, 0xc0, 0xdc, 0x44, 0xc8, 0x4e, 0xc4},
		[]byte{0x4b, 0x46, 0xab, 0x65, 0xbd, 0x04, 0xaf, 0x0d, 0xd7, 0x21, 0x71, 0x8a, 0x68, 0x66, 0x33, 0x90},
	},
	{
		[]byte{0xee, 0x81, 0x0d, 0xd5, 0x77, 0x9c, 0xe5, 0x0a, 0xa8, 0xdf, 0x38, 0x24, 0xb3, 0xf2, 0x59, 0x4c, 0xf9, 0xee, 0x68, 0xa9, 0xc6, 0xca, 0xa7, 0xa3, 0x67, 0x09, 0xf5, 0x54, 0x94, 0xd5},
		[]byte{0xc6, 0x3d, 0x0c, 0x63, 0x0b, 0xcd, 0x57, 0x9d, 0x07, 0x29, 0x01, 0xdb, 0xa8, 0x2a, 0xe2, 0xa6},
	},
	{
		[]byte{0x9a, 0x06, 0x18, 0xdb, 0x8a, 0x4e, 0xcc, 0x1c, 0x8f, 0x1f, 0xa8, 0xe9, 0x48, 0x2d, 0xbb, 0xe5, 0x66, 0x47, 0xaf, 0xe1, 0x07, 0x39, 0xe4, 0xbd, 0x00, 0xd2, 0x9d, 0xb0, 0xd9, 0xcc},
		[]byte{0x2d, 0x46, 0x7d, 0x6f, 0xc4, 0xaa, 0xb3, 0x09, 0xe8, 0xb4, 0x61, 0x82, 0x47, 0x56, 0x7e, 0x06},
	},
	{
		[]byte{0x30, 0x38, 0x38, 0xa2, 0xb4, 0xca, 0xb6, 0x6f, 0x25, 0xd5, 0x3b, 0xd3, 0x31, 0xb5, 0xee, 0x4c, 0x03, 0x5b, 0x6e, 0x12, 0xaf, 0xcd, 0x05, 0xc5, 0xa5, 0xa8, 0x5a, 0x38, 0x2f, 0xf4},
		[]byte{0x35, 0x3d, 0x49, 0x1e, 0xdc, 0x8f, 0x6f, 0x81, 0x91, 0xa3, 0xcb, 0x36, 0x6b, 0xaf, 0xe1, 0x4c},
	},
	{
		[]byte{0xa4, 0xc4, 0x11, 0xc5, 0x4d, 0x94, 0x72, 0x71, 0x43, 0x50, 0x4a, 0xec, 0xe5, 0x61, 0x3d, 0xa8, 0xc6, 0xee, 0x6d, 0xd2, 0x6b, 0xa5, 0xd0, 0x7e, 0xc0, 0xdc, 0x44, 0xc8, 0x4e, 0xc4, 0x81, 0x40, 0xb2, 0x44, 0xd5, 0xbe, 0x32, 0x3d, 0x25, 0x1a},
		[]byte{0x84, 0x39, 0xa3, 0x6b, 0x35, 0x9c, 0x7c, 0x1c, 0x5b, 0x69, 0x8c, 0x59, 0x4b, 0x56, 0xb0, 0x20},
	},
	{
		[]byte{0xee, 0x81, 0x0d, 0xd5, 0x77, 0x9c, 0xe5, 0x0a, 0xa8, 0xdf, 0x38, 0x24, 0xb3, 0xf2, 0x59, 0x4c, 0xf9, 0xee, 0x68, 0xa9, 0xc6, 0xca, 0xa7, 0xa3, 0x67, 0x09, 0xf5, 0x54, 0x94, 0xd5, 0xc5, 0x69, 0x71, 0xd3, 0x90, 0xbc, 0x57, 0xa6, 0x82, 0x23},
		[]byte{0x43, 0x0b, 0x6d, 0xb2, 0xce, 0x11, 0xce, 0x6b, 0xcd, 0x70, 0x76, 0x18, 0x7c, 0x44, 0xaa, 0x1b},
	},
	{
		[]byte{0x30, 0x38, 0x38, 0xa2, 0xb4, 0xca, 0xb6, 0x6f, 0x25, 0xd5, 0x3b, 0xd3, 0x31, 0xb5, 0xee, 0x4c, 0x03, 0x5b, 0x6e, 0x12, 0xaf, 0xcd, 0x05, 0xc5, 0xa5, 0xa8, 0x5a, 0x38, 0x2f, 0xf4, 0xbf, 0xae, 0x23, 0x08, 0x44, 0xae, 0xf7, 0xe2, 0xae, 0x5c},
		[]byte{0xa7, 0x79, 0x94, 0xae, 0xda, 0x51, 0xdf, 0x86, 0x8c, 0xd8, 0x2b, 0xd3, 0xb1, 0x13, 0x14, 0x01},
	},
	{
		[]byte{0x9a, 0x06, 0x18, 0xdb, 0x8a, 0x4e, 0xcc, 0x1c, 0x8f, 0x1f, 0xa8, 0xe9, 0x48, 0x2d, 0xbb, 0xe5, 0x66, 0x47, 0xaf, 0xe1, 0x07, 0x39, 0xe4, 0xbd, 0x00, 0xd2, 0x9d, 0xb0, 0xd9, 0xcc, 0xf4, 0xba, 0x5d, 0x55, 0x6f, 0x4c, 0x86, 0x9e, 0xbb, 0x9e},
		[]byte{0xfc, 0xff, 0x3f, 0xf1, 0x9e, 0x29, 0x11, 0xe5, 0x2c, 0xe2, 0x52, 0xe8, 0x5f, 0x04, 0x98, 0xda},
	},
	{
		[]byte{0xa4, 0xc4, 0x11, 0xc5, 0x4d, 0x94, 0x72, 0x71, 0x43, 0x50, 0x4a, 0xec, 0xe5, 0x61, 0x3d, 0xa8, 0xc6, 0xee, 0x6d, 0xd2, 0x6b, 0xa5, 0xd0, 0x7e, 0xc0, 0xdc, 0x44, 0xc8, 0x4e, 0xc4, 0x81, 0x40, 0xb2, 0x44, 0xd5, 0xbe, 0x32, 0x3d, 0x25, 0x1a, 0xcb, 0x3a, 0x4b, 0xb1, 0x98, 0x98, 0x89, 0x19, 0x73, 0xdb},
		[]byte{0xf2, 0x27, 0x2c, 0x25, 0x45, 0x04, 0xfc, 0x4c, 0xe7, 0xaf, 0x28, 0x73, 0xc1, 0xde, 0x18, 0x4f},
	},
	{
		[]byte{0xee, 0x81, 0x0d, 0xd5, 0x77, 0x9c, 0xe5, 0x0a, 0xa8, 0xdf, 0x38, 0x24, 0xb3, 0xf2, 0x59, 0x4c, 0xf9, 0xee, 0x68, 0xa9, 0xc6, 0xca, 0xa7, 0xa3, 0x67, 0x09, 0xf5, 0x54, 0x94, 0xd5, 0xc5, 0x69, 0x71, 0xd3, 0x90, 0xbc, 0x57, 0xa6, 0x82, 0x23, 0xd5, 0x88, 0xed, 0xb0, 0xe4, 0x8e, 0xc0, 0x73, 0xd2, 0x8b},
		[]byte{0x0f, 0x25, 0xa9, 0xbf, 0x94, 0x0d, 0xac, 0xa0, 0xa5, 0x20, 0x93, 0x7b, 0xe1, 0x75, 0x2f, 0x75},
	},
	{
		[]byte{0x9a, 0x06, 0x18, 0xdb, 0x8a, 0x4e, 0xcc, 0x1c, 0x8f, 0x1f, 0xa8, 0xe9, 0x48, 0x2d, 0xbb, 0xe5, 0x66, 0x47, 0xaf, 0xe1, 0x07, 0x39, 0xe4, 0xbd, 0x00, 0xd2, 0x9d, 0xb0, 0xd9, 0xcc, 0xf4, 0xba, 0x5d, 0x55, 0x6f, 0x4c, 0x86, 0x9e, 0xbb, 0x9e, 0xce, 0xeb, 0x01, 0x9a, 0xca, 0xa7, 0xae, 0x6c, 0x8c, 0xa7},
		[]byte{0x10, 0xe4, 0x1b, 0x10, 0xdf, 0x29, 0xf8, 0xc2, 0xea, 0x91, 0xfe, 0x7d, 0x12, 0x6f, 0x50, 0xe1},
	},
	{
		[]byte{0x30, 0x38, 0x38, 0xa2, 0xb4, 0xca, 0xb6, 0x6f, 0x25, 0xd5, 0x3b, 0xd3, 0x31, 0xb5, 0xee, 0x4c, 0x03, 0x5b, 0x6e, 0x12, 0xaf, 0xcd, 0x05, 0xc5, 0xa5, 0xa8, 0x5a, 0x38, 0x2f, 0xf4, 0xbf, 0xae, 0x23, 0x08, 0x44, 0xae, 0xf7, 0xe2, 0xae, 0x5c, 0x83, 0x46, 0x71, 0xc1, 0x63, 0x2a, 0xfc, 0xa9, 0x43, 0xeb},
		[]byte{0xb1, 0xca, 0xec, 0xdd, 0xef, 0x43, 0x3c, 0xeb, 0x4d, 0x8e, 0xd5, 0x54, 0xf1, 0x70, 0x2d, 0xbd},
	},
}

func TestDFunctionComputeHashValue(t *testing.T) {
	d := DFunction{}
	for _, test := range testsData {

		actualOut, err := d.ComputeHashValue(test.in)
		assert.NoError(t, err)

		assert.Equal(t, actualOut, test.out)
	}
}
