// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package sqlsmith

import (
	"math/rand"

	"github.com/cockroachdb/cockroach/pkg/util/syncutil"
)

// StatementWeight is the generic weight type.
type StatementWeight struct {
	weight int
	elem   statement
}

// NewStatementWeightedSampler creates a StatementSampler that produces
// Statements. They are returned at the relative frequency of the values of
// weights. All weights must be >= 1.
func NewWeightedStatementSampler(weights []StatementWeight, seed int64) *StatementSampler {
	sum := 0
	for _, w := range weights {
		if w.weight < 1 {
			panic("expected weight >= 1")
		}
		sum += w.weight
	}
	if sum == 0 {
		panic("expected weights")
	}
	samples := make([]statement, sum)
	pos := 0
	for _, w := range weights {
		for count := 0; count < w.weight; count++ {
			samples[pos] = w.elem
			pos++
		}
	}
	return &StatementSampler{
		rnd:     rand.New(rand.NewSource(seed)),
		samples: samples,
	}
}

// StatementSampler is a weighted statement sampler.
type StatementSampler struct {
	mu      syncutil.Mutex
	rnd     *rand.Rand
	samples []statement
}

// Next returns the next weighted sample.
func (w *StatementSampler) Next() statement {
	w.mu.Lock()
	v := w.samples[w.rnd.Intn(len(w.samples))]
	w.mu.Unlock()
	return v
}
