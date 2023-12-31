// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "联系人",
            "url": "http://www.swagger.io/support",
            "email": "584807419@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/re": {
            "get": {
                "description": "关注或者取消关注用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "用户关系操作",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "关注者",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "被关注者",
                        "name": "to_user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.BaseResp"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.BaseResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.BaseResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BaseResp": {
            "type": "object",
            "properties": {
                "status_code": {
                    "description": "业务响应码",
                    "type": "integer",
                    "example": 1
                },
                "status_msg": {
                    "description": "业务消息",
                    "type": "string",
                    "example": "xxxx"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0 版本",
	Host:             "videotools.cn",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "mock tiktok",
	Description:      "字节青训营-模仿抖音项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
