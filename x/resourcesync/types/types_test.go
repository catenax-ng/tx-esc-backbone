package types

import "github.com/catenax/esc-backbone/testutil"

const (
	Alice = testutil.Alice
	Bob   = testutil.Bob
	Carol = testutil.Carol
)

func createValidResouceKey(originator string, origResId string) ResourceKey {
	resourceKey, err := NewResourceKey(originator, origResId)
	if err != nil {
		panic(err)
	}
	return resourceKey
}
