/*
Package wye mixes a (comparably) “short-lived” [context.Context] into another
long-living context: kind of a “Y” joint. The opposite of Go's “⅄” pattern of
deriving new contexts from existing contexts.

If it weren't for the bizarre “design” of the [Podman REST API bindings], this
context Mixin “abomination” would never exist.

# Usage

First, set up a long-living Podman connection that later will get used over and
over again by service handlers. What we get back here is actually a context;
that's why we call a function named [bindings.NewConnection]...!

	conn, _ := bindings.NewConnection(context.Background())

Later, in some HTTP service handler:

	func handler(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := wye.Mixin(conn, r.Context())
		defer cancel() // ...won't touch the original contexts, but ensures proper ctx cleanup
		cntrs, _ := containers.List(ctx, nil, nil, nil, nil, nil, nil)
	}

[Podman REST API bindings]: https://pkg.go.dev/github.com/containers/podman/v2/pkg/bindings
[bindings.NewConnection]: https://github.com/containers/podman/v2/pkg/bindings.NewConnection
*/
package wye
