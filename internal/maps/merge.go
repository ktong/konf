// Copyright (c) 2023 The konf authors
// Use of this source code is governed by a MIT license found in the LICENSE file.

package maps

import "strings"

// Merge recursively merges the src map into the dst map.
// Key conflicts are resolved by preferring src,
// or recursively descending, if both values from src and dst are map.
// All keys in dst map are lower case.
func Merge(dst, src map[string]any) {
	for key, srcVal := range src {
		// Ensure key is lower case since the path is case-insensitive.
		key = strings.ToLower(key)

		// Add the srcVal if the key does not exist in the dst map.
		dstVal, exist := dst[key]
		if !exist {
			dst[key] = srcVal

			continue
		}

		// Direct override if the srcVal is not map[string]any.
		srcMap, succeed := srcVal.(map[string]any)
		if !succeed {
			dst[key] = srcVal

			continue
		}

		// Direct override if the dstVal is not map[string]any.
		dstMap, succeed := dstVal.(map[string]any)
		if !succeed {
			dst[key] = srcVal

			continue
		}

		// Merge if the srcVal and dstVal are both map[string]any.
		Merge(dstMap, srcMap)
	}
}
