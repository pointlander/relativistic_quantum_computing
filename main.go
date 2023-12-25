// Copyright 2023 The Relativistic Quantum Computing Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/itsubaki/q"
)

func main() {
	qsim := q.New()

	q0 := qsim.Zero()
	q1 := qsim.Zero()
	q2 := qsim.Zero()

	qsim.H(q0).H(q1).CNOT(q0, q1).CNOT(q1, q2)

	for _, s := range qsim.State() {
		fmt.Println(s)
	}

	m0 := qsim.Measure(q0)
	m1 := qsim.Measure(q1)
	qsim.Measure(q2)
	fmt.Println(m0.IsZero() == m1.IsZero())

	for _, s := range qsim.State() {
		fmt.Println(s)
	}
}
