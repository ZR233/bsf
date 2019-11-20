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
	"fmt"
	"github.com/ZR233/bsf/errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

func (b *Base) PostForm(key string, must bool) (value string, err error) {
	value = b.ctx.PostForm(key)
	if must && value == "" {
		err = fmt.Errorf("%w: [%s]is empty", errors.ErrParam, key)
	}
	return
}

func (b *Base) PostFormInt(key string, must bool) (value *int, err error) {

	valueStr, err := b.PostForm(key, must)
	if err != nil {
		return
	}
	if valueStr == "" {
		return
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		err = fmt.Errorf("%w: [%s](%s)is not int", errors.ErrParam, key, valueStr)
		return
	}
	value = &valueInt
	return
}
func (b *Base) PostFormTime(key, format string, must bool) (value *time.Time, err error) {

	valueStr, err := b.PostForm(key, must)
	if err != nil {
		return
	}
	if valueStr == "" {
		return
	}

	valueTime, err := time.Parse(format, valueStr)
	if err != nil {
		err = fmt.Errorf("%w: [%s](%s)time cannot match format[%s]", errors.ErrParam, key, valueStr, format)
		return
	}
	value = &valueTime
	return
}

func (b *Base) PostFormFloat64(key string, must bool) (value *float64, err error) {

	valueStr, err := b.PostForm(key, must)
	if err != nil {
		return
	}
	if valueStr == "" {
		return
	}

	valueF, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		err = fmt.Errorf("%w: [%s](%s)is not float", errors.ErrParam, key, valueStr)
		return
	}
	value = &valueF
	return
}
