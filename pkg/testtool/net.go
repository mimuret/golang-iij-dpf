package testtool

import "github.com/mimuret/golang-iij-dpf/pkg/types"

// for testing.
func MustParseIPNet(str string) *types.IPNet {
	n, err := types.ParseIPNet(str)
	if err != nil {
		panic(err)
	}
	return n
}
