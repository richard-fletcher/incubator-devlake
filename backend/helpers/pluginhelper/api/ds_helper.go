/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/srvhelper"
)

func NewDataSourceHelpers[
	C plugin.ToolLayerConnection,
	S plugin.ToolLayerScope,
	SC plugin.ToolLayerScopeConfig,
](
	basicRes context.BasicRes,
	pluginName string,
	scopeSearchColumns []string,
) (
	*srvhelper.ConnectionSrvHelper[C, S, SC],
	*DsConnectionApiHelper[C, S, SC],
	*srvhelper.ScopeSrvHelper[C, S, SC],
	*DsScopeApiHelper[C, S, SC],
	*srvhelper.ScopeConfigSrvHelper[C, S, SC],
	*DsScopeConfigApiHelper[C, S, SC],
) {
	connSrv := srvhelper.NewConnectionSrvHelper[C, S, SC](basicRes, pluginName)
	connApi := NewDsConnectionApiHelper[C, S, SC](basicRes, connSrv)
	scopeSrv := srvhelper.NewScopeSrvHelper[C, S, SC](basicRes, pluginName, scopeSearchColumns)
	scopeApi := NewDsScopeApiHelper[C, S, SC](basicRes, scopeSrv)
	scSrv := srvhelper.NewScopeConfigSrvHelper[C, S, SC](basicRes)
	scApi := NewDsScopeConfigApiHelper[C, S, SC](basicRes, scSrv)
	return connSrv, connApi, scopeSrv, scopeApi, scSrv, scApi
}
