// Code generated by counterfeiter. DO NOT EDIT.
package registryfakes

import (
	sync "sync"

	layout "github.com/google/go-containerregistry/pkg/v1/layout"
	image "github.com/pivotal/image-relocation/pkg/image"
	registry "github.com/pivotal/image-relocation/pkg/registry"
)

type FakeImage struct {
	AppendToLayoutStub        func(registry.LayoutPath, ...layout.Option) error
	appendToLayoutMutex       sync.RWMutex
	appendToLayoutArgsForCall []struct {
		arg1 registry.LayoutPath
		arg2 []layout.Option
	}
	appendToLayoutReturns struct {
		result1 error
	}
	appendToLayoutReturnsOnCall map[int]struct {
		result1 error
	}
	DigestStub        func() (image.Digest, error)
	digestMutex       sync.RWMutex
	digestArgsForCall []struct {
	}
	digestReturns struct {
		result1 image.Digest
		result2 error
	}
	digestReturnsOnCall map[int]struct {
		result1 image.Digest
		result2 error
	}
	WriteStub        func(image.Name) (image.Digest, int64, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 image.Name
	}
	writeReturns struct {
		result1 image.Digest
		result2 int64
		result3 error
	}
	writeReturnsOnCall map[int]struct {
		result1 image.Digest
		result2 int64
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImage) AppendToLayout(arg1 registry.LayoutPath, arg2 ...layout.Option) error {
	fake.appendToLayoutMutex.Lock()
	ret, specificReturn := fake.appendToLayoutReturnsOnCall[len(fake.appendToLayoutArgsForCall)]
	fake.appendToLayoutArgsForCall = append(fake.appendToLayoutArgsForCall, struct {
		arg1 registry.LayoutPath
		arg2 []layout.Option
	}{arg1, arg2})
	fake.recordInvocation("AppendToLayout", []interface{}{arg1, arg2})
	fake.appendToLayoutMutex.Unlock()
	if fake.AppendToLayoutStub != nil {
		return fake.AppendToLayoutStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.appendToLayoutReturns
	return fakeReturns.result1
}

func (fake *FakeImage) AppendToLayoutCallCount() int {
	fake.appendToLayoutMutex.RLock()
	defer fake.appendToLayoutMutex.RUnlock()
	return len(fake.appendToLayoutArgsForCall)
}

func (fake *FakeImage) AppendToLayoutCalls(stub func(registry.LayoutPath, ...layout.Option) error) {
	fake.appendToLayoutMutex.Lock()
	defer fake.appendToLayoutMutex.Unlock()
	fake.AppendToLayoutStub = stub
}

func (fake *FakeImage) AppendToLayoutArgsForCall(i int) (registry.LayoutPath, []layout.Option) {
	fake.appendToLayoutMutex.RLock()
	defer fake.appendToLayoutMutex.RUnlock()
	argsForCall := fake.appendToLayoutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImage) AppendToLayoutReturns(result1 error) {
	fake.appendToLayoutMutex.Lock()
	defer fake.appendToLayoutMutex.Unlock()
	fake.AppendToLayoutStub = nil
	fake.appendToLayoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImage) AppendToLayoutReturnsOnCall(i int, result1 error) {
	fake.appendToLayoutMutex.Lock()
	defer fake.appendToLayoutMutex.Unlock()
	fake.AppendToLayoutStub = nil
	if fake.appendToLayoutReturnsOnCall == nil {
		fake.appendToLayoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.appendToLayoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImage) Digest() (image.Digest, error) {
	fake.digestMutex.Lock()
	ret, specificReturn := fake.digestReturnsOnCall[len(fake.digestArgsForCall)]
	fake.digestArgsForCall = append(fake.digestArgsForCall, struct {
	}{})
	fake.recordInvocation("Digest", []interface{}{})
	fake.digestMutex.Unlock()
	if fake.DigestStub != nil {
		return fake.DigestStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.digestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) DigestCallCount() int {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	return len(fake.digestArgsForCall)
}

func (fake *FakeImage) DigestCalls(stub func() (image.Digest, error)) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = stub
}

func (fake *FakeImage) DigestReturns(result1 image.Digest, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	fake.digestReturns = struct {
		result1 image.Digest
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) DigestReturnsOnCall(i int, result1 image.Digest, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	if fake.digestReturnsOnCall == nil {
		fake.digestReturnsOnCall = make(map[int]struct {
			result1 image.Digest
			result2 error
		})
	}
	fake.digestReturnsOnCall[i] = struct {
		result1 image.Digest
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Write(arg1 image.Name) (image.Digest, int64, error) {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 image.Name
	}{arg1})
	fake.recordInvocation("Write", []interface{}{arg1})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.writeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeImage) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeImage) WriteCalls(stub func(image.Name) (image.Digest, int64, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeImage) WriteArgsForCall(i int) image.Name {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImage) WriteReturns(result1 image.Digest, result2 int64, result3 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 image.Digest
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImage) WriteReturnsOnCall(i int, result1 image.Digest, result2 int64, result3 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 image.Digest
			result2 int64
			result3 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 image.Digest
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendToLayoutMutex.RLock()
	defer fake.appendToLayoutMutex.RUnlock()
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImage) recordInvocation(key string, args []interface{}) {
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

var _ registry.Image = new(FakeImage)
