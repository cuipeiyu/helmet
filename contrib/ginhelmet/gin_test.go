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

package ginhelmet

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/cuipeiyu/helmet"
)

func TestMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(
		// set test data
		func(c *gin.Context) {
			c.Header(helmet.HeaderXPoweredBy, "go-helmet")
			c.Header(helmet.HeaderXPermittedCrossDomainPolicies, "")

			c.Next()
		},
		// override the test data
		Middleware(),
	)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "done")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Header().Get(helmet.HeaderContentSecurityPolicy), helmet.DefaultContentSecurityPolicyDirectives)
	assert.Equal(t, w.Header().Get(helmet.HeaderReferrerPolicy), "")
	assert.Contains(t, w.Header().Get(helmet.HeaderStrictTransportSecurity), "max-age")
	assert.Equal(t, w.Header().Get(helmet.HeaderXContentTypeOptions), string(helmet.XContentTypeOptionsNosniff))
	assert.Equal(t, w.Header().Get(helmet.HeaderXDNSPrefetchControl), string(helmet.XDNSPrefetchControlOff))
	assert.Equal(t, w.Header().Get(helmet.HeaderXDownloadOptions), string(helmet.XDownloadOptionsNoopen))
	assert.Equal(t, w.Header().Get(helmet.HeaderXFrameOptions), string(helmet.XFrameOptionsSameOrigin))
	assert.Equal(t, w.Header().Get(helmet.HeaderXPermittedCrossDomainPolicies), string(helmet.XPermittedCrossDomainPoliciesNone))
	assert.Equal(t, w.Header().Get(helmet.HeaderXPoweredBy), "")
}
