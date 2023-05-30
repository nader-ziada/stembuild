// Code generated by counterfeiter. DO NOT EDIT.
package packagersfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/stembuild/package_stemcell/packagers"
)

type FakeSource struct {
	ArtifactReaderStub        func() (io.Reader, error)
	artifactReaderMutex       sync.RWMutex
	artifactReaderArgsForCall []struct {
	}
	artifactReaderReturns struct {
		result1 io.Reader
		result2 error
	}
	artifactReaderReturnsOnCall map[int]struct {
		result1 io.Reader
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSource) ArtifactReader() (io.Reader, error) {
	fake.artifactReaderMutex.Lock()
	ret, specificReturn := fake.artifactReaderReturnsOnCall[len(fake.artifactReaderArgsForCall)]
	fake.artifactReaderArgsForCall = append(fake.artifactReaderArgsForCall, struct {
	}{})
	stub := fake.ArtifactReaderStub
	fakeReturns := fake.artifactReaderReturns
	fake.recordInvocation("ArtifactReader", []interface{}{})
	fake.artifactReaderMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSource) ArtifactReaderCallCount() int {
	fake.artifactReaderMutex.RLock()
	defer fake.artifactReaderMutex.RUnlock()
	return len(fake.artifactReaderArgsForCall)
}

func (fake *FakeSource) ArtifactReaderCalls(stub func() (io.Reader, error)) {
	fake.artifactReaderMutex.Lock()
	defer fake.artifactReaderMutex.Unlock()
	fake.ArtifactReaderStub = stub
}

func (fake *FakeSource) ArtifactReaderReturns(result1 io.Reader, result2 error) {
	fake.artifactReaderMutex.Lock()
	defer fake.artifactReaderMutex.Unlock()
	fake.ArtifactReaderStub = nil
	fake.artifactReaderReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ArtifactReaderReturnsOnCall(i int, result1 io.Reader, result2 error) {
	fake.artifactReaderMutex.Lock()
	defer fake.artifactReaderMutex.Unlock()
	fake.ArtifactReaderStub = nil
	if fake.artifactReaderReturnsOnCall == nil {
		fake.artifactReaderReturnsOnCall = make(map[int]struct {
			result1 io.Reader
			result2 error
		})
	}
	fake.artifactReaderReturnsOnCall[i] = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.artifactReaderMutex.RLock()
	defer fake.artifactReaderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSource) recordInvocation(key string, args []interface{}) {
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

var _ packagers.Source = new(FakeSource)
