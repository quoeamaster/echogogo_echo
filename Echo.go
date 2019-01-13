/*
 * Licensed to Echogogo under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Echogogo licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package echogogo_echo

import "net/http"
import "github.com/quoeamaster/echogogo_plugin"

type EchoModule struct {

}

func (m *EchoModule) DoAction(request http.Request, response http.Response, endPoint string, optionalMap ...map[string]interface{}) interface{}  {
	/* usage: optionalMap[0] => 1st element in the optional parameter which is a map[string]interface{} */
	return nil
}

func (m *EchoModule) GetRestConfig() echogogo.RestConfigModel {
	/* either read from a file or simply overwrite it programmatically.... */
	modelPtr := new(echogogo.RestConfigModel)

	modelPtr.ConsumeFormat = echogogo.FORMAT_JSON
	modelPtr.ProduceFormat = echogogo.FORMAT_XML_JSON
	modelPtr.Path = "/echo"
	modelPtr.EndPoints = []string { "/", "/json", "/xml" }

	return *modelPtr
}
