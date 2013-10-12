// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package xml2

import "testing"

func TestInit(t *testing.T) {
	XmlCheckVersion(20600)
}

/*
Tests that follow and the files that they read are translations
of examples in libxml2 and carry the following copyright and
permissions notice:
==================================================================
Except where otherwise noted in the source code (e.g. the files hash.c,
list.c and the trio files, which are covered by a similar licence but
with different Copyright notices) all the files are:

 Copyright (C) 1998-2003 Daniel Veillard.  All Rights Reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is fur-
nished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FIT-
NESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
DANIEL VEILLARD BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CON-
NECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Except as contained in this notice, the name of Daniel Veillard shall not
be used in advertising or otherwise to promote the sale, use or other deal-
ings in this Software without prior written authorization from him.
==================================================================
*/

func TestParse1(t *testing.T) {
	filename := "test1.xml"
	doc := XmlReadFile(filename, "", 0)
	if doc == nil {
		t.Errorf("Failed to parse %s\n")
	}
	XmlFreeDoc(doc)
	XmlCleanupParser()
	XmlMemoryDump()
}

func TestParse2(t *testing.T) {
	filename := "test2.xml"
	ctxt := XmlNewParserCtxt()
	if ctxt == nil {
		t.Errorf("Failed to allocate parser context\n")
	}
	doc := XmlCtxtReadFile(ctxt, filename, "", XML_PARSE_DTDVALID)
	if doc == nil {
		t.Errorf("Failed to parse %s\n")
	} else {
		if ctxt.Valid == 0 {
			t.Errorf("Failed to validate %s\n", filename)
		}
		XmlFreeDoc(doc)
	}
	XmlFreeParserCtxt(ctxt)
	XmlCleanupParser()
	XmlMemoryDump()
}

func TestParse3(t *testing.T) {
	content := "<doc/>"
	doc := XmlReadMemory(content, len(content), "noname.xml", "", 0)
	if doc == nil {
		t.Errorf("Failed to parse %s\n", content)
	}
	XmlFreeDoc(doc)
	XmlCleanupParser()
	XmlMemoryDump()
}
