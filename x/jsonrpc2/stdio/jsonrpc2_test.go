// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdio_test

import (
	"context"
	"io"
	"testing"

	"github.com/goplus/gop/x/jsonrpc2"
	"github.com/goplus/gop/x/jsonrpc2/jsonrpc2test/cases"
	"github.com/goplus/gop/x/jsonrpc2/stdio"
)

func TestStdio(t *testing.T) {
	ctx := context.Background()
	listener := stdio.Listener(true)
	cases.Test(t, ctx, listener, nil, false)
}

func TestDial(t *testing.T) {
	r, w := io.Pipe()
	stdio.Dial(r, w, jsonrpc2.BinderFunc(
		func(ctx context.Context, c *jsonrpc2.Connection) jsonrpc2.ConnectionOptions {
			return jsonrpc2.ConnectionOptions{}
		}), nil)
	if _, err := stdio.Dialer(r, w).Dial(context.Background()); err != stdio.ErrTooManyConnections {
		t.Fatal("stdio.Dialer:", err)
	}
}
