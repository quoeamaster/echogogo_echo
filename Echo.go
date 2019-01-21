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
	"github.com/buger/jsonparser"
	"github.com/quoeamaster/echogogo_plugin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type EchoModuleModel struct {
	EchoBody        string
	EchoBodyMap		map[string]interface{}
	Timestamp       time.Time
	TimestampEpoch  int64
}

/*
 * note: seems plugin api could only lookup 1st class level methods; hence if you have struct
 * implementing an interface's method... the designated methods won't be lookup-able
 *
 *	(removed params) response http.Response,
 */
func DoAction(request http.Request, endPoint string, optionalMap ...map[string]interface{}) interface{}  {
	/* usage: optionalMap[0] => 1st element in the optional parameter which is a map[string]interface{} */

	// MUST be json request... try to parse in json...
	bArray, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	modelPtr := new(EchoModuleModel)
	// have body, return this body content
	if len(bArray) > 0 {
		if strings.HasSuffix(request.URL.Path, "/xml") || strings.HasSuffix(request.URL.Path, "/xml/") {
			modelPtr.EchoBody = string(bArray)
			modelPtr.EchoBody = strings.Replace(modelPtr.EchoBody, "\n", "", -1)
			modelPtr.EchoBody = strings.Replace(modelPtr.EchoBody, "\t", "", -1)
		} else {
			m := make(map[string]interface{})
			err := jsonparser.ObjectEach(bArray, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				// fmt.Printf("idx(%v); key (%v)=> value(%v);; valueType (%v)  \n", offset, string(key), string(value), dataType)
				keyString := string(key)
				valString := string(value)
				m[keyString] = valString
				return nil
			})	// end -- iterate all fields from the root (or a given set of keys)
			if err != nil {
				return nil
			}
			modelPtr.EchoBodyMap = m
		}
	} else {
		modelPtr.EchoBody = "Hi, welcome to the echo server!!!"
	}
	modelPtr.Timestamp = time.Now().UTC()
	modelPtr.TimestampEpoch = modelPtr.Timestamp.UnixNano()

	return *modelPtr
}

// endpoint format => [http action]::[path] => e.g. GET::/ or POST::/name
func GetRestConfig() map[string]interface{} {
	/* TODO: either read from a file or simply overwrite it programmatically.... */
	mapModelPtr := make(map[string]interface{})

	mapModelPtr["consumeFormat"] = echogogo.FORMAT_JSON
	mapModelPtr["produceFormat"] = echogogo.FORMAT_XML_JSON
	mapModelPtr["path"] = "/echo"
	mapModelPtr["endPoints"] = []string {
		"GET::/", "GET::/json", "GET::/xml",
		"POST::/", "POST::/json", "POST::/xml",
	}

	return mapModelPtr
}
