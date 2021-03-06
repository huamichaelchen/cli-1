// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3CancelZdtPushActor struct {
	CancelDeploymentByAppNameAndSpaceStub        func(string, string) (v3action.Warnings, error)
	cancelDeploymentByAppNameAndSpaceMutex       sync.RWMutex
	cancelDeploymentByAppNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	cancelDeploymentByAppNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	cancelDeploymentByAppNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpace(arg1 string, arg2 string) (v3action.Warnings, error) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall[len(fake.cancelDeploymentByAppNameAndSpaceArgsForCall)]
	fake.cancelDeploymentByAppNameAndSpaceArgsForCall = append(fake.cancelDeploymentByAppNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CancelDeploymentByAppNameAndSpace", []interface{}{arg1, arg2})
	fake.cancelDeploymentByAppNameAndSpaceMutex.Unlock()
	if fake.CancelDeploymentByAppNameAndSpaceStub != nil {
		return fake.CancelDeploymentByAppNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.cancelDeploymentByAppNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceCallCount() int {
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	return len(fake.cancelDeploymentByAppNameAndSpaceArgsForCall)
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceCalls(stub func(string, string) (v3action.Warnings, error)) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.Lock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.Unlock()
	fake.CancelDeploymentByAppNameAndSpaceStub = stub
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceArgsForCall(i int) (string, string) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	argsForCall := fake.cancelDeploymentByAppNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.Lock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.Unlock()
	fake.CancelDeploymentByAppNameAndSpaceStub = nil
	fake.cancelDeploymentByAppNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3CancelZdtPushActor) CancelDeploymentByAppNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.cancelDeploymentByAppNameAndSpaceMutex.Lock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.Unlock()
	fake.CancelDeploymentByAppNameAndSpaceStub = nil
	if fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall == nil {
		fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.cancelDeploymentByAppNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3CancelZdtPushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelDeploymentByAppNameAndSpaceMutex.RLock()
	defer fake.cancelDeploymentByAppNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3CancelZdtPushActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3CancelZdtPushActor = new(FakeV3CancelZdtPushActor)
