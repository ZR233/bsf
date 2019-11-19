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
	"fmt"
	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				if err, ok := p.(error); ok {
					_ = context.Error(err)
				} else {
					_ = context.Error(fmt.Errorf("%s", p))

				}
			}
		}()
		context.Next()
	}
}
