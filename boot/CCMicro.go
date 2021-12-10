package boot

import (
	"encoding/json"
	"fmt"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"github.com/asim/go-micro/plugins/registry/nacos/v4"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/server"
	"os"
	"strings"
)

func Micro(args []string) {
	event.TriggerEvent(event.Start)
	marshal, _ := json.Marshal(args)
	cclog.Debug(fmt.Sprintf("args=%s", string(marshal)))
	service := CreateServiceWithHttpServer()
	service.Init()
	configRouter(service.Server())
	// Run service
	if err := service.Run(); err != nil {
		cclog.Error(err.Error())
		os.Exit(0)
	}
}

func configRouter(server server.Server) {

	hd := server.NewHandler(web.Engine)
	if err := server.Handle(hd); err != nil {
		cclog.Error(err.Error())
	}
}

func CreateRegistry() registry.Registry {
	clientConfig := constant.ClientConfig{}
	err := viper.UnmarshalKey("cc.cloud.nacos.config", &clientConfig)
	if err != nil {
		return nil
	}
	addr := config.GetStringOrDefault("cc.cloud.nacos.server-addr", "localhost:8848")
	addrs := strings.Split(addr, ",")
	return nacos.NewRegistry(nacos.WithAddress(addrs), nacos.WithClientConfig(clientConfig))
}

func CreateServiceWithHttpServer() micro.Service {
	webAddr := config.GetStringOrDefault("server.port", "8080")
	appName := config.GetStringOrDefault("cc.application.name", "go.micro")
	srv := httpServer.NewServer(
		server.Name(appName),
		server.Address(":"+webAddr),
	)

	return micro.NewService(
		micro.Server(srv),
		micro.Name(appName),
		micro.Registry(CreateRegistry()),
	)
}

func CreateService() micro.Service {
	webAddr := config.GetStringOrDefault("server.port", "8081")
	appName := config.GetStringOrDefault("cc.application.name", "go.micro")
	return micro.NewService(
		micro.Address(":"+webAddr),
		micro.Name(appName),
		micro.Registry(CreateRegistry()),
	)
}

func CreateClient() client.Client {

	s := selector.NewSelector(selector.Registry(CreateRegistry()))

	return httpClient.NewClient(client.Selector(s),
		client.ContentType("application/json"))
}
