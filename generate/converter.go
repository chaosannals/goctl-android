// MIT License
//
// Copyright (c) 2020 goctl-android
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

package generate

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

type (
	Plugin struct {
		Api           *spec.ApiSpec `json:"-"`
		ApiFilePath   string
		Style         string
		Dir           string
		ParentPackage string
	}
	Spec struct {
		Beans   []*Bean
		Service IService
	}

	Tag struct {
		tp   string
		name string
	}
	Type struct {
		Name string
	}
)

func (p *Plugin) SetParentPackage(parentPackage string) {
	p.ParentPackage = parentPackage
}

func (p *Plugin) Convert() (*Spec, error) {
	var (
		ret   Spec
		beans = make(map[string]*Bean)
	)

	// 遍历数据类型
	for _, each := range p.Api.Types {
		list, err := getBean(p.ParentPackage, each)
		if err != nil {
			return nil, err
		}

		for _, bean := range list {
			beans[bean.Name.Lower()] = bean
		}
	}

	// 遍历路由
	var r = make(map[string]spec.Route)
	for _, g := range p.Api.Service.Groups {
		prefix := g.GetAnnotation("prefix")
		for _, route := range g.Routes {
			r[fmt.Sprintf("%s%s", prefix, route.Path)] = route
		}
	}

	imports, routes := getRoute(r, beans)
	ret.Service = IService{
		ParentPackage: p.ParentPackage,
		Import:        strings.Join(trimList(imports), pathx.NL),
		Routes:        routes,
	}

	for _, each := range beans {
		ret.Beans = append(ret.Beans, each)
	}

	sort.Slice(ret.Beans, func(i, j int) bool {
		return ret.Beans[i].Name.Source() < ret.Beans[j].Name.Source()
	})

	sort.Slice(ret.Service.Routes, func(i, j int) bool {
		return ret.Service.Routes[i].Path < ret.Service.Routes[j].Path
	})

	return &ret, nil
}

func trimList(list []string) []string {
	var ret []string
	for _, each := range list {
		tmp := strings.TrimSpace(each)
		if len(tmp) == 0 {
			continue
		}
		ret = append(ret, tmp)
	}
	return ret
}

func getRoute(in map[string]spec.Route, m map[string]*Bean) ([]string, []*Route) {
	var list []*Route
	var imports []string

	for urlPath, each := range in {
		handlerName := each.Handler

		doc := each.AtDoc.Properties["summary"]
		if len(doc) > 0 {
			doc = strings.ReplaceAll(doc, "'", "")
			doc = strings.ReplaceAll(doc, "`", "")
			doc = strings.ReplaceAll(doc, `"`, "")
			doc = "// " + doc
		}

		path, ids, idsExpr := parsePath(urlPath)
		var bean *Bean
		if each.RequestType != nil {
			bean = m[strings.ToLower(each.RequestType.Name())]
		}

		var queryId []string
		var bodyPrefix string
		var queryExpr, pathIdExpr, headersExpr, headerIdsExpr string
		var showRequestBody bool
		if bean != nil {
			imports = append(imports, bean.Import)
			for _, query := range bean.FormTag {
				queryId = append(queryId, fmt.Sprintf("req.get%s()", stringx.From(query).ToCamel()))
			}
			queryExpr = bean.GetQuery()
			showRequestBody = len(bean.JsonTag) > 0
			if showRequestBody {
				queryExpr = queryExpr + ", "
			}
			pathIdExpr = toRetrofitPath(ids, bean)
			if len(queryId) > 0 {
				pathIdExpr = pathIdExpr + ", "
			}
			headersExpr, headerIdsExpr = bean.GetHeaders()
			if len(headersExpr) > 0 && (len(pathIdExpr) > 0 || len(queryId) > 0 || showRequestBody) {
				headersExpr = headersExpr + ", "
				headerIdsExpr = headerIdsExpr + ","
			}
			if !showRequestBody && (len(pathIdExpr) > 0 || len(queryId) > 0 || len(headersExpr) > 0) {
				bodyPrefix = ", "
			}
		}

		requestBeanName := ""
		if each.RequestType != nil {
			requestBeanName = stringx.From(each.RequestType.Name()).Title()
		}
		responseBeanName := ""
		if each.ResponseType != nil {
			responseBeanName = stringx.From(each.ResponseType.Name()).Title()
		}

		list = append(list, &Route{
			MethodName:       stringx.From(handlerName).Untitle(),
			Method:           strings.ToUpper(each.Method),
			Path:             path,
			RequestBeanName:  requestBeanName,
			ResponseBeanName: responseBeanName,
			HasRequest:       each.RequestType != nil,
			ShowRequestBody:  showRequestBody,
			HasResponse:      each.ResponseType != nil,
			HavePath:         len(ids) > 0,
			PathId:           strings.Join(idsExpr, ","),
			PathIdExpr:       pathIdExpr,
			QueryId:          strings.Join(queryId, ","),
			HaveQuery:        len(queryId) > 0,
			QueryIdExpr:      queryExpr,
			HaveHeaders:      len(headersExpr) > 0,
			HeaderIdsExpr:    headerIdsExpr,
			HeadersExpr:      headersExpr,
			BodyPrefix:       bodyPrefix,
			Doc:              doc,
		})
	}

	sort.Strings(imports) // 固定排序
	return imports, list
}

func parsePath(path string) (string, []string, []string) {
	p := strings.Split(path, "/")
	var list, ids, idsExpr []string
	for _, each := range p {
		if strings.Contains(each, ":") {
			id := strings.ReplaceAll(each, ":", "")
			list = append(list, "{"+id+"}")
			ids = append(ids, id)
			idsExpr = append(idsExpr, "req.get"+stringx.From(id).ToCamel()+"()")
			continue
		}

		list = append(list, each)
	}
	return strings.Join(list, "/"), ids, idsExpr
}

func toRetrofitPath(ids []string, bean *Bean) string {
	if bean == nil {
		return ""
	}
	var list []string
	for _, each := range ids {
		m := bean.GetMember(each)
		if m == nil {
			continue
		}

		list = append(list, fmt.Sprintf(`@Path("%s") %s %s`, each, m.TypeName, each))
	}
	return strings.Join(list, ", ")
}

func getBean(parentPackage string, tp spec.Type) ([]*Bean, error) {
	var bean Bean
	var list []*Bean
	bean.Name = stringx.From(tp.Name())
	bean.ParentPackage = parentPackage

	definedType, ok := tp.(spec.DefineStruct)
	if !ok {
		return nil, fmt.Errorf("type %s not supported", tp.Name())
	}

	for _, m := range definedType.Members {
		externalBeans, err := getBeans(parentPackage, m, &bean)
		if err != nil {
			return nil, err
		}

		list = append(list, externalBeans...)
	}
	return list, nil
}

func getBeans(parentPackage string, member spec.Member, bean *Bean) ([]*Bean, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, member.Type)
	if err != nil {
		return nil, err
	}

	tag := NewTag(member.Tag)
	name := tag.GetTag()
	if tag.IsJson() {
		bean.JsonTag = append(bean.JsonTag, name)
	}
	if tag.IsPath() {
		bean.PathTag = append(bean.PathTag, name)
	}
	if tag.IsForm() {
		bean.FormTag = append(bean.FormTag, name)
	}
	if tag.IsHeader() {
		bean.HeaderTag = append(bean.HeaderTag, name)
	}

	bean.Import = strings.Join(imports, pathx.NL)
	comment := strings.Join(member.Type.Comments(), " ")
	doc := strings.Join(member.Docs, pathx.NL)
	if len(comment) > 0 {
		comment = "// " + comment
	}
	if len(doc) > 0 {
		doc = "// " + doc
	}

	bean.Members = append(bean.Members, &Member{
		Name:     stringx.From(member.Name),
		Field:    name,
		TypeName: typeName,
		Comment:  comment,
		Doc:      doc,
		Tag:      tag.tp,
	})
	beans = append(beans, bean)
	return beans, nil
}

func getTypeName(parentPackage string, expr interface{}) ([]*Bean, []string, string, error) {
	set := collection.NewSet()
	switch v := expr.(type) {
	case spec.PrimitiveType:
		imp, typeName := toJavaType(parentPackage, v.Name())
		set.AddStr(imp)
		return nil, set.KeysStr(), typeName, nil
	case spec.PointerType:
		return getTypeName(parentPackage, v.Type)
	case spec.MapType:
		set.AddStr("import java.util.HashMap;")
		beans, imports, typeName, err := toJavaMap(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		set.AddStr(imports...)
		return beans, set.KeysStr(), typeName, nil
	case spec.ArrayType:
		set.AddStr("import java.util.ArrayList;")
		beans, imports, typeName, err := toJavaArray(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		set.AddStr(imports...)
		return beans, set.KeysStr(), typeName, nil
	case spec.InterfaceType:
		return nil, nil, "Object", nil
	case spec.Type:
		beans, err := getBean(parentPackage, v)
		if err != nil {
			return nil, nil, "", err
		}

		imp, typeName := toJavaType(parentPackage, v.Name())
		set.AddStr(imp)
		return beans, set.KeysStr(), typeName, nil
	case Type:
		return nil, nil, v.Name, nil
	default:
		return nil, nil, "", fmt.Errorf("unsupported type(%s): %v", reflect.TypeOf(expr).Name(), v)
	}
}

func toJavaArray(parentPackage string, a spec.ArrayType) ([]*Bean, []string, string, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, a.Value)
	if err != nil {
		return nil, nil, "", err
	}

	return beans, imports, fmt.Sprintf("ArrayList<%s>", typeName), nil
}

func toJavaMap(parentPackage string, m spec.MapType) ([]*Bean, []string, string, error) {
	beans, imports, typeName, err := getTypeName(parentPackage, m.Value)
	if err != nil {
		return nil, nil, "", err
	}

	return beans, imports, fmt.Sprintf("HashMap<String,%s>", typeName), nil
}

func toJavaType(parentPackage, goType string) (string, string) {
	switch goType {
	case "bool":
		return "", "boolean"
	case "uint8", "uint16", "uint32", "int8", "int16", "int32", "int", "uint", "byte":
		return "", "int"
	case "uint64", "int64":
		return "", "long"
	case "float32":
		return "", "float"
	case "float64":
		return "", "double"
	case "string":
		return "", "String"
	default:
		return fmt.Sprintf("import %s.bean.%s;", parentPackage, goType), goType
	}
}

func NewTag(tagExpr string) *Tag {
	tagExpr = strings.ReplaceAll(tagExpr, "`", "")
	tagExpr = strings.ReplaceAll(tagExpr, `"`, "")
	commaIndex := strings.Index(tagExpr, ",")
	if commaIndex > 0 {
		tagExpr = tagExpr[:commaIndex]
	}
	splits := strings.Split(tagExpr, ":")
	var (
		tp, name string
	)
	if len(splits) == 2 {
		tp = splits[0]
		name = splits[1]
	}

	return &Tag{
		tp:   tp,
		name: name,
	}
}

func (t *Tag) IsJson() bool {
	return t.tp == "json"
}

func (t *Tag) IsPath() bool {
	return t.tp == "path"
}

func (t *Tag) IsForm() bool {
	return t.tp == "form"
}

func (t *Tag) IsHeader() bool {
	return t.tp == "header"
}

func (t *Tag) GetTag() string {
	return t.name
}
