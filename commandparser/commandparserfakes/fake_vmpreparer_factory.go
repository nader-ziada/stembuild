// Code generated by counterfeiter. DO NOT EDIT.
package commandparserfakes

import (
	"sync"

	"github.com/cloudfoundry/stembuild/commandparser"
	"github.com/cloudfoundry/stembuild/construct/config"
)

type FakeVMPreparerFactory struct {
	NewStub        func(config.SourceConfig, commandparser.VCenterManager) (commandparser.VmConstruct, error)
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		arg1 config.SourceConfig
		arg2 commandparser.VCenterManager
	}
	newReturns struct {
		result1 commandparser.VmConstruct
		result2 error
	}
	newReturnsOnCall map[int]struct {
		result1 commandparser.VmConstruct
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVMPreparerFactory) New(arg1 config.SourceConfig, arg2 commandparser.VCenterManager) (commandparser.VmConstruct, error) {
	fake.newMutex.Lock()
	ret, specificReturn := fake.newReturnsOnCall[len(fake.newArgsForCall)]
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		arg1 config.SourceConfig
		arg2 commandparser.VCenterManager
	}{arg1, arg2})
	stub := fake.NewStub
	fakeReturns := fake.newReturns
	fake.recordInvocation("New", []interface{}{arg1, arg2})
	fake.newMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVMPreparerFactory) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeVMPreparerFactory) NewCalls(stub func(config.SourceConfig, commandparser.VCenterManager) (commandparser.VmConstruct, error)) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = stub
}

func (fake *FakeVMPreparerFactory) NewArgsForCall(i int) (config.SourceConfig, commandparser.VCenterManager) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	argsForCall := fake.newArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVMPreparerFactory) NewReturns(result1 commandparser.VmConstruct, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 commandparser.VmConstruct
		result2 error
	}{result1, result2}
}

func (fake *FakeVMPreparerFactory) NewReturnsOnCall(i int, result1 commandparser.VmConstruct, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	if fake.newReturnsOnCall == nil {
		fake.newReturnsOnCall = make(map[int]struct {
			result1 commandparser.VmConstruct
			result2 error
		})
	}
	fake.newReturnsOnCall[i] = struct {
		result1 commandparser.VmConstruct
		result2 error
	}{result1, result2}
}

func (fake *FakeVMPreparerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVMPreparerFactory) recordInvocation(key string, args []interface{}) {
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

var _ commandparser.VMPreparerFactory = new(FakeVMPreparerFactory)
