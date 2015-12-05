// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at https://golang.org/project

// Changes Copyright 2015 The Myth Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner_test

import (
	"fmt"
	"go/token"
	"myth/scanner"
)

func ExampleScanner_Scan() {
	// src is the input that we want to tokenize.
	src := []byte("cos(x) + 1*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	// output:
	// 1:1	IDENT	"cos"
	// 1:4	(	""
	// 1:5	IDENT	"x"
	// 1:6	)	""
	// 1:8	+	""
	// 1:10	INT	"1"
	// 1:11	*	""
	// 1:12	IDENT	"sin"
	// 1:15	(	""
	// 1:16	IDENT	"x"
	// 1:17	)	""
	// 1:19	;	"\n"
	// 1:19	COMMENT	"// Euler"
}
