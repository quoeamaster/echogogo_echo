package echogogo_echo

import "net/http"
import "github.com/quoeamaster/echogogo_plugin"

type EchoModule struct {

}

func (m *EchoModule) DoAction(request http.Request, response http.Response, endPoint string, optionalMap ...map[string]interface{}) interface{}  {
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