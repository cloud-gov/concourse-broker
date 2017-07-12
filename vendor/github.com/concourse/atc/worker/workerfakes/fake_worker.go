// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
	"github.com/cppforlife/go-semi-semantic/version"
)

type FakeWorker struct {
	FindOrCreateContainerStub        func(lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, db.ResourceUser, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes) (worker.Container, error)
	findOrCreateContainerMutex       sync.RWMutex
	findOrCreateContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 db.ResourceUser
		arg5 db.ContainerOwner
		arg6 db.ContainerMetadata
		arg7 worker.ContainerSpec
		arg8 creds.VersionedResourceTypes
	}
	findOrCreateContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	findOrCreateContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 error
	}
	FindContainerByHandleStub        func(lager.Logger, int, string) (worker.Container, bool, error)
	findContainerByHandleMutex       sync.RWMutex
	findContainerByHandleArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findContainerByHandleReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	findContainerByHandleReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	FindResourceTypeByPathStub        func(path string) (atc.WorkerResourceType, bool)
	findResourceTypeByPathMutex       sync.RWMutex
	findResourceTypeByPathArgsForCall []struct {
		path string
	}
	findResourceTypeByPathReturns struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	findResourceTypeByPathReturnsOnCall map[int]struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	lookupVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	SatisfyingStub        func(lager.Logger, worker.WorkerSpec, creds.VersionedResourceTypes) (worker.Worker, error)
	satisfyingMutex       sync.RWMutex
	satisfyingArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
		arg3 creds.VersionedResourceTypes
	}
	satisfyingReturns struct {
		result1 worker.Worker
		result2 error
	}
	satisfyingReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	AllSatisfyingStub        func(lager.Logger, worker.WorkerSpec, creds.VersionedResourceTypes) ([]worker.Worker, error)
	allSatisfyingMutex       sync.RWMutex
	allSatisfyingArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
		arg3 creds.VersionedResourceTypes
	}
	allSatisfyingReturns struct {
		result1 []worker.Worker
		result2 error
	}
	allSatisfyingReturnsOnCall map[int]struct {
		result1 []worker.Worker
		result2 error
	}
	RunningWorkersStub        func(lager.Logger) ([]worker.Worker, error)
	runningWorkersMutex       sync.RWMutex
	runningWorkersArgsForCall []struct {
		arg1 lager.Logger
	}
	runningWorkersReturns struct {
		result1 []worker.Worker
		result2 error
	}
	runningWorkersReturnsOnCall map[int]struct {
		result1 []worker.Worker
		result2 error
	}
	ActiveContainersStub        func() int
	activeContainersMutex       sync.RWMutex
	activeContainersArgsForCall []struct{}
	activeContainersReturns     struct {
		result1 int
	}
	activeContainersReturnsOnCall map[int]struct {
		result1 int
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
		result1 string
	}
	descriptionReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ResourceTypesStub        func() []atc.WorkerResourceType
	resourceTypesMutex       sync.RWMutex
	resourceTypesArgsForCall []struct{}
	resourceTypesReturns     struct {
		result1 []atc.WorkerResourceType
	}
	resourceTypesReturnsOnCall map[int]struct {
		result1 []atc.WorkerResourceType
	}
	TagsStub        func() atc.Tags
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct{}
	tagsReturns     struct {
		result1 atc.Tags
	}
	tagsReturnsOnCall map[int]struct {
		result1 atc.Tags
	}
	UptimeStub        func() time.Duration
	uptimeMutex       sync.RWMutex
	uptimeArgsForCall []struct{}
	uptimeReturns     struct {
		result1 time.Duration
	}
	uptimeReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	IsOwnedByTeamStub        func() bool
	isOwnedByTeamMutex       sync.RWMutex
	isOwnedByTeamArgsForCall []struct{}
	isOwnedByTeamReturns     struct {
		result1 bool
	}
	isOwnedByTeamReturnsOnCall map[int]struct {
		result1 bool
	}
	IsVersionCompatibleStub        func(lager.Logger, *version.Version) bool
	isVersionCompatibleMutex       sync.RWMutex
	isVersionCompatibleArgsForCall []struct {
		arg1 lager.Logger
		arg2 *version.Version
	}
	isVersionCompatibleReturns struct {
		result1 bool
	}
	isVersionCompatibleReturnsOnCall map[int]struct {
		result1 bool
	}
	FindVolumeForResourceCacheStub        func(logger lager.Logger, resourceCache *db.UsedResourceCache) (worker.Volume, bool, error)
	findVolumeForResourceCacheMutex       sync.RWMutex
	findVolumeForResourceCacheArgsForCall []struct {
		logger        lager.Logger
		resourceCache *db.UsedResourceCache
	}
	findVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeForResourceCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindVolumeForTaskCacheStub        func(lager.Logger, int, int, string, string) (worker.Volume, bool, error)
	findVolumeForTaskCacheMutex       sync.RWMutex
	findVolumeForTaskCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 string
		arg5 string
	}
	findVolumeForTaskCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeForTaskCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) FindOrCreateContainer(arg1 lager.Logger, arg2 <-chan os.Signal, arg3 worker.ImageFetchingDelegate, arg4 db.ResourceUser, arg5 db.ContainerOwner, arg6 db.ContainerMetadata, arg7 worker.ContainerSpec, arg8 creds.VersionedResourceTypes) (worker.Container, error) {
	fake.findOrCreateContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateContainerReturnsOnCall[len(fake.findOrCreateContainerArgsForCall)]
	fake.findOrCreateContainerArgsForCall = append(fake.findOrCreateContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 db.ResourceUser
		arg5 db.ContainerOwner
		arg6 db.ContainerMetadata
		arg7 worker.ContainerSpec
		arg8 creds.VersionedResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("FindOrCreateContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.findOrCreateContainerMutex.Unlock()
	if fake.FindOrCreateContainerStub != nil {
		return fake.FindOrCreateContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateContainerReturns.result1, fake.findOrCreateContainerReturns.result2
}

func (fake *FakeWorker) FindOrCreateContainerCallCount() int {
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	return len(fake.findOrCreateContainerArgsForCall)
}

func (fake *FakeWorker) FindOrCreateContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, db.ResourceUser, db.ContainerOwner, db.ContainerMetadata, worker.ContainerSpec, creds.VersionedResourceTypes) {
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	return fake.findOrCreateContainerArgsForCall[i].arg1, fake.findOrCreateContainerArgsForCall[i].arg2, fake.findOrCreateContainerArgsForCall[i].arg3, fake.findOrCreateContainerArgsForCall[i].arg4, fake.findOrCreateContainerArgsForCall[i].arg5, fake.findOrCreateContainerArgsForCall[i].arg6, fake.findOrCreateContainerArgsForCall[i].arg7, fake.findOrCreateContainerArgsForCall[i].arg8
}

func (fake *FakeWorker) FindOrCreateContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateContainerStub = nil
	fake.findOrCreateContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindOrCreateContainerReturnsOnCall(i int, result1 worker.Container, result2 error) {
	fake.FindOrCreateContainerStub = nil
	if fake.findOrCreateContainerReturnsOnCall == nil {
		fake.findOrCreateContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 error
		})
	}
	fake.findOrCreateContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindContainerByHandle(arg1 lager.Logger, arg2 int, arg3 string) (worker.Container, bool, error) {
	fake.findContainerByHandleMutex.Lock()
	ret, specificReturn := fake.findContainerByHandleReturnsOnCall[len(fake.findContainerByHandleArgsForCall)]
	fake.findContainerByHandleArgsForCall = append(fake.findContainerByHandleArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindContainerByHandle", []interface{}{arg1, arg2, arg3})
	fake.findContainerByHandleMutex.Unlock()
	if fake.FindContainerByHandleStub != nil {
		return fake.FindContainerByHandleStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findContainerByHandleReturns.result1, fake.findContainerByHandleReturns.result2, fake.findContainerByHandleReturns.result3
}

func (fake *FakeWorker) FindContainerByHandleCallCount() int {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return len(fake.findContainerByHandleArgsForCall)
}

func (fake *FakeWorker) FindContainerByHandleArgsForCall(i int) (lager.Logger, int, string) {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return fake.findContainerByHandleArgsForCall[i].arg1, fake.findContainerByHandleArgsForCall[i].arg2, fake.findContainerByHandleArgsForCall[i].arg3
}

func (fake *FakeWorker) FindContainerByHandleReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerByHandleStub = nil
	fake.findContainerByHandleReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) FindContainerByHandleReturnsOnCall(i int, result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerByHandleStub = nil
	if fake.findContainerByHandleReturnsOnCall == nil {
		fake.findContainerByHandleReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 bool
			result3 error
		})
	}
	fake.findContainerByHandleReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) FindResourceTypeByPath(path string) (atc.WorkerResourceType, bool) {
	fake.findResourceTypeByPathMutex.Lock()
	ret, specificReturn := fake.findResourceTypeByPathReturnsOnCall[len(fake.findResourceTypeByPathArgsForCall)]
	fake.findResourceTypeByPathArgsForCall = append(fake.findResourceTypeByPathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("FindResourceTypeByPath", []interface{}{path})
	fake.findResourceTypeByPathMutex.Unlock()
	if fake.FindResourceTypeByPathStub != nil {
		return fake.FindResourceTypeByPathStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findResourceTypeByPathReturns.result1, fake.findResourceTypeByPathReturns.result2
}

func (fake *FakeWorker) FindResourceTypeByPathCallCount() int {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return len(fake.findResourceTypeByPathArgsForCall)
}

func (fake *FakeWorker) FindResourceTypeByPathArgsForCall(i int) string {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return fake.findResourceTypeByPathArgsForCall[i].path
}

func (fake *FakeWorker) FindResourceTypeByPathReturns(result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	fake.findResourceTypeByPathReturns = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeWorker) FindResourceTypeByPathReturnsOnCall(i int, result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	if fake.findResourceTypeByPathReturnsOnCall == nil {
		fake.findResourceTypeByPathReturnsOnCall = make(map[int]struct {
			result1 atc.WorkerResourceType
			result2 bool
		})
	}
	fake.findResourceTypeByPathReturnsOnCall[i] = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeWorker) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	ret, specificReturn := fake.lookupVolumeReturnsOnCall[len(fake.lookupVolumeArgsForCall)]
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
}

func (fake *FakeWorker) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeWorker) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeWorker) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) LookupVolumeReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	if fake.lookupVolumeReturnsOnCall == nil {
		fake.lookupVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.lookupVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) Satisfying(arg1 lager.Logger, arg2 worker.WorkerSpec, arg3 creds.VersionedResourceTypes) (worker.Worker, error) {
	fake.satisfyingMutex.Lock()
	ret, specificReturn := fake.satisfyingReturnsOnCall[len(fake.satisfyingArgsForCall)]
	fake.satisfyingArgsForCall = append(fake.satisfyingArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
		arg3 creds.VersionedResourceTypes
	}{arg1, arg2, arg3})
	fake.recordInvocation("Satisfying", []interface{}{arg1, arg2, arg3})
	fake.satisfyingMutex.Unlock()
	if fake.SatisfyingStub != nil {
		return fake.SatisfyingStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.satisfyingReturns.result1, fake.satisfyingReturns.result2
}

func (fake *FakeWorker) SatisfyingCallCount() int {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return len(fake.satisfyingArgsForCall)
}

func (fake *FakeWorker) SatisfyingArgsForCall(i int) (lager.Logger, worker.WorkerSpec, creds.VersionedResourceTypes) {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return fake.satisfyingArgsForCall[i].arg1, fake.satisfyingArgsForCall[i].arg2, fake.satisfyingArgsForCall[i].arg3
}

func (fake *FakeWorker) SatisfyingReturns(result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	fake.satisfyingReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) SatisfyingReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	if fake.satisfyingReturnsOnCall == nil {
		fake.satisfyingReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.satisfyingReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) AllSatisfying(arg1 lager.Logger, arg2 worker.WorkerSpec, arg3 creds.VersionedResourceTypes) ([]worker.Worker, error) {
	fake.allSatisfyingMutex.Lock()
	ret, specificReturn := fake.allSatisfyingReturnsOnCall[len(fake.allSatisfyingArgsForCall)]
	fake.allSatisfyingArgsForCall = append(fake.allSatisfyingArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
		arg3 creds.VersionedResourceTypes
	}{arg1, arg2, arg3})
	fake.recordInvocation("AllSatisfying", []interface{}{arg1, arg2, arg3})
	fake.allSatisfyingMutex.Unlock()
	if fake.AllSatisfyingStub != nil {
		return fake.AllSatisfyingStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allSatisfyingReturns.result1, fake.allSatisfyingReturns.result2
}

func (fake *FakeWorker) AllSatisfyingCallCount() int {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return len(fake.allSatisfyingArgsForCall)
}

func (fake *FakeWorker) AllSatisfyingArgsForCall(i int) (lager.Logger, worker.WorkerSpec, creds.VersionedResourceTypes) {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return fake.allSatisfyingArgsForCall[i].arg1, fake.allSatisfyingArgsForCall[i].arg2, fake.allSatisfyingArgsForCall[i].arg3
}

func (fake *FakeWorker) AllSatisfyingReturns(result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	fake.allSatisfyingReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) AllSatisfyingReturnsOnCall(i int, result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	if fake.allSatisfyingReturnsOnCall == nil {
		fake.allSatisfyingReturnsOnCall = make(map[int]struct {
			result1 []worker.Worker
			result2 error
		})
	}
	fake.allSatisfyingReturnsOnCall[i] = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) RunningWorkers(arg1 lager.Logger) ([]worker.Worker, error) {
	fake.runningWorkersMutex.Lock()
	ret, specificReturn := fake.runningWorkersReturnsOnCall[len(fake.runningWorkersArgsForCall)]
	fake.runningWorkersArgsForCall = append(fake.runningWorkersArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("RunningWorkers", []interface{}{arg1})
	fake.runningWorkersMutex.Unlock()
	if fake.RunningWorkersStub != nil {
		return fake.RunningWorkersStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runningWorkersReturns.result1, fake.runningWorkersReturns.result2
}

func (fake *FakeWorker) RunningWorkersCallCount() int {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return len(fake.runningWorkersArgsForCall)
}

func (fake *FakeWorker) RunningWorkersArgsForCall(i int) lager.Logger {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return fake.runningWorkersArgsForCall[i].arg1
}

func (fake *FakeWorker) RunningWorkersReturns(result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	fake.runningWorkersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) RunningWorkersReturnsOnCall(i int, result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	if fake.runningWorkersReturnsOnCall == nil {
		fake.runningWorkersReturnsOnCall = make(map[int]struct {
			result1 []worker.Worker
			result2 error
		})
	}
	fake.runningWorkersReturnsOnCall[i] = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ActiveContainers() int {
	fake.activeContainersMutex.Lock()
	ret, specificReturn := fake.activeContainersReturnsOnCall[len(fake.activeContainersArgsForCall)]
	fake.activeContainersArgsForCall = append(fake.activeContainersArgsForCall, struct{}{})
	fake.recordInvocation("ActiveContainers", []interface{}{})
	fake.activeContainersMutex.Unlock()
	if fake.ActiveContainersStub != nil {
		return fake.ActiveContainersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.activeContainersReturns.result1
}

func (fake *FakeWorker) ActiveContainersCallCount() int {
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	return len(fake.activeContainersArgsForCall)
}

func (fake *FakeWorker) ActiveContainersReturns(result1 int) {
	fake.ActiveContainersStub = nil
	fake.activeContainersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) ActiveContainersReturnsOnCall(i int, result1 int) {
	fake.ActiveContainersStub = nil
	if fake.activeContainersReturnsOnCall == nil {
		fake.activeContainersReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.activeContainersReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) Description() string {
	fake.descriptionMutex.Lock()
	ret, specificReturn := fake.descriptionReturnsOnCall[len(fake.descriptionArgsForCall)]
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.descriptionReturns.result1
}

func (fake *FakeWorker) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeWorker) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) DescriptionReturnsOnCall(i int, result1 string) {
	fake.DescriptionStub = nil
	if fake.descriptionReturnsOnCall == nil {
		fake.descriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.descriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeWorker) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorker) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) ResourceTypes() []atc.WorkerResourceType {
	fake.resourceTypesMutex.Lock()
	ret, specificReturn := fake.resourceTypesReturnsOnCall[len(fake.resourceTypesArgsForCall)]
	fake.resourceTypesArgsForCall = append(fake.resourceTypesArgsForCall, struct{}{})
	fake.recordInvocation("ResourceTypes", []interface{}{})
	fake.resourceTypesMutex.Unlock()
	if fake.ResourceTypesStub != nil {
		return fake.ResourceTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceTypesReturns.result1
}

func (fake *FakeWorker) ResourceTypesCallCount() int {
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	return len(fake.resourceTypesArgsForCall)
}

func (fake *FakeWorker) ResourceTypesReturns(result1 []atc.WorkerResourceType) {
	fake.ResourceTypesStub = nil
	fake.resourceTypesReturns = struct {
		result1 []atc.WorkerResourceType
	}{result1}
}

func (fake *FakeWorker) ResourceTypesReturnsOnCall(i int, result1 []atc.WorkerResourceType) {
	fake.ResourceTypesStub = nil
	if fake.resourceTypesReturnsOnCall == nil {
		fake.resourceTypesReturnsOnCall = make(map[int]struct {
			result1 []atc.WorkerResourceType
		})
	}
	fake.resourceTypesReturnsOnCall[i] = struct {
		result1 []atc.WorkerResourceType
	}{result1}
}

func (fake *FakeWorker) Tags() atc.Tags {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct{}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tagsReturns.result1
}

func (fake *FakeWorker) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeWorker) TagsReturns(result1 atc.Tags) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeWorker) TagsReturnsOnCall(i int, result1 atc.Tags) {
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 atc.Tags
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeWorker) Uptime() time.Duration {
	fake.uptimeMutex.Lock()
	ret, specificReturn := fake.uptimeReturnsOnCall[len(fake.uptimeArgsForCall)]
	fake.uptimeArgsForCall = append(fake.uptimeArgsForCall, struct{}{})
	fake.recordInvocation("Uptime", []interface{}{})
	fake.uptimeMutex.Unlock()
	if fake.UptimeStub != nil {
		return fake.UptimeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uptimeReturns.result1
}

func (fake *FakeWorker) UptimeCallCount() int {
	fake.uptimeMutex.RLock()
	defer fake.uptimeMutex.RUnlock()
	return len(fake.uptimeArgsForCall)
}

func (fake *FakeWorker) UptimeReturns(result1 time.Duration) {
	fake.UptimeStub = nil
	fake.uptimeReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeWorker) UptimeReturnsOnCall(i int, result1 time.Duration) {
	fake.UptimeStub = nil
	if fake.uptimeReturnsOnCall == nil {
		fake.uptimeReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.uptimeReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeWorker) IsOwnedByTeam() bool {
	fake.isOwnedByTeamMutex.Lock()
	ret, specificReturn := fake.isOwnedByTeamReturnsOnCall[len(fake.isOwnedByTeamArgsForCall)]
	fake.isOwnedByTeamArgsForCall = append(fake.isOwnedByTeamArgsForCall, struct{}{})
	fake.recordInvocation("IsOwnedByTeam", []interface{}{})
	fake.isOwnedByTeamMutex.Unlock()
	if fake.IsOwnedByTeamStub != nil {
		return fake.IsOwnedByTeamStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isOwnedByTeamReturns.result1
}

func (fake *FakeWorker) IsOwnedByTeamCallCount() int {
	fake.isOwnedByTeamMutex.RLock()
	defer fake.isOwnedByTeamMutex.RUnlock()
	return len(fake.isOwnedByTeamArgsForCall)
}

func (fake *FakeWorker) IsOwnedByTeamReturns(result1 bool) {
	fake.IsOwnedByTeamStub = nil
	fake.isOwnedByTeamReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) IsOwnedByTeamReturnsOnCall(i int, result1 bool) {
	fake.IsOwnedByTeamStub = nil
	if fake.isOwnedByTeamReturnsOnCall == nil {
		fake.isOwnedByTeamReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isOwnedByTeamReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) IsVersionCompatible(arg1 lager.Logger, arg2 *version.Version) bool {
	fake.isVersionCompatibleMutex.Lock()
	ret, specificReturn := fake.isVersionCompatibleReturnsOnCall[len(fake.isVersionCompatibleArgsForCall)]
	fake.isVersionCompatibleArgsForCall = append(fake.isVersionCompatibleArgsForCall, struct {
		arg1 lager.Logger
		arg2 *version.Version
	}{arg1, arg2})
	fake.recordInvocation("IsVersionCompatible", []interface{}{arg1, arg2})
	fake.isVersionCompatibleMutex.Unlock()
	if fake.IsVersionCompatibleStub != nil {
		return fake.IsVersionCompatibleStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isVersionCompatibleReturns.result1
}

func (fake *FakeWorker) IsVersionCompatibleCallCount() int {
	fake.isVersionCompatibleMutex.RLock()
	defer fake.isVersionCompatibleMutex.RUnlock()
	return len(fake.isVersionCompatibleArgsForCall)
}

func (fake *FakeWorker) IsVersionCompatibleArgsForCall(i int) (lager.Logger, *version.Version) {
	fake.isVersionCompatibleMutex.RLock()
	defer fake.isVersionCompatibleMutex.RUnlock()
	return fake.isVersionCompatibleArgsForCall[i].arg1, fake.isVersionCompatibleArgsForCall[i].arg2
}

func (fake *FakeWorker) IsVersionCompatibleReturns(result1 bool) {
	fake.IsVersionCompatibleStub = nil
	fake.isVersionCompatibleReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) IsVersionCompatibleReturnsOnCall(i int, result1 bool) {
	fake.IsVersionCompatibleStub = nil
	if fake.isVersionCompatibleReturnsOnCall == nil {
		fake.isVersionCompatibleReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isVersionCompatibleReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) FindVolumeForResourceCache(logger lager.Logger, resourceCache *db.UsedResourceCache) (worker.Volume, bool, error) {
	fake.findVolumeForResourceCacheMutex.Lock()
	ret, specificReturn := fake.findVolumeForResourceCacheReturnsOnCall[len(fake.findVolumeForResourceCacheArgsForCall)]
	fake.findVolumeForResourceCacheArgsForCall = append(fake.findVolumeForResourceCacheArgsForCall, struct {
		logger        lager.Logger
		resourceCache *db.UsedResourceCache
	}{logger, resourceCache})
	fake.recordInvocation("FindVolumeForResourceCache", []interface{}{logger, resourceCache})
	fake.findVolumeForResourceCacheMutex.Unlock()
	if fake.FindVolumeForResourceCacheStub != nil {
		return fake.FindVolumeForResourceCacheStub(logger, resourceCache)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findVolumeForResourceCacheReturns.result1, fake.findVolumeForResourceCacheReturns.result2, fake.findVolumeForResourceCacheReturns.result3
}

func (fake *FakeWorker) FindVolumeForResourceCacheCallCount() int {
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findVolumeForResourceCacheArgsForCall)
}

func (fake *FakeWorker) FindVolumeForResourceCacheArgsForCall(i int) (lager.Logger, *db.UsedResourceCache) {
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	return fake.findVolumeForResourceCacheArgsForCall[i].logger, fake.findVolumeForResourceCacheArgsForCall[i].resourceCache
}

func (fake *FakeWorker) FindVolumeForResourceCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeForResourceCacheStub = nil
	fake.findVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) FindVolumeForResourceCacheReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeForResourceCacheStub = nil
	if fake.findVolumeForResourceCacheReturnsOnCall == nil {
		fake.findVolumeForResourceCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeForResourceCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) FindVolumeForTaskCache(arg1 lager.Logger, arg2 int, arg3 int, arg4 string, arg5 string) (worker.Volume, bool, error) {
	fake.findVolumeForTaskCacheMutex.Lock()
	ret, specificReturn := fake.findVolumeForTaskCacheReturnsOnCall[len(fake.findVolumeForTaskCacheArgsForCall)]
	fake.findVolumeForTaskCacheArgsForCall = append(fake.findVolumeForTaskCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindVolumeForTaskCache", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findVolumeForTaskCacheMutex.Unlock()
	if fake.FindVolumeForTaskCacheStub != nil {
		return fake.FindVolumeForTaskCacheStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findVolumeForTaskCacheReturns.result1, fake.findVolumeForTaskCacheReturns.result2, fake.findVolumeForTaskCacheReturns.result3
}

func (fake *FakeWorker) FindVolumeForTaskCacheCallCount() int {
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	return len(fake.findVolumeForTaskCacheArgsForCall)
}

func (fake *FakeWorker) FindVolumeForTaskCacheArgsForCall(i int) (lager.Logger, int, int, string, string) {
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	return fake.findVolumeForTaskCacheArgsForCall[i].arg1, fake.findVolumeForTaskCacheArgsForCall[i].arg2, fake.findVolumeForTaskCacheArgsForCall[i].arg3, fake.findVolumeForTaskCacheArgsForCall[i].arg4, fake.findVolumeForTaskCacheArgsForCall[i].arg5
}

func (fake *FakeWorker) FindVolumeForTaskCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeForTaskCacheStub = nil
	fake.findVolumeForTaskCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) FindVolumeForTaskCacheReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeForTaskCacheStub = nil
	if fake.findVolumeForTaskCacheReturnsOnCall == nil {
		fake.findVolumeForTaskCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeForTaskCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.uptimeMutex.RLock()
	defer fake.uptimeMutex.RUnlock()
	fake.isOwnedByTeamMutex.RLock()
	defer fake.isOwnedByTeamMutex.RUnlock()
	fake.isVersionCompatibleMutex.RLock()
	defer fake.isVersionCompatibleMutex.RUnlock()
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
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

var _ worker.Worker = new(FakeWorker)
