# archive example

This example uses `ecosystem::archive` from the library root. Its runnable
example builds a small project bundle through the generic streaming TAR writer,
reads it back, repacks it as Deflate ZIP, checks the metadata and source payload,
and exercises the tar.gz convenience API.
Both the example and archive library use pure GoML without a `go.mod` or Go
adapter sources.

From the library root, run `goml run --example basic` or `goml test --example basic`.
Run `(cd ../workflows && just ecosystem-test archive)` there to validate the
library, example and independent downstream snapshot.

This example shares the library root manifest and its dependencies. From the library root, run `goml test` to build and test the library and its examples.
