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

package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetContext() *gin.Context
}

type Base struct {
	ctx *gin.Context
}

func NewBase(ctx *gin.Context) *Base {
	b := &Base{}
	b.ctx = ctx
	return b
}

func (b Base) GetContext() *gin.Context {
	return b.ctx
}

func (b Base) BsfHandleErrorDoNotUseThisMethod(err error) {
	if err != nil {
		b.ctx.Set("error", err)
	}
}

type Handler func() error
