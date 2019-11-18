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

package bsf

import (
	"fmt"
	"github.com/ZR233/bsf/controller"
	"github.com/gin-gonic/gin"
	"reflect"
	"regexp"
	"strings"
)

type Engine struct {
	*gin.Engine
}

/*
自动注册controller类的所有公有方法
路由地址为 relativePath/controller名/方法名
*/
func (e *Engine) AutoRegisterController(relativePath string, controllerPtr controller.Controller) *Engine {
	autoRegisterControllerToIRoutes(e, relativePath, controllerPtr)
	return e
}

func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *Group {
	g := Group{}
	g.RouterGroup = e.Engine.Group(relativePath, handlers...)
	return &g
}

/*
自动注册controller类的所有公有方法
路由地址为 relativePath/controller名/方法名
*/
func (r *Group) AutoRegisterController(relativePath string, controllerPtr controller.Controller) *Group {
	autoRegisterControllerToIRoutes(r, relativePath, controllerPtr)
	return r
}

type Group struct {
	*gin.RouterGroup
}

func autoRegisterControllerToIRoutes(routers gin.IRoutes, relativePath string, controllerPtr controller.Controller) {
	methodMap := parseController(controllerPtr)

	for methodPath, methodFunc := range methodMap {
		methodPathArr := strings.Split(methodPath, "/")

		httpMethod, methodName := getHttpMethod(methodPathArr[1])

		finalPath := strings.Join([]string{
			relativePath,
			methodPathArr[0],
			methodName,
		}, "/")

		switch httpMethod {
		case HttpMethodAny:
			routers.Any(finalPath, methodFunc)
		case HttpMethodGet:
			routers.GET(finalPath, methodFunc)
		case HttpMethodPost:
			routers.POST(finalPath, methodFunc)
		case HttpMethodDELETE:
			routers.DELETE(finalPath, methodFunc)
		case HttpMethodPATCH:
			routers.PATCH(finalPath, methodFunc)
		case HttpMethodPUT:
			routers.PUT(finalPath, methodFunc)
		case HttpMethodOPTIONS:
			routers.OPTIONS(finalPath, methodFunc)
		case HttpMethodHEAD:
			routers.HEAD(finalPath, methodFunc)
		}
	}
}

func isControllerFunc(method reflect.Value) (r bool) {
	f := func() error {
		return nil
	}
	return method.String() == reflect.ValueOf(f).String()
}
func controllerFieldOk(valueOfControllerPtr reflect.Value) bool {
	baseController := getBaseControllerValue(valueOfControllerPtr)
	return baseController.String() == "<*controller.Base Value>"
}

func newValueOfControllerPtr(typeOfControllerPtr reflect.Type) reflect.Value {
	return reflect.New(typeOfControllerPtr.Elem())
}

func getBaseControllerValue(valueOfControllerPtr reflect.Value) reflect.Value {
	return valueOfControllerPtr.Elem().FieldByName("Base")
}

func ginHandlerFunc(context *gin.Context, typeOfControllerPtr reflect.Type, methodIter int) {

	valueOfControllerPtr := newValueOfControllerPtr(typeOfControllerPtr)
	baseController := getBaseControllerValue(valueOfControllerPtr)

	base := controller.NewBase(context)
	valueOfBase := reflect.ValueOf(base)
	baseController.Set(valueOfBase)

	in := make([]reflect.Value, 0)

	values := valueOfControllerPtr.Method(methodIter).Call(in)
	e := values[0].Interface().(error)
	if e != nil {
		_ = context.Error(e)
	}

	base.BsfHandleErrorDoNotUseThisMethod(e)
}

func parseController(controllerPtr controller.Controller) (result map[string]gin.HandlerFunc) {
	result = make(map[string]gin.HandlerFunc)

	typeOfControllerPtr := reflect.TypeOf(controllerPtr)
	valueOfControllerPtr := newValueOfControllerPtr(typeOfControllerPtr)

	name := typeOfControllerPtr.String()
	namePath := strings.Split(name, ".")
	name = namePath[1]

	if !controllerFieldOk(valueOfControllerPtr) {
		panic(fmt.Errorf("%w: [%s] did not have Base controller pointer member", ErrController, name))
	}

	num := typeOfControllerPtr.NumMethod()

	excludeMethods := map[string]bool{
		"GetContext":                       true,
		"BsfHandleErrorDoNotUseThisMethod": true,
	}

	for i := 0; i < num; i++ {
		methodName := typeOfControllerPtr.Method(i).Name
		//排除controller基类方法
		if _, ok := excludeMethods[methodName]; ok {
			continue
		}
		if !isControllerFunc(valueOfControllerPtr.Method(i)) {
			panic(fmt.Errorf("%w: class [%s] method [%s] format error", ErrController, namePath, methodName))
		}

		methodPath := name + "/" + methodName

		iter := i

		result[methodPath] = func(ctx *gin.Context) {
			ginHandlerFunc(ctx, typeOfControllerPtr, iter)
		}
	}

	return
}

type HttpMethod int

const (
	HttpMethodAny HttpMethod = iota
	HttpMethodGet
	HttpMethodPost
	HttpMethodDELETE
	HttpMethodPATCH
	HttpMethodPUT
	HttpMethodOPTIONS
	HttpMethodHEAD
)

func getHttpMethod(methodName string) (h HttpMethod, NewName string) {
	h = HttpMethodAny
	NewName = methodName
	methodNameBytes := []byte(methodName)
	reg := regexp.MustCompile(`^(HttpGet|HttpPost|HttpDelete|HttpPatch|HttpPut|HttpOptions|HttpHead)`)
	loc := reg.FindIndex(methodNameBytes)
	if loc != nil {
		method := string(methodNameBytes[:loc[1]])
		switch method {
		case "HttpGet":
			h = HttpMethodGet
		case "HttpPost":
			h = HttpMethodPost
		case "HttpDelete":
			h = HttpMethodDELETE
		case "HttpPatch":
			h = HttpMethodPATCH
		case "HttpPut":
			h = HttpMethodPUT
		case "HttpOptions":
			h = HttpMethodOPTIONS
		case "HttpHead":
			h = HttpMethodHEAD
		}
		NewName = string(methodNameBytes[loc[1]:])
	}
	return
}
