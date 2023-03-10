// Copyright (c) 2023 The konf authors
// Use of this source code is governed by a MIT license found in the LICENSE file.

package flag_test

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"

	kflag "github.com/ktong/konf/provider/flag"
)

func BenchmarkNew(b *testing.B) {
	set := &flag.FlagSet{}
	set.String("k", "v", "")
	b.ResetTimer()

	var loader kflag.Flag
	for i := 0; i < b.N; i++ {
		loader = kflag.New(kflag.WithFlagSet(set))
	}
	b.StopTimer()

	values, err := loader.Load()
	require.NoError(b, err)
	require.Equal(b, "v", values["k"])
}

func BenchmarkLoad(b *testing.B) {
	set := &flag.FlagSet{}
	set.String("k", "v", "")
	loader := kflag.New(kflag.WithFlagSet(set))
	b.ResetTimer()

	var (
		values map[string]any
		err    error
	)
	for i := 0; i < b.N; i++ {
		values, err = loader.Load()
	}
	b.StopTimer()

	require.NoError(b, err)
	require.Equal(b, "v", values["k"])
}
