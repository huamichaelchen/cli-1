// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeCreateServiceBrokerActor struct {
	CreateServiceBrokerStub        func(string, string, string, string, string) (v2action.ServiceBroker, v2action.Warnings, error)
	createServiceBrokerMutex       sync.RWMutex
	createServiceBrokerArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	createServiceBrokerReturns struct {
		result1 v2action.ServiceBroker
		result2 v2action.Warnings
		result3 error
	}
	createServiceBrokerReturnsOnCall map[int]struct {
		result1 v2action.ServiceBroker
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBroker(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (v2action.ServiceBroker, v2action.Warnings, error) {
	fake.createServiceBrokerMutex.Lock()
	ret, specificReturn := fake.createServiceBrokerReturnsOnCall[len(fake.createServiceBrokerArgsForCall)]
	fake.createServiceBrokerArgsForCall = append(fake.createServiceBrokerArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreateServiceBroker", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createServiceBrokerMutex.Unlock()
	if fake.CreateServiceBrokerStub != nil {
		return fake.CreateServiceBrokerStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createServiceBrokerReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerCallCount() int {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	return len(fake.createServiceBrokerArgsForCall)
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerCalls(stub func(string, string, string, string, string) (v2action.ServiceBroker, v2action.Warnings, error)) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = stub
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerArgsForCall(i int) (string, string, string, string, string) {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	argsForCall := fake.createServiceBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerReturns(result1 v2action.ServiceBroker, result2 v2action.Warnings, result3 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	fake.createServiceBrokerReturns = struct {
		result1 v2action.ServiceBroker
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateServiceBrokerActor) CreateServiceBrokerReturnsOnCall(i int, result1 v2action.ServiceBroker, result2 v2action.Warnings, result3 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	if fake.createServiceBrokerReturnsOnCall == nil {
		fake.createServiceBrokerReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceBroker
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createServiceBrokerReturnsOnCall[i] = struct {
		result1 v2action.ServiceBroker
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateServiceBrokerActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateServiceBrokerActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v6.CreateServiceBrokerActor = new(FakeCreateServiceBrokerActor)