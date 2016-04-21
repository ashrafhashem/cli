// This file was generated by counterfeiter
package spacequotasfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/spacequotas"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeSpaceQuotaRepository struct {
	FindByNameStub        func(name string) (quota models.SpaceQuota, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	FindByOrgStub        func(guid string) (quota []models.SpaceQuota, apiErr error)
	findByOrgMutex       sync.RWMutex
	findByOrgArgsForCall []struct {
		guid string
	}
	findByOrgReturns struct {
		result1 []models.SpaceQuota
		result2 error
	}
	FindByGuidStub        func(guid string) (quota models.SpaceQuota, apiErr error)
	findByGuidMutex       sync.RWMutex
	findByGuidArgsForCall []struct {
		guid string
	}
	findByGuidReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	FindByNameAndOrgGuidStub        func(spaceQuotaName string, orgGuid string) (quota models.SpaceQuota, apiErr error)
	findByNameAndOrgGuidMutex       sync.RWMutex
	findByNameAndOrgGuidArgsForCall []struct {
		spaceQuotaName string
		orgGuid        string
	}
	findByNameAndOrgGuidReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	AssociateSpaceWithQuotaStub        func(spaceGuid string, quotaGuid string) error
	associateSpaceWithQuotaMutex       sync.RWMutex
	associateSpaceWithQuotaArgsForCall []struct {
		spaceGuid string
		quotaGuid string
	}
	associateSpaceWithQuotaReturns struct {
		result1 error
	}
	UnassignQuotaFromSpaceStub        func(spaceGuid string, quotaGuid string) error
	unassignQuotaFromSpaceMutex       sync.RWMutex
	unassignQuotaFromSpaceArgsForCall []struct {
		spaceGuid string
		quotaGuid string
	}
	unassignQuotaFromSpaceReturns struct {
		result1 error
	}
	CreateStub        func(quota models.SpaceQuota) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		quota models.SpaceQuota
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(quota models.SpaceQuota) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		quota models.SpaceQuota
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(quotaGuid string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		quotaGuid string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeSpaceQuotaRepository) FindByName(name string) (quota models.SpaceQuota, apiErr error) {
	fake.findByNameMutex.Lock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeSpaceQuotaRepository) FindByNameReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByOrg(guid string) (quota []models.SpaceQuota, apiErr error) {
	fake.findByOrgMutex.Lock()
	fake.findByOrgArgsForCall = append(fake.findByOrgArgsForCall, struct {
		guid string
	}{guid})
	fake.findByOrgMutex.Unlock()
	if fake.FindByOrgStub != nil {
		return fake.FindByOrgStub(guid)
	} else {
		return fake.findByOrgReturns.result1, fake.findByOrgReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByOrgCallCount() int {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return len(fake.findByOrgArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByOrgArgsForCall(i int) string {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return fake.findByOrgArgsForCall[i].guid
}

func (fake *FakeSpaceQuotaRepository) FindByOrgReturns(result1 []models.SpaceQuota, result2 error) {
	fake.FindByOrgStub = nil
	fake.findByOrgReturns = struct {
		result1 []models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByGuid(guid string) (quota models.SpaceQuota, apiErr error) {
	fake.findByGuidMutex.Lock()
	fake.findByGuidArgsForCall = append(fake.findByGuidArgsForCall, struct {
		guid string
	}{guid})
	fake.findByGuidMutex.Unlock()
	if fake.FindByGuidStub != nil {
		return fake.FindByGuidStub(guid)
	} else {
		return fake.findByGuidReturns.result1, fake.findByGuidReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByGuidCallCount() int {
	fake.findByGuidMutex.RLock()
	defer fake.findByGuidMutex.RUnlock()
	return len(fake.findByGuidArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByGuidArgsForCall(i int) string {
	fake.findByGuidMutex.RLock()
	defer fake.findByGuidMutex.RUnlock()
	return fake.findByGuidArgsForCall[i].guid
}

func (fake *FakeSpaceQuotaRepository) FindByGuidReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByGuidStub = nil
	fake.findByGuidReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGuid(spaceQuotaName string, orgGuid string) (quota models.SpaceQuota, apiErr error) {
	fake.findByNameAndOrgGuidMutex.Lock()
	fake.findByNameAndOrgGuidArgsForCall = append(fake.findByNameAndOrgGuidArgsForCall, struct {
		spaceQuotaName string
		orgGuid        string
	}{spaceQuotaName, orgGuid})
	fake.findByNameAndOrgGuidMutex.Unlock()
	if fake.FindByNameAndOrgGuidStub != nil {
		return fake.FindByNameAndOrgGuidStub(spaceQuotaName, orgGuid)
	} else {
		return fake.findByNameAndOrgGuidReturns.result1, fake.findByNameAndOrgGuidReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGuidCallCount() int {
	fake.findByNameAndOrgGuidMutex.RLock()
	defer fake.findByNameAndOrgGuidMutex.RUnlock()
	return len(fake.findByNameAndOrgGuidArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGuidArgsForCall(i int) (string, string) {
	fake.findByNameAndOrgGuidMutex.RLock()
	defer fake.findByNameAndOrgGuidMutex.RUnlock()
	return fake.findByNameAndOrgGuidArgsForCall[i].spaceQuotaName, fake.findByNameAndOrgGuidArgsForCall[i].orgGuid
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGuidReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByNameAndOrgGuidStub = nil
	fake.findByNameAndOrgGuidReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuota(spaceGuid string, quotaGuid string) error {
	fake.associateSpaceWithQuotaMutex.Lock()
	fake.associateSpaceWithQuotaArgsForCall = append(fake.associateSpaceWithQuotaArgsForCall, struct {
		spaceGuid string
		quotaGuid string
	}{spaceGuid, quotaGuid})
	fake.associateSpaceWithQuotaMutex.Unlock()
	if fake.AssociateSpaceWithQuotaStub != nil {
		return fake.AssociateSpaceWithQuotaStub(spaceGuid, quotaGuid)
	} else {
		return fake.associateSpaceWithQuotaReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaCallCount() int {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return len(fake.associateSpaceWithQuotaArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaArgsForCall(i int) (string, string) {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return fake.associateSpaceWithQuotaArgsForCall[i].spaceGuid, fake.associateSpaceWithQuotaArgsForCall[i].quotaGuid
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaReturns(result1 error) {
	fake.AssociateSpaceWithQuotaStub = nil
	fake.associateSpaceWithQuotaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpace(spaceGuid string, quotaGuid string) error {
	fake.unassignQuotaFromSpaceMutex.Lock()
	fake.unassignQuotaFromSpaceArgsForCall = append(fake.unassignQuotaFromSpaceArgsForCall, struct {
		spaceGuid string
		quotaGuid string
	}{spaceGuid, quotaGuid})
	fake.unassignQuotaFromSpaceMutex.Unlock()
	if fake.UnassignQuotaFromSpaceStub != nil {
		return fake.UnassignQuotaFromSpaceStub(spaceGuid, quotaGuid)
	} else {
		return fake.unassignQuotaFromSpaceReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceCallCount() int {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return len(fake.unassignQuotaFromSpaceArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceArgsForCall(i int) (string, string) {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return fake.unassignQuotaFromSpaceArgsForCall[i].spaceGuid, fake.unassignQuotaFromSpaceArgsForCall[i].quotaGuid
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceReturns(result1 error) {
	fake.UnassignQuotaFromSpaceStub = nil
	fake.unassignQuotaFromSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Create(quota models.SpaceQuota) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		quota models.SpaceQuota
	}{quota})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(quota)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) CreateArgsForCall(i int) models.SpaceQuota {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].quota
}

func (fake *FakeSpaceQuotaRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Update(quota models.SpaceQuota) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		quota models.SpaceQuota
	}{quota})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(quota)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UpdateArgsForCall(i int) models.SpaceQuota {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].quota
}

func (fake *FakeSpaceQuotaRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Delete(quotaGuid string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		quotaGuid string
	}{quotaGuid})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(quotaGuid)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].quotaGuid
}

func (fake *FakeSpaceQuotaRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ spacequotas.SpaceQuotaRepository = new(FakeSpaceQuotaRepository)