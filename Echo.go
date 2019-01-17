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
package main

import (
	"github.com/quoeamaster/echogogo_plugin"
	"net/http"
)

type EchoModule struct {

}

/* (m *EchoModule)  */
/*
 * note: seems plugin api could only lookup 1st class level methods; hence if you have struct
 * implementing an interface's method... the designated methods won't be lookup-able
 */
func DoAction(request http.Request, response http.Response, endPoint string, optionalMap ...map[string]interface{}) interface{}  {
	/* usage: optionalMap[0] => 1st element in the optional parameter which is a map[string]interface{} */
	return nil
}

/* (m *EchoModule)  */
// TODO: change it back to a Map instead???
// TODO: create another method with a String param => see if it works or not... OR map param
func GetRestConfig() map[string]interface{} {
	/* either read from a file or simply overwrite it programmatically.... */
	/*
	modelPtr := new(echogogo.RestConfigModel)

	modelPtr.ConsumeFormat = echogogo.FORMAT_JSON
	modelPtr.ProduceFormat = echogogo.FORMAT_XML_JSON
	modelPtr.Path = "/echo"
	modelPtr.EndPoints = []string { "/", "/json", "/xml" }

	return *modelPtr
	*/
	mapModelPtr := make(map[string]interface{})

	mapModelPtr["consumeFormat"] = echogogo.FORMAT_JSON
	mapModelPtr["produceFormat"] = echogogo.FORMAT_XML_JSON
	mapModelPtr["path"] = "/echo"
	mapModelPtr["endPoints"] = []string { "/", "/json", "/xml" }

	return mapModelPtr
}
