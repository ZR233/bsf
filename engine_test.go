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
	con "gitee.com/ZR233/bsf/controller"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

type testController struct {
	*con.Base
}

func (t testController) PostMethod() error {
	return nil
}

type testControllerFieldError struct {
	*gin.Engine
}

type testControllerMethodError struct {
	*con.Base
}

func (testControllerMethodError) PostMethod2() {
	return
}

func Test_parseController(t *testing.T) {
	r := parseController(&testController{})
	println(len(r))
}

func Test_ginHandlerFunc(t *testing.T) {
	typeOfController := reflect.TypeOf(&testController{})
	typeOfControllerMethodError := reflect.TypeOf(&testControllerMethodError{})
	typeOfControllerFieldError := reflect.TypeOf(&testControllerFieldError{})
	type args struct {
		context          *gin.Context
		typeOfController reflect.Type
		methodIter       int
	}
	tests := []struct {
		name string
		args args
	}{
		{"成员错误", args{&gin.Context{}, typeOfControllerFieldError, 2}},
		{"normal", args{&gin.Context{}, typeOfController, 2}},
		{"方法格式错误", args{&gin.Context{}, typeOfControllerMethodError, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ginHandlerFunc(tt.args.context, tt.args.typeOfController, tt.args.methodIter)
		})
	}
}

func Test_isControllerFunc(t *testing.T) {
	methodNormal := reflect.New(reflect.TypeOf(&testController{}).Elem())
	methodError := reflect.New(reflect.TypeOf(&testControllerMethodError{}).Elem())

	type args struct {
		method reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		wantR bool
	}{
		{"格式正确", args{methodNormal.Method(2)}, true},
		{"格式错误", args{methodError.Method(2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := isControllerFunc(tt.args.method); gotR != tt.wantR {
				t.Errorf("isControllerFunc() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_controllerFieldOk(t *testing.T) {
	controllerNormal := reflect.New(reflect.TypeOf(&testController{}).Elem())
	controllerError := reflect.New(reflect.TypeOf(&testControllerFieldError{}).Elem())

	type args struct {
		valueOfController reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"格式正确", args{controllerNormal}, true},
		{"格式错误", args{controllerError}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllerFieldOk(tt.args.valueOfController); got != tt.want {
				t.Errorf("controllerFieldOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHttpMethod(t *testing.T) {
	type args struct {
		methodName string
	}
	tests := []struct {
		name        string
		args        args
		wantH       HttpMethod
		wantNewName string
	}{
		{"1", args{"HttpPostTest"}, HttpMethodPost, "Test"},
		{"2", args{"GetTest2"}, HttpMethodAny, "GetTest2"},
		{"3", args{"Test3HttpGet"}, HttpMethodAny, "Test3HttpGet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotNewName := getHttpMethod(tt.args.methodName)
			if gotH != tt.wantH {
				t.Errorf("getHttpMethod() gotH = %v, want %v", gotH, tt.wantH)
			}
			if gotNewName != tt.wantNewName {
				t.Errorf("getHttpMethod() gotNewName = %v, want %v", gotNewName, tt.wantNewName)
			}
		})
	}
}
