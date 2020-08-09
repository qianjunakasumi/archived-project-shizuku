/***********************************************************************************************************************
***  P R O J E C T  --  S H I Z U K U                                                   Q I A N J U N A K A S U M I  ***
************************************************************************************************************************
* Basic:
*
*   Package Name : shizuku
*   File Name    : main.go
*   File Path    : internal/shizuku/
*   Author       : Qianjunakasumi
*   Description  : 启动 UEHARA 应用
*
*----------------------------------------------------------------------------------------------------------------------*
* Summary:
*   func Start() error -- 启动 UEHARA 连接，初始化应用信息
*
*----------------------------------------------------------------------------------------------------------------------*
* Copyright:
*
*   Copyright (C) 2020-present QianjuNakasumi
*
*   E-mail qianjunakasumi@gmail.com
*
*   This program is free software: you can redistribute it and/or modify
*   it under the terms of the GNU Affero General Public License as published
*   by the Free Software Foundation, either version 3 of the License, or
*   (at your option) any later version.
*
*   This program is distributed in the hope that it will be useful,
*   but WITHOUT ANY WARRANTY; without even the implied warranty of
*   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*   GNU Affero General Public License for more details.
*
*   You should have received a copy of the GNU Affero General Public License
*   along with this program.  If not, see https://github.com/qianjunakasumi/project-shizuku/blob/master/LICENSE.
*----------------------------------------------------------------------------------------------------------------------*/

package shizuku

import "github.com/qianjunakasumi/project-shizuku/internal/uehara"

func Start() error {

	if err := uehara.Connect(); err != nil {

		return err

	}

	uehara.InitApp()

	return nil

}
