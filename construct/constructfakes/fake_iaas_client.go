// Code generated by counterfeiter. DO NOT EDIT.
package constructfakes

import (
	"sync"

	"github.com/cloudfoundry/stembuild/construct"
)

type FakeIaasClient struct {
	IsPoweredOffStub        func(string) (bool, error)
	isPoweredOffMutex       sync.RWMutex
	isPoweredOffArgsForCall []struct {
		arg1 string
	}
	isPoweredOffReturns struct {
		result1 bool
		result2 error
	}
	isPoweredOffReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	MakeDirectoryStub        func(string, string, string, string) error
	makeDirectoryMutex       sync.RWMutex
	makeDirectoryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	makeDirectoryReturns struct {
		result1 error
	}
	makeDirectoryReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(string, string, string, string, ...string) (string, error)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 []string
	}
	startReturns struct {
		result1 string
		result2 error
	}
	startReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UploadArtifactStub        func(string, string, string, string, string) error
	uploadArtifactMutex       sync.RWMutex
	uploadArtifactArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	uploadArtifactReturns struct {
		result1 error
	}
	uploadArtifactReturnsOnCall map[int]struct {
		result1 error
	}
	WaitForExitStub        func(string, string, string, string) (int, error)
	waitForExitMutex       sync.RWMutex
	waitForExitArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	waitForExitReturns struct {
		result1 int
		result2 error
	}
	waitForExitReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIaasClient) IsPoweredOff(arg1 string) (bool, error) {
	fake.isPoweredOffMutex.Lock()
	ret, specificReturn := fake.isPoweredOffReturnsOnCall[len(fake.isPoweredOffArgsForCall)]
	fake.isPoweredOffArgsForCall = append(fake.isPoweredOffArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsPoweredOffStub
	fakeReturns := fake.isPoweredOffReturns
	fake.recordInvocation("IsPoweredOff", []interface{}{arg1})
	fake.isPoweredOffMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIaasClient) IsPoweredOffCallCount() int {
	fake.isPoweredOffMutex.RLock()
	defer fake.isPoweredOffMutex.RUnlock()
	return len(fake.isPoweredOffArgsForCall)
}

func (fake *FakeIaasClient) IsPoweredOffCalls(stub func(string) (bool, error)) {
	fake.isPoweredOffMutex.Lock()
	defer fake.isPoweredOffMutex.Unlock()
	fake.IsPoweredOffStub = stub
}

func (fake *FakeIaasClient) IsPoweredOffArgsForCall(i int) string {
	fake.isPoweredOffMutex.RLock()
	defer fake.isPoweredOffMutex.RUnlock()
	argsForCall := fake.isPoweredOffArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIaasClient) IsPoweredOffReturns(result1 bool, result2 error) {
	fake.isPoweredOffMutex.Lock()
	defer fake.isPoweredOffMutex.Unlock()
	fake.IsPoweredOffStub = nil
	fake.isPoweredOffReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) IsPoweredOffReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isPoweredOffMutex.Lock()
	defer fake.isPoweredOffMutex.Unlock()
	fake.IsPoweredOffStub = nil
	if fake.isPoweredOffReturnsOnCall == nil {
		fake.isPoweredOffReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isPoweredOffReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) MakeDirectory(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.makeDirectoryMutex.Lock()
	ret, specificReturn := fake.makeDirectoryReturnsOnCall[len(fake.makeDirectoryArgsForCall)]
	fake.makeDirectoryArgsForCall = append(fake.makeDirectoryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.MakeDirectoryStub
	fakeReturns := fake.makeDirectoryReturns
	fake.recordInvocation("MakeDirectory", []interface{}{arg1, arg2, arg3, arg4})
	fake.makeDirectoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIaasClient) MakeDirectoryCallCount() int {
	fake.makeDirectoryMutex.RLock()
	defer fake.makeDirectoryMutex.RUnlock()
	return len(fake.makeDirectoryArgsForCall)
}

func (fake *FakeIaasClient) MakeDirectoryCalls(stub func(string, string, string, string) error) {
	fake.makeDirectoryMutex.Lock()
	defer fake.makeDirectoryMutex.Unlock()
	fake.MakeDirectoryStub = stub
}

func (fake *FakeIaasClient) MakeDirectoryArgsForCall(i int) (string, string, string, string) {
	fake.makeDirectoryMutex.RLock()
	defer fake.makeDirectoryMutex.RUnlock()
	argsForCall := fake.makeDirectoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeIaasClient) MakeDirectoryReturns(result1 error) {
	fake.makeDirectoryMutex.Lock()
	defer fake.makeDirectoryMutex.Unlock()
	fake.MakeDirectoryStub = nil
	fake.makeDirectoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) MakeDirectoryReturnsOnCall(i int, result1 error) {
	fake.makeDirectoryMutex.Lock()
	defer fake.makeDirectoryMutex.Unlock()
	fake.MakeDirectoryStub = nil
	if fake.makeDirectoryReturnsOnCall == nil {
		fake.makeDirectoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeDirectoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) Start(arg1 string, arg2 string, arg3 string, arg4 string, arg5 ...string) (string, error) {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 []string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIaasClient) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeIaasClient) StartCalls(stub func(string, string, string, string, ...string) (string, error)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeIaasClient) StartArgsForCall(i int) (string, string, string, string, []string) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeIaasClient) StartReturns(result1 string, result2 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) StartReturnsOnCall(i int, result1 string, result2 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) UploadArtifact(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.uploadArtifactMutex.Lock()
	ret, specificReturn := fake.uploadArtifactReturnsOnCall[len(fake.uploadArtifactArgsForCall)]
	fake.uploadArtifactArgsForCall = append(fake.uploadArtifactArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.UploadArtifactStub
	fakeReturns := fake.uploadArtifactReturns
	fake.recordInvocation("UploadArtifact", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.uploadArtifactMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIaasClient) UploadArtifactCallCount() int {
	fake.uploadArtifactMutex.RLock()
	defer fake.uploadArtifactMutex.RUnlock()
	return len(fake.uploadArtifactArgsForCall)
}

func (fake *FakeIaasClient) UploadArtifactCalls(stub func(string, string, string, string, string) error) {
	fake.uploadArtifactMutex.Lock()
	defer fake.uploadArtifactMutex.Unlock()
	fake.UploadArtifactStub = stub
}

func (fake *FakeIaasClient) UploadArtifactArgsForCall(i int) (string, string, string, string, string) {
	fake.uploadArtifactMutex.RLock()
	defer fake.uploadArtifactMutex.RUnlock()
	argsForCall := fake.uploadArtifactArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeIaasClient) UploadArtifactReturns(result1 error) {
	fake.uploadArtifactMutex.Lock()
	defer fake.uploadArtifactMutex.Unlock()
	fake.UploadArtifactStub = nil
	fake.uploadArtifactReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) UploadArtifactReturnsOnCall(i int, result1 error) {
	fake.uploadArtifactMutex.Lock()
	defer fake.uploadArtifactMutex.Unlock()
	fake.UploadArtifactStub = nil
	if fake.uploadArtifactReturnsOnCall == nil {
		fake.uploadArtifactReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadArtifactReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) WaitForExit(arg1 string, arg2 string, arg3 string, arg4 string) (int, error) {
	fake.waitForExitMutex.Lock()
	ret, specificReturn := fake.waitForExitReturnsOnCall[len(fake.waitForExitArgsForCall)]
	fake.waitForExitArgsForCall = append(fake.waitForExitArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.WaitForExitStub
	fakeReturns := fake.waitForExitReturns
	fake.recordInvocation("WaitForExit", []interface{}{arg1, arg2, arg3, arg4})
	fake.waitForExitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIaasClient) WaitForExitCallCount() int {
	fake.waitForExitMutex.RLock()
	defer fake.waitForExitMutex.RUnlock()
	return len(fake.waitForExitArgsForCall)
}

func (fake *FakeIaasClient) WaitForExitCalls(stub func(string, string, string, string) (int, error)) {
	fake.waitForExitMutex.Lock()
	defer fake.waitForExitMutex.Unlock()
	fake.WaitForExitStub = stub
}

func (fake *FakeIaasClient) WaitForExitArgsForCall(i int) (string, string, string, string) {
	fake.waitForExitMutex.RLock()
	defer fake.waitForExitMutex.RUnlock()
	argsForCall := fake.waitForExitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeIaasClient) WaitForExitReturns(result1 int, result2 error) {
	fake.waitForExitMutex.Lock()
	defer fake.waitForExitMutex.Unlock()
	fake.WaitForExitStub = nil
	fake.waitForExitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) WaitForExitReturnsOnCall(i int, result1 int, result2 error) {
	fake.waitForExitMutex.Lock()
	defer fake.waitForExitMutex.Unlock()
	fake.WaitForExitStub = nil
	if fake.waitForExitReturnsOnCall == nil {
		fake.waitForExitReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.waitForExitReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIaasClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isPoweredOffMutex.RLock()
	defer fake.isPoweredOffMutex.RUnlock()
	fake.makeDirectoryMutex.RLock()
	defer fake.makeDirectoryMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.uploadArtifactMutex.RLock()
	defer fake.uploadArtifactMutex.RUnlock()
	fake.waitForExitMutex.RLock()
	defer fake.waitForExitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIaasClient) recordInvocation(key string, args []interface{}) {
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

var _ construct.IaasClient = new(FakeIaasClient)
