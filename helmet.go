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

import "net/http"

type Helmet struct {
	contentSecurityPolicy         *ContentSecurityPolicy
	crossOriginEmbedderPolicy     *CrossOriginEmbedderPolicy
	crossOriginOpenerPolicy       *CrossOriginOpenerPolicy
	crossOriginResourcePolicy     *CrossOriginResourcePolicy
	originAgentCluster            *OriginAgentCluster
	referrerPolicy                *ReferrerPolicy
	strictTransportSecurity       *StrictTransportSecurity
	xContentTypeOptions           *XContentTypeOptions
	xDNSPrefetchControl           *XDNSPrefetchControl
	xDownloadOptions              *XDownloadOptions
	xFrameOptions                 *XFrameOptions
	xPermittedCrossDomainPolicies *XPermittedCrossDomainPolicies
	xPoweredBy                    *XPoweredBy
	xXSSProtection                *XXSSProtection
}

func NewHelmet(funs ...OptionFunc) *Helmet {
	h := &Helmet{
		contentSecurityPolicy:         NewContentSecurityPolicy(funs...),
		crossOriginEmbedderPolicy:     NewCrossOriginEmbedderPolicy(funs...),
		crossOriginOpenerPolicy:       NewCrossOriginOpenerPolicy(funs...),
		crossOriginResourcePolicy:     NewCrossOriginResourcePolicy(funs...),
		originAgentCluster:            NewOriginAgentCluster(funs...),
		referrerPolicy:                NewReferrerPolicy(funs...),
		strictTransportSecurity:       NewStrictTransportSecurity(funs...),
		xContentTypeOptions:           NewXContentTypeOptions(funs...),
		xDNSPrefetchControl:           NewXDNSPrefetchControl(funs...),
		xDownloadOptions:              NewXDownloadOptions(funs...),
		xFrameOptions:                 NewXFrameOptions(funs...),
		xPermittedCrossDomainPolicies: NewXPermittedCrossDomainPolicies(funs...),
		xPoweredBy:                    NewXPoweredBy(funs...),
		xXSSProtection:                NewXXSSProtection(funs...),
	}

	return h
}

var _ http.Handler = (*Helmet)(nil)

func (h *Helmet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.contentSecurityPolicy.ServeHTTP(rw, r)
	h.crossOriginEmbedderPolicy.ServeHTTP(rw, r)
	h.crossOriginOpenerPolicy.ServeHTTP(rw, r)
	h.crossOriginResourcePolicy.ServeHTTP(rw, r)
	h.originAgentCluster.ServeHTTP(rw, r)
	h.referrerPolicy.ServeHTTP(rw, r)
	h.strictTransportSecurity.ServeHTTP(rw, r)
	h.xContentTypeOptions.ServeHTTP(rw, r)
	h.xDNSPrefetchControl.ServeHTTP(rw, r)
	h.xDownloadOptions.ServeHTTP(rw, r)
	h.xFrameOptions.ServeHTTP(rw, r)
	h.xPermittedCrossDomainPolicies.ServeHTTP(rw, r)
	h.xPoweredBy.ServeHTTP(rw, r)
	h.xXSSProtection.ServeHTTP(rw, r)
}
