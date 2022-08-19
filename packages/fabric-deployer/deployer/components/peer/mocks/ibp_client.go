// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/IBM-Blockchain/fabric-deployer/deployer/components/peer"
	"k8s.io/apimachinery/pkg/runtime"
)

type IBPOperatorClient struct {
	CreateCRStub        func(string, string, interface{}) error
	createCRMutex       sync.RWMutex
	createCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}
	createCRReturns struct {
		result1 error
	}
	createCRReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCRStub        func(string, string, string) error
	deleteCRMutex       sync.RWMutex
	deleteCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	deleteCRReturns struct {
		result1 error
	}
	deleteCRReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllCRStub        func(string, string, runtime.Object) error
	getAllCRMutex       sync.RWMutex
	getAllCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 runtime.Object
	}
	getAllCRReturns struct {
		result1 error
	}
	getAllCRReturnsOnCall map[int]struct {
		result1 error
	}
	GetCRStub        func(string, string, string, runtime.Object) error
	getCRMutex       sync.RWMutex
	getCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 runtime.Object
	}
	getCRReturns struct {
		result1 error
	}
	getCRReturnsOnCall map[int]struct {
		result1 error
	}
	PatchCRStub        func(string, string, string, []byte) error
	patchCRMutex       sync.RWMutex
	patchCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []byte
	}
	patchCRReturns struct {
		result1 error
	}
	patchCRReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateCRStub        func(string, string, string, []byte) error
	updateCRMutex       sync.RWMutex
	updateCRArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []byte
	}
	updateCRReturns struct {
		result1 error
	}
	updateCRReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IBPOperatorClient) CreateCR(arg1 string, arg2 string, arg3 interface{}) error {
	fake.createCRMutex.Lock()
	ret, specificReturn := fake.createCRReturnsOnCall[len(fake.createCRArgsForCall)]
	fake.createCRArgsForCall = append(fake.createCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}{arg1, arg2, arg3})
	stub := fake.CreateCRStub
	fakeReturns := fake.createCRReturns
	fake.recordInvocation("CreateCR", []interface{}{arg1, arg2, arg3})
	fake.createCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) CreateCRCallCount() int {
	fake.createCRMutex.RLock()
	defer fake.createCRMutex.RUnlock()
	return len(fake.createCRArgsForCall)
}

func (fake *IBPOperatorClient) CreateCRCalls(stub func(string, string, interface{}) error) {
	fake.createCRMutex.Lock()
	defer fake.createCRMutex.Unlock()
	fake.CreateCRStub = stub
}

func (fake *IBPOperatorClient) CreateCRArgsForCall(i int) (string, string, interface{}) {
	fake.createCRMutex.RLock()
	defer fake.createCRMutex.RUnlock()
	argsForCall := fake.createCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IBPOperatorClient) CreateCRReturns(result1 error) {
	fake.createCRMutex.Lock()
	defer fake.createCRMutex.Unlock()
	fake.CreateCRStub = nil
	fake.createCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) CreateCRReturnsOnCall(i int, result1 error) {
	fake.createCRMutex.Lock()
	defer fake.createCRMutex.Unlock()
	fake.CreateCRStub = nil
	if fake.createCRReturnsOnCall == nil {
		fake.createCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) DeleteCR(arg1 string, arg2 string, arg3 string) error {
	fake.deleteCRMutex.Lock()
	ret, specificReturn := fake.deleteCRReturnsOnCall[len(fake.deleteCRArgsForCall)]
	fake.deleteCRArgsForCall = append(fake.deleteCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeleteCRStub
	fakeReturns := fake.deleteCRReturns
	fake.recordInvocation("DeleteCR", []interface{}{arg1, arg2, arg3})
	fake.deleteCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) DeleteCRCallCount() int {
	fake.deleteCRMutex.RLock()
	defer fake.deleteCRMutex.RUnlock()
	return len(fake.deleteCRArgsForCall)
}

func (fake *IBPOperatorClient) DeleteCRCalls(stub func(string, string, string) error) {
	fake.deleteCRMutex.Lock()
	defer fake.deleteCRMutex.Unlock()
	fake.DeleteCRStub = stub
}

func (fake *IBPOperatorClient) DeleteCRArgsForCall(i int) (string, string, string) {
	fake.deleteCRMutex.RLock()
	defer fake.deleteCRMutex.RUnlock()
	argsForCall := fake.deleteCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IBPOperatorClient) DeleteCRReturns(result1 error) {
	fake.deleteCRMutex.Lock()
	defer fake.deleteCRMutex.Unlock()
	fake.DeleteCRStub = nil
	fake.deleteCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) DeleteCRReturnsOnCall(i int, result1 error) {
	fake.deleteCRMutex.Lock()
	defer fake.deleteCRMutex.Unlock()
	fake.DeleteCRStub = nil
	if fake.deleteCRReturnsOnCall == nil {
		fake.deleteCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) GetAllCR(arg1 string, arg2 string, arg3 runtime.Object) error {
	fake.getAllCRMutex.Lock()
	ret, specificReturn := fake.getAllCRReturnsOnCall[len(fake.getAllCRArgsForCall)]
	fake.getAllCRArgsForCall = append(fake.getAllCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 runtime.Object
	}{arg1, arg2, arg3})
	stub := fake.GetAllCRStub
	fakeReturns := fake.getAllCRReturns
	fake.recordInvocation("GetAllCR", []interface{}{arg1, arg2, arg3})
	fake.getAllCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) GetAllCRCallCount() int {
	fake.getAllCRMutex.RLock()
	defer fake.getAllCRMutex.RUnlock()
	return len(fake.getAllCRArgsForCall)
}

func (fake *IBPOperatorClient) GetAllCRCalls(stub func(string, string, runtime.Object) error) {
	fake.getAllCRMutex.Lock()
	defer fake.getAllCRMutex.Unlock()
	fake.GetAllCRStub = stub
}

func (fake *IBPOperatorClient) GetAllCRArgsForCall(i int) (string, string, runtime.Object) {
	fake.getAllCRMutex.RLock()
	defer fake.getAllCRMutex.RUnlock()
	argsForCall := fake.getAllCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IBPOperatorClient) GetAllCRReturns(result1 error) {
	fake.getAllCRMutex.Lock()
	defer fake.getAllCRMutex.Unlock()
	fake.GetAllCRStub = nil
	fake.getAllCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) GetAllCRReturnsOnCall(i int, result1 error) {
	fake.getAllCRMutex.Lock()
	defer fake.getAllCRMutex.Unlock()
	fake.GetAllCRStub = nil
	if fake.getAllCRReturnsOnCall == nil {
		fake.getAllCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getAllCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) GetCR(arg1 string, arg2 string, arg3 string, arg4 runtime.Object) error {
	fake.getCRMutex.Lock()
	ret, specificReturn := fake.getCRReturnsOnCall[len(fake.getCRArgsForCall)]
	fake.getCRArgsForCall = append(fake.getCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 runtime.Object
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetCRStub
	fakeReturns := fake.getCRReturns
	fake.recordInvocation("GetCR", []interface{}{arg1, arg2, arg3, arg4})
	fake.getCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) GetCRCallCount() int {
	fake.getCRMutex.RLock()
	defer fake.getCRMutex.RUnlock()
	return len(fake.getCRArgsForCall)
}

func (fake *IBPOperatorClient) GetCRCalls(stub func(string, string, string, runtime.Object) error) {
	fake.getCRMutex.Lock()
	defer fake.getCRMutex.Unlock()
	fake.GetCRStub = stub
}

func (fake *IBPOperatorClient) GetCRArgsForCall(i int) (string, string, string, runtime.Object) {
	fake.getCRMutex.RLock()
	defer fake.getCRMutex.RUnlock()
	argsForCall := fake.getCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *IBPOperatorClient) GetCRReturns(result1 error) {
	fake.getCRMutex.Lock()
	defer fake.getCRMutex.Unlock()
	fake.GetCRStub = nil
	fake.getCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) GetCRReturnsOnCall(i int, result1 error) {
	fake.getCRMutex.Lock()
	defer fake.getCRMutex.Unlock()
	fake.GetCRStub = nil
	if fake.getCRReturnsOnCall == nil {
		fake.getCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) PatchCR(arg1 string, arg2 string, arg3 string, arg4 []byte) error {
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.patchCRMutex.Lock()
	ret, specificReturn := fake.patchCRReturnsOnCall[len(fake.patchCRArgsForCall)]
	fake.patchCRArgsForCall = append(fake.patchCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []byte
	}{arg1, arg2, arg3, arg4Copy})
	stub := fake.PatchCRStub
	fakeReturns := fake.patchCRReturns
	fake.recordInvocation("PatchCR", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.patchCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) PatchCRCallCount() int {
	fake.patchCRMutex.RLock()
	defer fake.patchCRMutex.RUnlock()
	return len(fake.patchCRArgsForCall)
}

func (fake *IBPOperatorClient) PatchCRCalls(stub func(string, string, string, []byte) error) {
	fake.patchCRMutex.Lock()
	defer fake.patchCRMutex.Unlock()
	fake.PatchCRStub = stub
}

func (fake *IBPOperatorClient) PatchCRArgsForCall(i int) (string, string, string, []byte) {
	fake.patchCRMutex.RLock()
	defer fake.patchCRMutex.RUnlock()
	argsForCall := fake.patchCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *IBPOperatorClient) PatchCRReturns(result1 error) {
	fake.patchCRMutex.Lock()
	defer fake.patchCRMutex.Unlock()
	fake.PatchCRStub = nil
	fake.patchCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) PatchCRReturnsOnCall(i int, result1 error) {
	fake.patchCRMutex.Lock()
	defer fake.patchCRMutex.Unlock()
	fake.PatchCRStub = nil
	if fake.patchCRReturnsOnCall == nil {
		fake.patchCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.patchCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) UpdateCR(arg1 string, arg2 string, arg3 string, arg4 []byte) error {
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.updateCRMutex.Lock()
	ret, specificReturn := fake.updateCRReturnsOnCall[len(fake.updateCRArgsForCall)]
	fake.updateCRArgsForCall = append(fake.updateCRArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 []byte
	}{arg1, arg2, arg3, arg4Copy})
	stub := fake.UpdateCRStub
	fakeReturns := fake.updateCRReturns
	fake.recordInvocation("UpdateCR", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.updateCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IBPOperatorClient) UpdateCRCallCount() int {
	fake.updateCRMutex.RLock()
	defer fake.updateCRMutex.RUnlock()
	return len(fake.updateCRArgsForCall)
}

func (fake *IBPOperatorClient) UpdateCRCalls(stub func(string, string, string, []byte) error) {
	fake.updateCRMutex.Lock()
	defer fake.updateCRMutex.Unlock()
	fake.UpdateCRStub = stub
}

func (fake *IBPOperatorClient) UpdateCRArgsForCall(i int) (string, string, string, []byte) {
	fake.updateCRMutex.RLock()
	defer fake.updateCRMutex.RUnlock()
	argsForCall := fake.updateCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *IBPOperatorClient) UpdateCRReturns(result1 error) {
	fake.updateCRMutex.Lock()
	defer fake.updateCRMutex.Unlock()
	fake.UpdateCRStub = nil
	fake.updateCRReturns = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) UpdateCRReturnsOnCall(i int, result1 error) {
	fake.updateCRMutex.Lock()
	defer fake.updateCRMutex.Unlock()
	fake.UpdateCRStub = nil
	if fake.updateCRReturnsOnCall == nil {
		fake.updateCRReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateCRReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IBPOperatorClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCRMutex.RLock()
	defer fake.createCRMutex.RUnlock()
	fake.deleteCRMutex.RLock()
	defer fake.deleteCRMutex.RUnlock()
	fake.getAllCRMutex.RLock()
	defer fake.getAllCRMutex.RUnlock()
	fake.getCRMutex.RLock()
	defer fake.getCRMutex.RUnlock()
	fake.patchCRMutex.RLock()
	defer fake.patchCRMutex.RUnlock()
	fake.updateCRMutex.RLock()
	defer fake.updateCRMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IBPOperatorClient) recordInvocation(key string, args []interface{}) {
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

var _ peer.IBPOperatorClient = new(IBPOperatorClient)