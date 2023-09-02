/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package build_test

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/x/build"
)

var (
	ctx = build.Default()
)

func init() {
	ctx.LoadConfig = func(cfg *cl.Config) {
		cfg.NoFileLine = true
	}
	build.RegisterClassFileType(".tspx", "MyGame", []*build.Class{
		{Ext: ".tspx", Class: "Sprite"},
	}, "github.com/goplus/gop/cl/internal/spx")
}

func gopClTest(t *testing.T, gopcode interface{}, expected string) {
	gopClTestEx(t, "main.gop", gopcode, expected)
}

func gopClTestEx(t *testing.T, filename string, gopcode interface{}, expected string) {
	data, err := ctx.BuildFile(filename, gopcode)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != expected {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

func testKind(t *testing.T, name string, proj, class bool) {
	isProj, ok := build.ClassKind(name)
	if isProj != proj || ok != class {
		t.Fatal("check classkind failed", name, isProj, ok)
	}
}

func TestKind(t *testing.T) {
	testKind(t, "Cat.gox", false, false)
	testKind(t, "Cat.spx", false, true)
	testKind(t, "main.spx", true, true)
	testKind(t, "main.gmx", true, true)
	testKind(t, "Cat.tspx", false, true)
	testKind(t, "main.tspx", true, true)
}

func TestGop(t *testing.T) {
	var src = `
println "Go+"
`
	var expect = `package main

import fmt "fmt"

func main() {
	fmt.Println("Go+")
}
`
	gopClTest(t, src, expect)
	gopClTest(t, []byte(src), expect)
	gopClTest(t, bytes.NewBufferString(src), expect)
	gopClTestEx(t, `./_testdata/hello/main.gop`, nil, expect)

	f, err := os.Open("./_testdata/hello/main.gop")
	if err != nil {
		t.Fatal("open failed", err)
	}
	defer f.Close()
	gopClTest(t, f, expect)
}

func TestGox(t *testing.T) {
	gopClTestEx(t, "Rect.gox", `
println "Go+"
`, `package main

import fmt "fmt"

type Rect struct {
}

func (this *Rect) Main() {
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import fmt "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	Buffer
	v int
}

func (this *Rect) Main() {
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	*Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import fmt "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	*Buffer
	v int
}

func (this *Rect) Main() {
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	*bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	fmt "fmt"
	bytes "bytes"
)

type Rect struct {
	*bytes.Buffer
	v int
}

func (this *Rect) Main() {
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	fmt "fmt"
	bytes "bytes"
)

type Rect struct {
	bytes.Buffer
	v int
}

func (this *Rect) Main() {
	fmt.Println("Go+")
}
func main() {
}
`)
}

func TestBig(t *testing.T) {
	gopClTest(t, `
a := 1/2r
println a+1/2r
`, `package main

import (
	fmt "fmt"
	ng "github.com/goplus/gop/builtin/ng"
	big "math/big"
)

func main() {
	a := ng.Bigrat_Init__2(big.NewRat(1, 2))
	fmt.Println(a.Gop_Add(ng.Bigrat_Init__2(big.NewRat(1, 2))))
}
`)
}

func TestIoxLines(t *testing.T) {
	gopClTest(t, `
import "io"

var r io.Reader

for line <- lines(r) {
	println line
}
`, `package main

import (
	fmt "fmt"
	iox "github.com/goplus/gop/builtin/iox"
	io "io"
)

var r io.Reader

func main() {
	for _gop_it := iox.Lines(r).Gop_Enum(); ; {
		var _gop_ok bool
		line, _gop_ok := _gop_it.Next()
		if !_gop_ok {
			break
		}
		fmt.Println(line)
	}
}
`)
}

func TestErrorWrap(t *testing.T) {
	gopClTest(t, `
import (
    "strconv"
)

func add(x, y string) (int, error) {
    return strconv.Atoi(x)? + strconv.Atoi(y)?, nil
}

func addSafe(x, y string) int {
    return strconv.Atoi(x)?:0 + strconv.Atoi(y)?:0
}

println add("100", "23")!

sum, err := add("10", "abc")
println sum, err

println addSafe("10", "abc")
`, `package main

import (
	fmt "fmt"
	strconv "strconv"
	errors "github.com/qiniu/x/errors"
)

func add(x string, y string) (int, error) {
	var _autoGo_1 int
	{
		var _gop_err error
		_autoGo_1, _gop_err = strconv.Atoi(x)
		if _gop_err != nil {
			_gop_err = errors.NewFrame(_gop_err, "strconv.Atoi(x)", "main.gop", 7, "main.add")
			return 0, _gop_err
		}
		goto _autoGo_2
	_autoGo_2:
	}
	var _autoGo_3 int
	{
		var _gop_err error
		_autoGo_3, _gop_err = strconv.Atoi(y)
		if _gop_err != nil {
			_gop_err = errors.NewFrame(_gop_err, "strconv.Atoi(y)", "main.gop", 7, "main.add")
			return 0, _gop_err
		}
		goto _autoGo_4
	_autoGo_4:
	}
	return _autoGo_1 + _autoGo_3, nil
}
func addSafe(x string, y string) int {
	return func() (_gop_ret int) {
		var _gop_err error
		_gop_ret, _gop_err = strconv.Atoi(x)
		if _gop_err != nil {
			return 0
		}
		return
	}() + func() (_gop_ret int) {
		var _gop_err error
		_gop_ret, _gop_err = strconv.Atoi(y)
		if _gop_err != nil {
			return 0
		}
		return
	}()
}
func main() {
	fmt.Println(func() (_gop_ret int) {
		var _gop_err error
		_gop_ret, _gop_err = add("100", "23")
		if _gop_err != nil {
			_gop_err = errors.NewFrame(_gop_err, "add(\"100\", \"23\")", "main.gop", 14, "main.main")
			panic(_gop_err)
		}
		return
	}())
	sum, err := add("10", "abc")
	fmt.Println(sum, err)
	fmt.Println(addSafe("10", "abc"))
}
`)
}

func TestSpx(t *testing.T) {
	gopClTestEx(t, "main.tspx", `println "hi"`, `package main

import (
	fmt "fmt"
	spx "github.com/goplus/gop/cl/internal/spx"
)

type MyGame struct {
	spx.MyGame
}

func (this *MyGame) MainEntry() {
	fmt.Println("hi")
}
func main() {
	spx.Gopt_MyGame_Main(new(MyGame))
}
`)
	gopClTestEx(t, "Cat.tspx", `println "hi"`, `package main

import (
	fmt "fmt"
	spx "github.com/goplus/gop/cl/internal/spx"
)

type Cat struct {
	spx.Sprite
}

func (this *Cat) Main() {
	fmt.Println("hi")
}
func main() {
}
`)
}

func testFromDir(t *testing.T, relDir string) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed:", err)
	}
	dir = path.Join(dir, relDir)
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal("ReadDir failed:", err)
	}
	for _, fi := range fis {
		name := fi.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			testFrom(t, name, dir+"/"+name)
		})
	}
}

func testFrom(t *testing.T, name, dir string) {
	data, err := ctx.BuildDir(dir)
	if err != nil {
		t.Fatal("BuildDir failed:", err)
	}
	if chk, err := ioutil.ReadFile(filepath.Join(dir, name+".expect")); err == nil {
		if bytes.Compare(data, chk) != 0 {
			t.Fatalf("-- %v output check error --\n%v\n--\n%v", name, string(data), string(chk))
		}
	}
}

func TestFromTestdata(t *testing.T) {
	testFromDir(t, "./_testdata")
}

type localFS struct{}

func (p localFS) ReadDir(dirname string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func (p localFS) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (p localFS) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func TestFS(t *testing.T) {
	var expect = []byte(`package main

import fmt "fmt"

func main() {
	fmt.Println("Go+")
}
`)
	data, err := ctx.BuildFSDir(localFS{}, "./_testdata/hello")
	if err != nil {
		t.Fatal("build fs dir failed", err)
	}
	if bytes.Compare(data, expect) != 0 {
		t.Fatal("build fs data failed", string(data))
	}
}
