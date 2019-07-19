// Code generated by counterfeiter. DO NOT EDIT.
package commandparserfakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry-incubator/stembuild/commandparser"
	"github.com/cloudfoundry-incubator/stembuild/iaas_cli/iaas_clients/guest_manager"
	"github.com/cloudfoundry-incubator/stembuild/iaas_cli/iaas_clients/vcenter_manager"
	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/object"
)

type FakeVCenterManager struct {
	FindVMStub        func(context.Context, string) (*object.VirtualMachine, error)
	findVMMutex       sync.RWMutex
	findVMArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	findVMReturns struct {
		result1 *object.VirtualMachine
		result2 error
	}
	findVMReturnsOnCall map[int]struct {
		result1 *object.VirtualMachine
		result2 error
	}
	GuestManagerStub        func(context.Context, vcenter_manager.OpsManager, string, string) (*guest_manager.GuestManager, error)
	guestManagerMutex       sync.RWMutex
	guestManagerArgsForCall []struct {
		arg1 context.Context
		arg2 vcenter_manager.OpsManager
		arg3 string
		arg4 string
	}
	guestManagerReturns struct {
		result1 *guest_manager.GuestManager
		result2 error
	}
	guestManagerReturnsOnCall map[int]struct {
		result1 *guest_manager.GuestManager
		result2 error
	}
	LoginStub        func(context.Context) error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 context.Context
	}
	loginReturns struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	OperationsManagerStub        func(context.Context, *object.VirtualMachine) *guest.OperationsManager
	operationsManagerMutex       sync.RWMutex
	operationsManagerArgsForCall []struct {
		arg1 context.Context
		arg2 *object.VirtualMachine
	}
	operationsManagerReturns struct {
		result1 *guest.OperationsManager
	}
	operationsManagerReturnsOnCall map[int]struct {
		result1 *guest.OperationsManager
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVCenterManager) FindVM(arg1 context.Context, arg2 string) (*object.VirtualMachine, error) {
	fake.findVMMutex.Lock()
	ret, specificReturn := fake.findVMReturnsOnCall[len(fake.findVMArgsForCall)]
	fake.findVMArgsForCall = append(fake.findVMArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FindVM", []interface{}{arg1, arg2})
	fake.findVMMutex.Unlock()
	if fake.FindVMStub != nil {
		return fake.FindVMStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findVMReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVCenterManager) FindVMCallCount() int {
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	return len(fake.findVMArgsForCall)
}

func (fake *FakeVCenterManager) FindVMCalls(stub func(context.Context, string) (*object.VirtualMachine, error)) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = stub
}

func (fake *FakeVCenterManager) FindVMArgsForCall(i int) (context.Context, string) {
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	argsForCall := fake.findVMArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVCenterManager) FindVMReturns(result1 *object.VirtualMachine, result2 error) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = nil
	fake.findVMReturns = struct {
		result1 *object.VirtualMachine
		result2 error
	}{result1, result2}
}

func (fake *FakeVCenterManager) FindVMReturnsOnCall(i int, result1 *object.VirtualMachine, result2 error) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = nil
	if fake.findVMReturnsOnCall == nil {
		fake.findVMReturnsOnCall = make(map[int]struct {
			result1 *object.VirtualMachine
			result2 error
		})
	}
	fake.findVMReturnsOnCall[i] = struct {
		result1 *object.VirtualMachine
		result2 error
	}{result1, result2}
}

func (fake *FakeVCenterManager) GuestManager(arg1 context.Context, arg2 vcenter_manager.OpsManager, arg3 string, arg4 string) (*guest_manager.GuestManager, error) {
	fake.guestManagerMutex.Lock()
	ret, specificReturn := fake.guestManagerReturnsOnCall[len(fake.guestManagerArgsForCall)]
	fake.guestManagerArgsForCall = append(fake.guestManagerArgsForCall, struct {
		arg1 context.Context
		arg2 vcenter_manager.OpsManager
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GuestManager", []interface{}{arg1, arg2, arg3, arg4})
	fake.guestManagerMutex.Unlock()
	if fake.GuestManagerStub != nil {
		return fake.GuestManagerStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.guestManagerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVCenterManager) GuestManagerCallCount() int {
	fake.guestManagerMutex.RLock()
	defer fake.guestManagerMutex.RUnlock()
	return len(fake.guestManagerArgsForCall)
}

func (fake *FakeVCenterManager) GuestManagerCalls(stub func(context.Context, vcenter_manager.OpsManager, string, string) (*guest_manager.GuestManager, error)) {
	fake.guestManagerMutex.Lock()
	defer fake.guestManagerMutex.Unlock()
	fake.GuestManagerStub = stub
}

func (fake *FakeVCenterManager) GuestManagerArgsForCall(i int) (context.Context, vcenter_manager.OpsManager, string, string) {
	fake.guestManagerMutex.RLock()
	defer fake.guestManagerMutex.RUnlock()
	argsForCall := fake.guestManagerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVCenterManager) GuestManagerReturns(result1 *guest_manager.GuestManager, result2 error) {
	fake.guestManagerMutex.Lock()
	defer fake.guestManagerMutex.Unlock()
	fake.GuestManagerStub = nil
	fake.guestManagerReturns = struct {
		result1 *guest_manager.GuestManager
		result2 error
	}{result1, result2}
}

func (fake *FakeVCenterManager) GuestManagerReturnsOnCall(i int, result1 *guest_manager.GuestManager, result2 error) {
	fake.guestManagerMutex.Lock()
	defer fake.guestManagerMutex.Unlock()
	fake.GuestManagerStub = nil
	if fake.guestManagerReturnsOnCall == nil {
		fake.guestManagerReturnsOnCall = make(map[int]struct {
			result1 *guest_manager.GuestManager
			result2 error
		})
	}
	fake.guestManagerReturnsOnCall[i] = struct {
		result1 *guest_manager.GuestManager
		result2 error
	}{result1, result2}
}

func (fake *FakeVCenterManager) Login(arg1 context.Context) error {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Login", []interface{}{arg1})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.loginReturns
	return fakeReturns.result1
}

func (fake *FakeVCenterManager) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeVCenterManager) LoginCalls(stub func(context.Context) error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeVCenterManager) LoginArgsForCall(i int) context.Context {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVCenterManager) LoginReturns(result1 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVCenterManager) LoginReturnsOnCall(i int, result1 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVCenterManager) OperationsManager(arg1 context.Context, arg2 *object.VirtualMachine) *guest.OperationsManager {
	fake.operationsManagerMutex.Lock()
	ret, specificReturn := fake.operationsManagerReturnsOnCall[len(fake.operationsManagerArgsForCall)]
	fake.operationsManagerArgsForCall = append(fake.operationsManagerArgsForCall, struct {
		arg1 context.Context
		arg2 *object.VirtualMachine
	}{arg1, arg2})
	fake.recordInvocation("OperationsManager", []interface{}{arg1, arg2})
	fake.operationsManagerMutex.Unlock()
	if fake.OperationsManagerStub != nil {
		return fake.OperationsManagerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.operationsManagerReturns
	return fakeReturns.result1
}

func (fake *FakeVCenterManager) OperationsManagerCallCount() int {
	fake.operationsManagerMutex.RLock()
	defer fake.operationsManagerMutex.RUnlock()
	return len(fake.operationsManagerArgsForCall)
}

func (fake *FakeVCenterManager) OperationsManagerCalls(stub func(context.Context, *object.VirtualMachine) *guest.OperationsManager) {
	fake.operationsManagerMutex.Lock()
	defer fake.operationsManagerMutex.Unlock()
	fake.OperationsManagerStub = stub
}

func (fake *FakeVCenterManager) OperationsManagerArgsForCall(i int) (context.Context, *object.VirtualMachine) {
	fake.operationsManagerMutex.RLock()
	defer fake.operationsManagerMutex.RUnlock()
	argsForCall := fake.operationsManagerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVCenterManager) OperationsManagerReturns(result1 *guest.OperationsManager) {
	fake.operationsManagerMutex.Lock()
	defer fake.operationsManagerMutex.Unlock()
	fake.OperationsManagerStub = nil
	fake.operationsManagerReturns = struct {
		result1 *guest.OperationsManager
	}{result1}
}

func (fake *FakeVCenterManager) OperationsManagerReturnsOnCall(i int, result1 *guest.OperationsManager) {
	fake.operationsManagerMutex.Lock()
	defer fake.operationsManagerMutex.Unlock()
	fake.OperationsManagerStub = nil
	if fake.operationsManagerReturnsOnCall == nil {
		fake.operationsManagerReturnsOnCall = make(map[int]struct {
			result1 *guest.OperationsManager
		})
	}
	fake.operationsManagerReturnsOnCall[i] = struct {
		result1 *guest.OperationsManager
	}{result1}
}

func (fake *FakeVCenterManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	fake.guestManagerMutex.RLock()
	defer fake.guestManagerMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.operationsManagerMutex.RLock()
	defer fake.operationsManagerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVCenterManager) recordInvocation(key string, args []interface{}) {
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

var _ commandparser.VCenterManager = new(FakeVCenterManager)
