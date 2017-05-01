// This file was generated by counterfeiter
package pluginactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakePluginMetadata struct {
	GetMetadataStub        func(pluginPath string) (configv3.Plugin, error)
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct {
		pluginPath string
	}
	getMetadataReturns struct {
		result1 configv3.Plugin
		result2 error
	}
	getMetadataReturnsOnCall map[int]struct {
		result1 configv3.Plugin
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginMetadata) GetMetadata(pluginPath string) (configv3.Plugin, error) {
	fake.getMetadataMutex.Lock()
	ret, specificReturn := fake.getMetadataReturnsOnCall[len(fake.getMetadataArgsForCall)]
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct {
		pluginPath string
	}{pluginPath})
	fake.recordInvocation("GetMetadata", []interface{}{pluginPath})
	fake.getMetadataMutex.Unlock()
	if fake.GetMetadataStub != nil {
		return fake.GetMetadataStub(pluginPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMetadataReturns.result1, fake.getMetadataReturns.result2
}

func (fake *FakePluginMetadata) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *FakePluginMetadata) GetMetadataArgsForCall(i int) string {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return fake.getMetadataArgsForCall[i].pluginPath
}

func (fake *FakePluginMetadata) GetMetadataReturns(result1 configv3.Plugin, result2 error) {
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 configv3.Plugin
		result2 error
	}{result1, result2}
}

func (fake *FakePluginMetadata) GetMetadataReturnsOnCall(i int, result1 configv3.Plugin, result2 error) {
	fake.GetMetadataStub = nil
	if fake.getMetadataReturnsOnCall == nil {
		fake.getMetadataReturnsOnCall = make(map[int]struct {
			result1 configv3.Plugin
			result2 error
		})
	}
	fake.getMetadataReturnsOnCall[i] = struct {
		result1 configv3.Plugin
		result2 error
	}{result1, result2}
}

func (fake *FakePluginMetadata) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePluginMetadata) recordInvocation(key string, args []interface{}) {
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

var _ pluginaction.PluginMetadata = new(FakePluginMetadata)