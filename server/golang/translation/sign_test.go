package translation

import (
	"testing"
)

// According to https://api.fanyi.baidu.com/product/113#
func TestSigning(t *testing.T) {
	s := sign("apple", "1435660288", "2015063000000001", "12345678")
	if s != "f89f9594663708c1605f3d736d01d2d4" {
		t.Fatal("MD5 not match")
	}
}
