// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mtls

import (
	"testing"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/echo"
	"istio.io/istio/pkg/test/framework/components/environment"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/tests/integration/security/util/reachability"
)

// This test verifies reachability under different authN scenario when automtls enabled
// - app A to app B using mTLS.
// In each test, the steps are:
// - Configure authn policy.
// - Wait for config propagation.
// - Send HTTP/gRPC requests between apps.
func TestMtlsStrict(t *testing.T) {
	framework.NewTest(t).
		Run(func(ctx framework.TestContext) {

			rctx := reachability.CreateContext(ctx, g, p)
			systemNM := namespace.ClaimSystemNamespaceOrFail(ctx, ctx)

			testCases := []reachability.TestCase{
				{
					ConfigFile:          "global-mtls-on-no-dr.yaml",
					Namespace:           systemNM,
					RequiredEnvironment: environment.Kube,
					Include: func(src echo.Instance, opts echo.CallOptions) bool {
						// Exclude calls to the headless TCP port.
						if opts.Target == rctx.Headless && opts.PortName == "tcp" {
							return false
						}

						return true
					},
					ExpectSuccess: func(src echo.Instance, opts echo.CallOptions) bool {
						// When mTLS is in STRICT mode, DR's TLS settings are default to mTLS so the result would
						// be the same as having global DR rule.
						if opts.Target == rctx.Naked {
							// calls to naked should always succeed.
							return true
						}

						// If source is naked, and destination is not, expect failure.
						return !(src == rctx.Naked && opts.Target != rctx.Naked)
					},
				},
				{
					ConfigFile:          "global-plaintext.yaml",
					Namespace:           systemNM,
					RequiredEnvironment: environment.Kube,
					Include: func(src echo.Instance, opts echo.CallOptions) bool {
						// Exclude calls to the headless TCP port.
						if opts.Target == rctx.Headless && opts.PortName == "tcp" {
							return false
						}

						return true
					},
					ExpectSuccess: func(src echo.Instance, opts echo.CallOptions) bool {
						// When mTLS is disabled, all traffic should work.
						return true
					},
				},
			}
			rctx.Run(testCases)
		})
}
