# autok3s
[![Build Status](http://drone-pandaria.cnrancher.com/api/badges/cnrancher/autok3s/status.svg)](http://drone-pandaria.cnrancher.com/cnrancher/autok3s)
[![Go Report Card](https://goreportcard.com/badge/github.com/cnrancher/autok3s)](https://goreportcard.com/report/github.com/cnrancher/autok3s) 
![GitHub release](https://img.shields.io/github/v/release/cnrancher/autok3s.svg?color=default)
[![License: apache-2.0](https://img.shields.io/badge/License-apache2-default.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://github.com/cnrancher/autok3s/pulls)

AutoK3s是用于在多个公有云平台上快速创建和管理K3s群集的轻量级工具。它可以帮助用户快速完成K3s集群的个性化配置，同时提供便捷的`kubectl`访问功能。

其他语言:
- [English](README.md)
- [Chinese Simplified (中文简体)](README_zhCN.md)

## 设计理念
该工具使用云厂商的SDK来创建和管理主机，然后使用SSH将K3s群集安装到远程主机。您也可以使用它将主机作为`masters/agents`节点加入K3s集群。同时自动将`kubeconfig`合并并存储在`$HOME/.autok3s/.kube/config`中，这对于用户访问群集是必需的。
然后用户可以使用`autok3s kubectl`命令快速访问集群。

使用 [viper](https://github.com/spf13/viper) 绑定参数和配置文件。 autok3s会生成一个配置文件，将云厂商的访问信息存储在指定位置（`$HOME/.autok3s/config.yaml`），以减少多次运行要传递的参数的数量。

集群成功创建后，会生成位于`$HOME/.autok3s/.state`目录下的状态文件，以记录在此主机上创建的集群信息。

## 已支持云厂商
有关更多用法的详细信息，请参见下面的链接：

- [alibaba](docs/alibaba/README_zhCN.md) - 使用阿里云SDK管理主机，然后使用SSH安装或加入K3s集群和主机。
- [tencent](docs/tencent/README_zhCN.md) - 使用腾讯云SDK管理主机，然后使用SSH安装或加入K3s集群和主机。
- [native](docs/native/README_zhCN.md) - 不集成Cloud SDK，仅使用SSH来安装或加入K3s集群和主机。

## 演示视频
使用命令 `autok3s -d create --provider alibaba` 创建K3s集群。

[![asciicast](https://asciinema.org/a/whwyjSfGv7lZdjaenTDCRejDW.svg)](https://asciinema.org/a/whwyjSfGv7lZdjaenTDCRejDW)

## 开发者指南
使用 `Makefile` 管理项目的编译、测试与打包。
项目支持使用 `dapper`，`dapper`安装步骤请参考[dapper](https://github.com/rancher/dapper)。

- 更新依赖: `GO111MODULE=on go mod vendor`
- 编译: `BY=dapper make autok3s`
- 测试: `BY=dapper make autok3s unit`
- 打包: `BY=dapper make autok3s package only`

# 开源协议

Copyright (c) 2020 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
