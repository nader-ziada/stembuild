// Code generated by counterfeiter. DO NOT EDIT.
package vcenter_managerfakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry/stembuild/iaas_cli/iaas_clients/vcenter_manager"
	"github.com/vmware/govmomi/guest"
)

type FakeOpsManager struct {
	FileManagerStub        func(context.Context) (*guest.FileManager, error)
	fileManagerMutex       sync.RWMutex
	fileManagerArgsForCall []struct {
		arg1 context.Context
	}
	fileManagerReturns struct {
		result1 *guest.FileManager
		result2 error
	}
	fileManagerReturnsOnCall map[int]struct {
		result1 *guest.FileManager
		result2 error
	}
	ProcessManagerStub        func(context.Context) (*guest.ProcessManager, error)
	processManagerMutex       sync.RWMutex
	processManagerArgsForCall []struct {
		arg1 context.Context
	}
	processManagerReturns struct {
		result1 *guest.ProcessManager
		result2 error
	}
	processManagerReturnsOnCall map[int]struct {
		result1 *guest.ProcessManager
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpsManager) FileManager(arg1 context.Context) (*guest.FileManager, error) {
	fake.fileManagerMutex.Lock()
	ret, specificReturn := fake.fileManagerReturnsOnCall[len(fake.fileManagerArgsForCall)]
	fake.fileManagerArgsForCall = append(fake.fileManagerArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FileManagerStub
	fakeReturns := fake.fileManagerReturns
	fake.recordInvocation("FileManager", []interface{}{arg1})
	fake.fileManagerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) FileManagerCallCount() int {
	fake.fileManagerMutex.RLock()
	defer fake.fileManagerMutex.RUnlock()
	return len(fake.fileManagerArgsForCall)
}

func (fake *FakeOpsManager) FileManagerCalls(stub func(context.Context) (*guest.FileManager, error)) {
	fake.fileManagerMutex.Lock()
	defer fake.fileManagerMutex.Unlock()
	fake.FileManagerStub = stub
}

func (fake *FakeOpsManager) FileManagerArgsForCall(i int) context.Context {
	fake.fileManagerMutex.RLock()
	defer fake.fileManagerMutex.RUnlock()
	argsForCall := fake.fileManagerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsManager) FileManagerReturns(result1 *guest.FileManager, result2 error) {
	fake.fileManagerMutex.Lock()
	defer fake.fileManagerMutex.Unlock()
	fake.FileManagerStub = nil
	fake.fileManagerReturns = struct {
		result1 *guest.FileManager
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) FileManagerReturnsOnCall(i int, result1 *guest.FileManager, result2 error) {
	fake.fileManagerMutex.Lock()
	defer fake.fileManagerMutex.Unlock()
	fake.FileManagerStub = nil
	if fake.fileManagerReturnsOnCall == nil {
		fake.fileManagerReturnsOnCall = make(map[int]struct {
			result1 *guest.FileManager
			result2 error
		})
	}
	fake.fileManagerReturnsOnCall[i] = struct {
		result1 *guest.FileManager
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) ProcessManager(arg1 context.Context) (*guest.ProcessManager, error) {
	fake.processManagerMutex.Lock()
	ret, specificReturn := fake.processManagerReturnsOnCall[len(fake.processManagerArgsForCall)]
	fake.processManagerArgsForCall = append(fake.processManagerArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ProcessManagerStub
	fakeReturns := fake.processManagerReturns
	fake.recordInvocation("ProcessManager", []interface{}{arg1})
	fake.processManagerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) ProcessManagerCallCount() int {
	fake.processManagerMutex.RLock()
	defer fake.processManagerMutex.RUnlock()
	return len(fake.processManagerArgsForCall)
}

func (fake *FakeOpsManager) ProcessManagerCalls(stub func(context.Context) (*guest.ProcessManager, error)) {
	fake.processManagerMutex.Lock()
	defer fake.processManagerMutex.Unlock()
	fake.ProcessManagerStub = stub
}

func (fake *FakeOpsManager) ProcessManagerArgsForCall(i int) context.Context {
	fake.processManagerMutex.RLock()
	defer fake.processManagerMutex.RUnlock()
	argsForCall := fake.processManagerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsManager) ProcessManagerReturns(result1 *guest.ProcessManager, result2 error) {
	fake.processManagerMutex.Lock()
	defer fake.processManagerMutex.Unlock()
	fake.ProcessManagerStub = nil
	fake.processManagerReturns = struct {
		result1 *guest.ProcessManager
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) ProcessManagerReturnsOnCall(i int, result1 *guest.ProcessManager, result2 error) {
	fake.processManagerMutex.Lock()
	defer fake.processManagerMutex.Unlock()
	fake.ProcessManagerStub = nil
	if fake.processManagerReturnsOnCall == nil {
		fake.processManagerReturnsOnCall = make(map[int]struct {
			result1 *guest.ProcessManager
			result2 error
		})
	}
	fake.processManagerReturnsOnCall[i] = struct {
		result1 *guest.ProcessManager
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileManagerMutex.RLock()
	defer fake.fileManagerMutex.RUnlock()
	fake.processManagerMutex.RLock()
	defer fake.processManagerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpsManager) recordInvocation(key string, args []interface{}) {
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

var _ vcenter_manager.OpsManager = new(FakeOpsManager)
