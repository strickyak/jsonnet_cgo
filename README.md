# `jsonnet_cgo`

Simple golang cgo wrapper around JSonnet VM.

Everything in libjsonnet.h is covered except the multi-file evaluators.

See jsonnet_test.go for how to use it.

## Quick example:

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
