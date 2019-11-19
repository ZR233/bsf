/*
 * Copyright (c) [2019] [zr]
 *    [bsf] is licensed under the Mulan PSL v1.
 *    You can use this software according to the terms and conditions of the Mulan PSL v1.
 *    You may obtain a copy of Mulan PSL v1 at:
 *       http://license.coscl.org.cn/MulanPSL
 *    THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 *    IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 *    PURPOSE.
 *
 *    See the Mulan PSL v1 for more details.
 */

package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginTime := time.Now()

		defer func() {
			code := 0
			memo := ""
			logMsg := ""

			userid := 0
			src := ""

			endTime := time.Now()

			postForm := c.Request.PostForm
			params, _ := json.Marshal(postForm)

			url := c.Request.URL.Path
			event := log.WithFields(logrus.Fields{
				"execTime": endTime.Sub(beginTime) / time.Millisecond,
				"trace":    url,
				//"time":      beginTime,
				"optUserId": userid,
				"src":       src,
				"code":      code,
				"params":    string(params),
			})
			event.Time = beginTime
			if code != 0 {
				event.Warn(logMsg)
			} else {
				event.Info(logMsg)
			}

		}()

		c.Next()
	}
}
