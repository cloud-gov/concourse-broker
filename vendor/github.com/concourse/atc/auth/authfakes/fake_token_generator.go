// This file was generated by counterfeiter
package authfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/auth"
)

type FakeTokenGenerator struct {
	GenerateTokenStub        func(expiration time.Time, teamName string, isAdmin bool) (auth.TokenType, auth.TokenValue, error)
	generateTokenMutex       sync.RWMutex
	generateTokenArgsForCall []struct {
		expiration time.Time
		teamName   string
		isAdmin    bool
	}
	generateTokenReturns struct {
		result1 auth.TokenType
		result2 auth.TokenValue
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenGenerator) GenerateToken(expiration time.Time, teamName string, isAdmin bool) (auth.TokenType, auth.TokenValue, error) {
	fake.generateTokenMutex.Lock()
	fake.generateTokenArgsForCall = append(fake.generateTokenArgsForCall, struct {
		expiration time.Time
		teamName   string
		isAdmin    bool
	}{expiration, teamName, isAdmin})
	fake.recordInvocation("GenerateToken", []interface{}{expiration, teamName, isAdmin})
	fake.generateTokenMutex.Unlock()
	if fake.GenerateTokenStub != nil {
		return fake.GenerateTokenStub(expiration, teamName, isAdmin)
	} else {
		return fake.generateTokenReturns.result1, fake.generateTokenReturns.result2, fake.generateTokenReturns.result3
	}
}

func (fake *FakeTokenGenerator) GenerateTokenCallCount() int {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return len(fake.generateTokenArgsForCall)
}

func (fake *FakeTokenGenerator) GenerateTokenArgsForCall(i int) (time.Time, string, bool) {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return fake.generateTokenArgsForCall[i].expiration, fake.generateTokenArgsForCall[i].teamName, fake.generateTokenArgsForCall[i].isAdmin
}

func (fake *FakeTokenGenerator) GenerateTokenReturns(result1 auth.TokenType, result2 auth.TokenValue, result3 error) {
	fake.GenerateTokenStub = nil
	fake.generateTokenReturns = struct {
		result1 auth.TokenType
		result2 auth.TokenValue
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTokenGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenGenerator) recordInvocation(key string, args []interface{}) {
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

var _ auth.TokenGenerator = new(FakeTokenGenerator)