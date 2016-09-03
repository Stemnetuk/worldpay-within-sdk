package event

import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"

type Handler interface {
	BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int)
	EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int)
}