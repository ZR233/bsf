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
	"github.com/gin-gonic/gin"
)

type ErrorHandler func(err error, ctx *gin.Context)

func HandleError(handler ErrorHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if len(context.Errors) > 0 {
				err := context.Errors[0].Err
				if err != nil {
					handler(err, context)
				}
			}
		}()

		context.Next()
	}
}
