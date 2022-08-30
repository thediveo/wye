// Copyright 2021 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wye

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gleak"
)

type mykey string

const (
	canarykey = mykey("canary")
	canaryval = "cheep"
)

var _ = Describe("mixing podman connection with other contexts", func() {

	var longctx context.Context

	BeforeEach(func() {
		var cancel context.CancelFunc
		longctx, cancel = context.WithCancel(context.Background())
		longctx = context.WithValue(longctx, canarykey, canaryval)
		DeferCleanup(func() {
			// Grrrr ... I should really pay attention to my own advice: don't
			// pass the result of a Goroutines() call to Eventually(...), always
			// pass Goroutines itself.
			Eventually(Goroutines).WithTimeout(1 * time.Second).ShouldNot(HaveLeaked())
			cancel()
		})
	})

	It("doesn't cancel the original context if mix-in isn't cancellable", func() {
		longerctx, cancel := Mixin(longctx, context.Background())
		cancel()
		Expect(longctx.Err()).To(BeNil())
		Expect(longerctx.Err()).To(BeNil())
		Expect(longerctx.Value(canarykey)).To(Equal(canaryval))
	})

	When("mix-in context gets cancelled", func() {

		It("cancels the mixed context if mix-in is cancellable", func() {
			shortctx, shortcancel := context.WithCancel(context.Background())
			defer shortcancel()
			longerctx, cancel := Mixin(longctx, shortctx)
			cancel()
			Expect(longctx.Err()).To(BeNil())
			Expect(longerctx.Err()).NotTo(BeNil())
		})

		It("sets the deadline for the mixed context if mix-in has deadline", func() {
			shortctx, shortcancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer shortcancel()
			shortctx, shortcancel = context.WithCancel(shortctx)
			defer shortcancel()
			longerctx, cancel := Mixin(longctx, shortctx)
			cancel()
			Expect(longerctx.Err()).NotTo(BeNil())
		})

		It("cancels the mixed context when mix-in meets its deadline", func() {
			shortctx, shortcancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer shortcancel()
			shortctx, shortcancel = context.WithCancel(shortctx)
			defer shortcancel()
			longerctx, cancel := Mixin(longctx, shortctx)
			defer cancel()
			Eventually(longerctx.Err).WithTimeout(1 * time.Second).ShouldNot(BeNil())
		})

	})

	When("long-lived context gets cancelled", func() {

		It("cancels only the mixed context", func() {
			longctx, longcancel := context.WithCancel(longctx)
			longerctx, cancel := Mixin(longctx, context.Background())
			defer cancel()
			longcancel()
			Expect(longerctx.Err()).NotTo(BeNil())
		})

		It("cancels only the mixed context at deadline", func() {
			longctx, longcancel := context.WithTimeout(longctx, 500*time.Millisecond)
			defer longcancel()
			longerctx, cancel := Mixin(longctx, context.Background())
			defer cancel()
			Eventually(longerctx.Err).WithTimeout(1 * time.Second).ShouldNot(BeNil())
		})

	})

	Context("nil contexts", func() {

		It("cannot mix a context into a nil context", func() {
			Expect(func() { _, _ = Mixin(nil, context.Background()) }).To(
				PanicWith(MatchRegexp("cannot mix into nil")))
		})

		It("cannot mix a nil context in", func() {
			Expect(func() { _, _ = Mixin(context.Background(), nil) }).To(
				PanicWith(MatchRegexp("cannot mix-in nil context")))
		})

	})

})
