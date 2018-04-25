package registry

import "gitlab.alipay-inc.com/afe/mosn/pkg/upstream/servicediscovery/confreg/servermanager"

const CONFREG_SOFA_GROUP  = "SOFA"


type RegistryClient interface {
    PublishAsync(dataId string, data ...string)

    UnPublishAsync(dataId string, data ...string)

    SubscribeAsync(dataId string)

    UnSubscribeAsync(dataId string)

    PublishSync(dataId string, data ...string) error

    UnPublishSync(dataId string, data ...string) error

    SubscribeSync(dataId string) error

    UnSubscribeSync(dataId string) error

    GetRPCServerManager() servermanager.RPCServerManager
}