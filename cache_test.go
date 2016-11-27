package cache

import (
	"testing"
)

func Test_TagCache(t *testing.T) {

	c, err := New(Options{Adapter: "memory"})
	if err != nil {
		t.Fatal(err)
	}

	// base use
	err = c.Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}

	res := ""
	c.Get("da", &res)

	if res != "weisd" {
		t.Fatal("base put faield")
	}

	t.Log("ok")

	// use tags/namespace
	err = c.Tags([]string{"insion"}).Set("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}
	res = ""
	c.Tags([]string{"insion"}).Get("da", &res)

	if res != "weisd" {
		t.Fatal("tags put faield")
	}

	t.Log("ok")

	err = c.Tags([]string{"macross"}).Set("macross", "macross_contrib", 300)
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	c.Tags([]string{"macross"}).Get("macross", &res)

	if res != "macross_contrib" {
		t.Fatal("not macross_contrib")
	}

	t.Log("ok")

	// flush namespace
	err = c.Tags([]string{"macross"}).Flush()
	if err != nil {
		t.Fatal(err)
	}

	res = ""
	c.Tags([]string{"macross"}).Get("macross", &res)
	if res != "" {
		t.Fatal("flush faield")
	}

	res = ""
	c.Tags([]string{"macross"}).Get("bb", &res)
	if res != "" {
		t.Fatal("flush faield")
	}

	// still store in
	res = ""
	c.Tags([]string{"insion"}).Get("da", &res)
	if res != "weisd" {
		t.Fatal("where")
	}

	t.Log("ok")

}
