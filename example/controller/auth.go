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
	"errors"
	"github.com/ZR233/bsf/controller"
)

type Auth struct {
	*controller.Base
}

// Hello godoc
// @Summary 测试
// @tags Test
// @Description 测试
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200 {string} success
// @Router /Auth/Hello [post]
func (a *Auth) Hello() (err error) {

	err = errors.New("错误了")
	return
}

func (a Auth) HttpGetHello2() (err error) {

	return
}
