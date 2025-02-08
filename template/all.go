package template

import _ "embed"

//go:embed bean.tpl
var Bean string

//go:embed iservice.tpl
var IService string

//go:embed service.tpl
var Service string
