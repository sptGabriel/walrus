// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecases

import (
	"context"
	"sync"

	"github.com/matheusmosca/walrus/domain/entities"
	"github.com/matheusmosca/walrus/domain/vos"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			CreateTopicFunc: func(ctx context.Context, name vos.TopicName, topic entities.Topic) error {
// 				panic("mock out the CreateTopic method")
// 			},
// 			GetTopicFunc: func(ctx context.Context, topicName vos.TopicName) (entities.Topic, error) {
// 				panic("mock out the GetTopic method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// CreateTopicFunc mocks the CreateTopic method.
	CreateTopicFunc func(ctx context.Context, name vos.TopicName, topic entities.Topic) error

	// GetTopicFunc mocks the GetTopic method.
	GetTopicFunc func(ctx context.Context, topicName vos.TopicName) (entities.Topic, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateTopic holds details about calls to the CreateTopic method.
		CreateTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name vos.TopicName
			// Topic is the topic argument value.
			Topic entities.Topic
		}
		// GetTopic holds details about calls to the GetTopic method.
		GetTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TopicName is the topicName argument value.
			TopicName vos.TopicName
		}
	}
	lockCreateTopic sync.RWMutex
	lockGetTopic    sync.RWMutex
}

// CreateTopic calls CreateTopicFunc.
func (mock *RepositoryMock) CreateTopic(ctx context.Context, name vos.TopicName, topic entities.Topic) error {
	if mock.CreateTopicFunc == nil {
		panic("RepositoryMock.CreateTopicFunc: method is nil but Repository.CreateTopic was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Name  vos.TopicName
		Topic entities.Topic
	}{
		Ctx:   ctx,
		Name:  name,
		Topic: topic,
	}
	mock.lockCreateTopic.Lock()
	mock.calls.CreateTopic = append(mock.calls.CreateTopic, callInfo)
	mock.lockCreateTopic.Unlock()
	return mock.CreateTopicFunc(ctx, name, topic)
}

// CreateTopicCalls gets all the calls that were made to CreateTopic.
// Check the length with:
//     len(mockedRepository.CreateTopicCalls())
func (mock *RepositoryMock) CreateTopicCalls() []struct {
	Ctx   context.Context
	Name  vos.TopicName
	Topic entities.Topic
} {
	var calls []struct {
		Ctx   context.Context
		Name  vos.TopicName
		Topic entities.Topic
	}
	mock.lockCreateTopic.RLock()
	calls = mock.calls.CreateTopic
	mock.lockCreateTopic.RUnlock()
	return calls
}

// GetTopic calls GetTopicFunc.
func (mock *RepositoryMock) GetTopic(ctx context.Context, topicName vos.TopicName) (entities.Topic, error) {
	if mock.GetTopicFunc == nil {
		panic("RepositoryMock.GetTopicFunc: method is nil but Repository.GetTopic was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		TopicName vos.TopicName
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockGetTopic.Lock()
	mock.calls.GetTopic = append(mock.calls.GetTopic, callInfo)
	mock.lockGetTopic.Unlock()
	return mock.GetTopicFunc(ctx, topicName)
}

// GetTopicCalls gets all the calls that were made to GetTopic.
// Check the length with:
//     len(mockedRepository.GetTopicCalls())
func (mock *RepositoryMock) GetTopicCalls() []struct {
	Ctx       context.Context
	TopicName vos.TopicName
} {
	var calls []struct {
		Ctx       context.Context
		TopicName vos.TopicName
	}
	mock.lockGetTopic.RLock()
	calls = mock.calls.GetTopic
	mock.lockGetTopic.RUnlock()
	return calls
}
