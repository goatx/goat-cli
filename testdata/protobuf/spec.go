package protobuf

import (
	"context"

	"github.com/goatx/goat"
	"github.com/goatx/goat/protobuf"
)

type ClientStateMachine struct {
	goat.StateMachine
	Service *ServiceStateMachine
}

type ServiceStateMachine struct {
	goat.StateMachine
	Client *ClientStateMachine
}

type ClientState struct {
	goat.State
}

type ServiceState struct {
	goat.State
}

type Request struct {
	protobuf.Message[*ClientStateMachine, *ServiceStateMachine]
}

type Response struct {
	protobuf.Message[*ServiceStateMachine, *ClientStateMachine]
}

func createProtobufModel() {
	clientSpec := goat.NewStateMachineSpec(&ClientStateMachine{})
	serviceSpec := protobuf.NewServiceSpec(&ServiceStateMachine{})

	clientState := &ClientState{}
	serviceState := &ServiceState{}

	clientSpec.DefineStates(clientState).SetInitialState(clientState)
	serviceSpec.DefineStates(serviceState).SetInitialState(serviceState)

	goat.OnEntry(clientSpec, clientState,
		func(ctx context.Context, client *ClientStateMachine) {
			request := &Request{}
			protobuf.SendTo(ctx, client.Service, request)
		})

	protobuf.OnMessage(serviceSpec, serviceState, "HandleRequest",
		func(ctx context.Context, request *Request, service *ServiceStateMachine) protobuf.Response[*Response] {
			response := &Response{}
			return protobuf.SendTo(ctx, service.Client, response)
		})
}
