package echogogo_echo

import "net/http"
import "github.com/quoeamaster/echogogo_plugin"

type EchoModule struct {

}

func (m *EchoModule) DoAction(request http.Request, response http.Response) interface{}  {
	return nil
}

func (m *EchoModule) GetRestConfig() echogogo.RestConfigModel {
	return *new(echogogo.RestConfigModel)
}