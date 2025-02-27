// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: balance.go
//
// Generated by this command:
//
//	mockgen -source=balance.go -destination=mock/balance.go -package=mock -typed
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/celestia-indexer/internal/storage"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIBalance is a mock of IBalance interface.
type MockIBalance struct {
	ctrl     *gomock.Controller
	recorder *MockIBalanceMockRecorder
}

// MockIBalanceMockRecorder is the mock recorder for MockIBalance.
type MockIBalanceMockRecorder struct {
	mock *MockIBalance
}

// NewMockIBalance creates a new mock instance.
func NewMockIBalance(ctrl *gomock.Controller) *MockIBalance {
	mock := &MockIBalance{ctrl: ctrl}
	mock.recorder = &MockIBalanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBalance) EXPECT() *MockIBalanceMockRecorder {
	return m.recorder
}

// CursorList mocks base method.
func (m *MockIBalance) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIBalanceMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IBalanceCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIBalance)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IBalanceCursorListCall{Call: call}
}

// IBalanceCursorListCall wrap *gomock.Call
type IBalanceCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceCursorListCall) Return(arg0 []*storage.Balance, arg1 error) *IBalanceCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Balance, error)) *IBalanceCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Balance, error)) *IBalanceCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIBalance) GetByID(ctx context.Context, id uint64) (*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIBalanceMockRecorder) GetByID(ctx, id any) *IBalanceGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIBalance)(nil).GetByID), ctx, id)
	return &IBalanceGetByIDCall{Call: call}
}

// IBalanceGetByIDCall wrap *gomock.Call
type IBalanceGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceGetByIDCall) Return(arg0 *storage.Balance, arg1 error) *IBalanceGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceGetByIDCall) Do(f func(context.Context, uint64) (*storage.Balance, error)) *IBalanceGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Balance, error)) *IBalanceGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIBalance) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIBalanceMockRecorder) IsNoRows(err any) *IBalanceIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIBalance)(nil).IsNoRows), err)
	return &IBalanceIsNoRowsCall{Call: call}
}

// IBalanceIsNoRowsCall wrap *gomock.Call
type IBalanceIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceIsNoRowsCall) Return(arg0 bool) *IBalanceIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceIsNoRowsCall) Do(f func(error) bool) *IBalanceIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceIsNoRowsCall) DoAndReturn(f func(error) bool) *IBalanceIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIBalance) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIBalanceMockRecorder) LastID(ctx any) *IBalanceLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIBalance)(nil).LastID), ctx)
	return &IBalanceLastIDCall{Call: call}
}

// IBalanceLastIDCall wrap *gomock.Call
type IBalanceLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceLastIDCall) Return(arg0 uint64, arg1 error) *IBalanceLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceLastIDCall) Do(f func(context.Context) (uint64, error)) *IBalanceLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IBalanceLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIBalance) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIBalanceMockRecorder) List(ctx, limit, offset, order any) *IBalanceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIBalance)(nil).List), ctx, limit, offset, order)
	return &IBalanceListCall{Call: call}
}

// IBalanceListCall wrap *gomock.Call
type IBalanceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceListCall) Return(arg0 []*storage.Balance, arg1 error) *IBalanceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Balance, error)) *IBalanceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Balance, error)) *IBalanceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIBalance) Save(ctx context.Context, m *storage.Balance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIBalanceMockRecorder) Save(ctx, m any) *IBalanceSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIBalance)(nil).Save), ctx, m)
	return &IBalanceSaveCall{Call: call}
}

// IBalanceSaveCall wrap *gomock.Call
type IBalanceSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceSaveCall) Return(arg0 error) *IBalanceSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceSaveCall) Do(f func(context.Context, *storage.Balance) error) *IBalanceSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceSaveCall) DoAndReturn(f func(context.Context, *storage.Balance) error) *IBalanceSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIBalance) Update(ctx context.Context, m *storage.Balance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIBalanceMockRecorder) Update(ctx, m any) *IBalanceUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIBalance)(nil).Update), ctx, m)
	return &IBalanceUpdateCall{Call: call}
}

// IBalanceUpdateCall wrap *gomock.Call
type IBalanceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBalanceUpdateCall) Return(arg0 error) *IBalanceUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBalanceUpdateCall) Do(f func(context.Context, *storage.Balance) error) *IBalanceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBalanceUpdateCall) DoAndReturn(f func(context.Context, *storage.Balance) error) *IBalanceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
