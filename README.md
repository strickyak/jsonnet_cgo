# `jsonnet_cgo`

Simple golang cgo wrapper around JSonnet VM.

Everything in libjsonnet.h is covered except the multi-file evaluators.

See jsonnet_test.go for how to use it.

## Quick example in golang:

        vm := jsonnet.Make()
        vm.ExtVar("color", "purple")

        x, err := vm.EvaluateSnippet(`Test_Demo`, `"dark " + std.extVar("color")`)

        if err != nil {
                panic(err)
        }
        if x != "\"dark purple\"\n" {
                panic("fail: we got " + x)
        }

        vm.Destroy()

## Quick examples with the command line demo program:

```
$ ( cd jsonnet_main/  ; go  build -x -a )
...
mv $WORK/b001/exe/a.out jsonnet_main
...
$ echo "{ a: 1, b: 2 }"  | jsonnet_main/jsonnet_main /dev/stdin
{
   "a": 1,
   "b": 2
}
$ cat test1.j
{
  shell: "/bin/sh",
  awk: "/usr/bin/awk",
}
$ jsonnet_main/jsonnet_main test1.j
{
   "awk": "/usr/bin/awk",
   "shell": "/bin/sh"
}
$ cat test2.j
local test1 = import "test1.j";

test1 {
  shell: "/bin/csh",
}
$ jsonnet_main/jsonnet_main test2.j
{
   "awk": "/usr/bin/awk",
   "shell": "/bin/csh"
}
$ echo ' std.extVar("a") + "bar" ' | jsonnet_main/jsonnet_main /dev/stdin a=foo
"foobar"
```

## LICENSES

Notice the various `LICENSE*` files.  I cannot offer legal advice,
but you might find that the Apache License is the most restrictive.

Most of this code comes from https://github.com/google/jsonnet
and is under the Apache License, Version 2.0, January 2004,
and our files that match filenames there are under that license.

Notice the `third_party/` directory in that distribution.
It has `json/` and `md5/` under their own licences, and our files
that match filenames there are under those licenses.

Anything new added here is under an MIT license in the plain `LICENSE` file.
