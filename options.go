// Copyright 2022 cuipeiyu (xiaocui1023@vip.qq.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helmet

import "time"

type OptionFunc func(*Option)

func WithEnable(yesno bool) OptionFunc {
	return func(h *Option) {
		h.enable = yesno
	}
}

func WithContentSecurityPolicyReportOnly(yesno bool) OptionFunc {
	return func(h *Option) {
		h.contentSecurityPolicyReportOnly = yesno
	}
}

func WithContentSecurityPolicyDirectives(directives string) OptionFunc {
	return func(h *Option) {
		h.contentSecurityPolicyDirectives = directives
	}
}

func WithReferrerPolicyDirectives(values ...ReferrerPolicyValue) OptionFunc {
	return func(h *Option) {
		h.referrerPolicyDirectives = values
	}
}

func WithCrossOriginEmbedderPolicy(value CrossOriginEmbedderPolicyValue) OptionFunc {
	return func(h *Option) {
		h.crossOriginEmbedderPolicy = value
	}
}

func WithCrossOriginOpenerPolicy(value CrossOriginOpenerPolicyValue) OptionFunc {
	return func(h *Option) {
		h.crossOriginOpenerPolicy = value
	}
}

func WithCrossOriginResourcePolicy(value CrossOriginResourcePolicyValue) OptionFunc {
	return func(h *Option) {
		h.crossOriginResourcePolicy = value
	}
}

func WithExpectCTMaxAge(maxage time.Duration) OptionFunc {
	return func(h *Option) {
		h.expectCTMaxAge = maxage
	}
}

func WithExpectCTEnforce(yesno bool) OptionFunc {
	return func(h *Option) {
		h.expectCTEnforce = yesno
	}
}

func WithExpectCTReportUri(uri string) OptionFunc {
	return func(h *Option) {
		h.expectCTReportUri = uri
	}
}

func WithStrictTransportSecurityMaxAge(maxage time.Duration) OptionFunc {
	return func(h *Option) {
		h.strictTransportSecurityMaxAge = maxage
	}
}

func WithStrictTransportSecurityIncludeSubDomains(yesno bool) OptionFunc {
	return func(h *Option) {
		h.strictTransportSecurityIncludeSubDomains = yesno
	}
}

func WithStrictTransportSecurityPreload(yesno bool) OptionFunc {
	return func(h *Option) {
		h.strictTransportSecurityPreload = yesno
	}
}

func WithXDNSPrefetchControl(value XDNSPrefetchControlValue) OptionFunc {
	return func(h *Option) {
		h.xDNSPrefetchControl = value
	}
}

func WithXFrameOptions(value XFrameOptionsValue) OptionFunc {
	return func(h *Option) {
		h.xFrameOptions = value
	}
}

func WithXPermittedCrossDomainPolicies(value XPermittedCrossDomainPoliciesValue) OptionFunc {
	return func(h *Option) {
		h.xPermittedCrossDomainPolicies = value
	}
}

func WithHideXPoweredBy(yesno bool) OptionFunc {
	return func(h *Option) {
		h.hideXPoweredBy = yesno
	}
}

func WithReplacementXPoweredBy(replacement string) OptionFunc {
	return func(h *Option) {
		h.replacementXPoweredBy = replacement
	}
}

type Option struct {
	enable bool

	// Content-Security-Policy
	contentSecurityPolicyReportOnly bool
	contentSecurityPolicyDirectives string

	// Cross-Origin-Embedder-Policy
	crossOriginEmbedderPolicy CrossOriginEmbedderPolicyValue

	// Cross-Origin-Opener-Policy
	crossOriginOpenerPolicy CrossOriginOpenerPolicyValue

	// Cross-Origin-Resource-Policy
	crossOriginResourcePolicy CrossOriginResourcePolicyValue

	// Expect-CT
	expectCTMaxAge    time.Duration
	expectCTEnforce   bool
	expectCTReportUri string

	// Referrer-Policy
	referrerPolicyDirectives []ReferrerPolicyValue

	// Strict-Transport-Security
	strictTransportSecurityMaxAge            time.Duration
	strictTransportSecurityIncludeSubDomains bool
	strictTransportSecurityPreload           bool

	// X-DNS-Prefetch-Control
	xDNSPrefetchControl XDNSPrefetchControlValue

	// X-Frame-Options
	xFrameOptions XFrameOptionsValue

	// X-Permitted-Cross-Domain-Policies
	xPermittedCrossDomainPolicies XPermittedCrossDomainPoliciesValue

	// X-Powered-By
	hideXPoweredBy        bool
	replacementXPoweredBy string
}
