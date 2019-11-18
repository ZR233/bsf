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

package main
import (
	"gitee.com/ZR233/bsf"
	"gitee.com/ZR233/bsf/example/controller"
	_ "gitee.com/ZR233/bsf/example/docs"
	"gitee.com/ZR233/bsf/example/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)



// @title Example API 文档
// @version 1.0
// @description bsf 示例文档
// @termsOfService http://swagger.io/terms/
func main(){
	engine := bsf.NewDefaultEngine()
	engine.Use(middleware.HandleError())
	engine.AutoRegisterController("", &controller.Auth{})
	engine.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	engine.Run(":28080")
}
