// Copyright (c) 2023 The konf authors
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build appengine || (!darwin && !dragonfly && !freebsd && !openbsd && !linux && !netbsd && !solaris && !windows)

package file

import "fmt"

func (f File) Watch(context.Context, func(map[string]any)) error {
	f.log(fmt.Sprintf("File.Watch does not supported on %s.", runtime.GOOS))

	return nil
}
