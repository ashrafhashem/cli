// This file was generated by counterfeiter
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeIsolationSegmentsActor struct {
	GetIsolationSegmentSummariesStub        func() ([]v3action.IsolationSegmentSummary, v3action.Warnings, error)
	getIsolationSegmentSummariesMutex       sync.RWMutex
	getIsolationSegmentSummariesArgsForCall []struct{}
	getIsolationSegmentSummariesReturns     struct {
		result1 []v3action.IsolationSegmentSummary
		result2 v3action.Warnings
		result3 error
	}
	getIsolationSegmentSummariesReturnsOnCall map[int]struct {
		result1 []v3action.IsolationSegmentSummary
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIsolationSegmentsActor) GetIsolationSegmentSummaries() ([]v3action.IsolationSegmentSummary, v3action.Warnings, error) {
	fake.getIsolationSegmentSummariesMutex.Lock()
	ret, specificReturn := fake.getIsolationSegmentSummariesReturnsOnCall[len(fake.getIsolationSegmentSummariesArgsForCall)]
	fake.getIsolationSegmentSummariesArgsForCall = append(fake.getIsolationSegmentSummariesArgsForCall, struct{}{})
	fake.recordInvocation("GetIsolationSegmentSummaries", []interface{}{})
	fake.getIsolationSegmentSummariesMutex.Unlock()
	if fake.GetIsolationSegmentSummariesStub != nil {
		return fake.GetIsolationSegmentSummariesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getIsolationSegmentSummariesReturns.result1, fake.getIsolationSegmentSummariesReturns.result2, fake.getIsolationSegmentSummariesReturns.result3
}

func (fake *FakeIsolationSegmentsActor) GetIsolationSegmentSummariesCallCount() int {
	fake.getIsolationSegmentSummariesMutex.RLock()
	defer fake.getIsolationSegmentSummariesMutex.RUnlock()
	return len(fake.getIsolationSegmentSummariesArgsForCall)
}

func (fake *FakeIsolationSegmentsActor) GetIsolationSegmentSummariesReturns(result1 []v3action.IsolationSegmentSummary, result2 v3action.Warnings, result3 error) {
	fake.GetIsolationSegmentSummariesStub = nil
	fake.getIsolationSegmentSummariesReturns = struct {
		result1 []v3action.IsolationSegmentSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIsolationSegmentsActor) GetIsolationSegmentSummariesReturnsOnCall(i int, result1 []v3action.IsolationSegmentSummary, result2 v3action.Warnings, result3 error) {
	fake.GetIsolationSegmentSummariesStub = nil
	if fake.getIsolationSegmentSummariesReturnsOnCall == nil {
		fake.getIsolationSegmentSummariesReturnsOnCall = make(map[int]struct {
			result1 []v3action.IsolationSegmentSummary
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getIsolationSegmentSummariesReturnsOnCall[i] = struct {
		result1 []v3action.IsolationSegmentSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIsolationSegmentsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIsolationSegmentSummariesMutex.RLock()
	defer fake.getIsolationSegmentSummariesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIsolationSegmentsActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.IsolationSegmentsActor = new(FakeIsolationSegmentsActor)
