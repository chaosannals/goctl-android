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

package action

import (
	"encoding/json"
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/goctl-android/generate"
)

func Android(ctx *cli.Context) error {
	pkg := ctx.String("package")
	std, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var plugin generate.Plugin
	plugin.ParentPackage = pkg
	err = json.Unmarshal(std, &plugin)
	if err != nil {
		return err
	}
	api, err := parser.Parse(plugin.ApiFilePath)
	if err != nil {
		return err
	}
	plugin.Api = api

	return generate.Do(plugin)
}
