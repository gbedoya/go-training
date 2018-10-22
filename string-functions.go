package main

import (
	st "strings"
)
import "fmt"

var p = fmt.Println

func StringFunctions() {

	p("Contains:", st.Contains("test", "es"))
	p("Count:", st.Count("test", "t"))
	p("HasPrefix:", st.HasPrefix("test", "te"))
	p("HasSuffix:", st.HasSuffix("test", "st"))
	p("Index:", st.Index("test", "e"))
	p()
	p("Join:", st.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat", st.Repeat("a", 5))
	p("Replace", st.Replace("foo", "o", "0", -1))
	p("Replace", st.Replace("foo", "o", "0", 1))
	p("ToLower:", st.ToLower("TEST"))
	p("ToUpper:", st.ToUpper("test"))
	p()
	p("Len:", len("hello"))
	p("Char:", "Hello"[1])
	p("String:", string("Hello"[1]))


}
