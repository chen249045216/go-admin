// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"html/template"
)

// Attribute is the component interface of template. Every component of
// template should implement it.
type Attribute struct {
	TemplateList map[string]string
}

// Page used in the template as a top variable.
type Page struct {
	// User is the login user.
	User models.UserModel

	// Menu is the left side menu of the template.
	Menu menu.Menu

	// Panel is the main content of template.
	Panel Panel

	// System contains some system info.
	System SystemInfo

	// UrlPrefix is the prefix of url.
	UrlPrefix string

	// Title is the title of the web page.
	Title string

	// Logo is the logo of the template.
	Logo template.HTML

	// MiniLogo is the downsizing logo of the template.
	MiniLogo template.HTML

	// ColorScheme is the color scheme of the template.
	ColorScheme string
}

// SystemInfo contains basic info of system.
type SystemInfo struct {
	Version string
}

// Panel contains the main content of the template which used as pjax.
type Panel struct {
	Content     template.HTML
	Title       string
	Description string
	Url         string
}

type GetPanel func() Panel

// Form is the form field with different options.
type Form struct {
	Field    string
	TypeName string
	Head     string
	Default  string
	Editable bool
	FormType string
	Value    string
	Options  []map[string]string
	FilterFn FieldFilterFn
	PostFun  FieldFilterFn
}

// RowModel contains ID and value of the single query result.
type RowModel struct {
	ID    int64
	Value string
	Row   map[string]interface{}
}

// FieldFilterFn is filter function of data.
type FieldFilterFn func(value RowModel) interface{}

// Field is the table field.
type Field struct {
	FilterFn  FieldFilterFn
	Field     string
	TypeName  string
	Head      string
	JoinTable []Join
	Sortable  bool
	Filter    bool
	Hide      bool
}

type Join struct {
	Table      string
	Field      string
	TableField string
	HasChild   bool
	JoinTable  *Join
}

// InfoPanel
type InfoPanel struct {
	FieldList   []Field
	Table       string
	Title       string
	Description string
	HeaderHtml  template.HTML
	FooterHtml  template.HTML
}

// FormPanel
type FormPanel struct {
	FormList    []Form
	Table       string
	Title       string
	Description string
	HeaderHtml  template.HTML
	FooterHtml  template.HTML
}
