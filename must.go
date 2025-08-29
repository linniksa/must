// Package must provides tiny helpers to unwrap results from functions that
// return (..., error). If err != nil, they panic(err).
package must

// P1 unwraps (v1, error) -> v1
func P1[P1 any](p1 P1, err error) P1 {
	if err != nil {
		panic(err)
	}
	return p1
}

// P2 unwraps (v1, v2, error) -> (v1, v2)
func P2[P1, P2 any](p1 P1, p2 P2, err error) (P1, P2) {
	if err != nil {
		panic(err)
	}
	return p1, p2
}

// P3 unwraps (v1, v2, v3, error) -> (v1, v2, v3)
func P3[P1, P2, P3 any](p1 P1, p2 P2, p3 P3, err error) (P1, P2, P3) {
	if err != nil {
		panic(err)
	}
	return p1, p2, p3
}

// P4 unwraps (v1, v2, v3, v4, error) -> (v1, v2, v3, v4)
func P4[P1, P2, P3, P4 any](p1 P1, p2 P2, p3 P3, p4 P4, err error) (P1, P2, P3, P4) {
	if err != nil {
		panic(err)
	}
	return p1, p2, p3, p4
}
