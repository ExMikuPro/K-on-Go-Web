// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/about": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统"
                ],
                "summary": "获取框架版本信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/admin/add/article": {
            "post": {
                "tags": [
                    "文章"
                ],
                "summary": "添加文章",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.articleTable"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/admin/add/category": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "添加分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "缩略名",
                        "name": "slug",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序",
                        "name": "order",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/admin/add/group": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "添加用户组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序",
                        "name": "order",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/admin/add/tag": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "添加标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "缩略名",
                        "name": "slug",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序",
                        "name": "order",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/admin/add/user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "mail",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "网站地址",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户昵称",
                        "name": "screenName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户组",
                        "name": "group",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/content/article/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "获取文章内容",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/content/category/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "获取分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分类ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/content/tag/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "获取标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/list/article": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "获取列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/list/category": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "获取分类列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/list/tag": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "获取标签列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/update/user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "mail",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "网址",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "screenName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户组",
                        "name": "group",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        },
        "/userLogin": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GeneralJSONHeader"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.GeneralJSONHeader": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "msg": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "service.articleTable": {
            "type": "object",
            "properties": {
                "allowComment": {
                    "description": "是否允许评论",
                    "type": "boolean"
                },
                "authorId": {
                    "description": "内容所属用户",
                    "type": "string"
                },
                "created": {
                    "description": "内容生成时的时间戳",
                    "type": "integer"
                },
                "id": {
                    "description": "主键，自增",
                    "type": "string"
                },
                "modified": {
                    "description": "内容修改时的时间戳",
                    "type": "integer"
                },
                "order": {
                    "description": "排序",
                    "type": "integer"
                },
                "password": {
                    "description": "受保护内容的密码",
                    "type": "string"
                },
                "slug": {
                    "description": "内容缩略名",
                    "type": "string"
                },
                "status": {
                    "description": "内容状态",
                    "type": "string"
                },
                "template": {
                    "description": "内容所使用的模版",
                    "type": "string"
                },
                "text": {
                    "description": "内容文字",
                    "type": "string"
                },
                "title": {
                    "description": "内容标题",
                    "type": "string"
                },
                "type": {
                    "description": "内容类别",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1206",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "AzuLLia",
	Description:      "一只轻量级无头CMS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
