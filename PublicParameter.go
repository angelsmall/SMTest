package main

import (
	"SMTest/SMInterface"
	"sync"
)

type PublicParameter struct {
	isInsertCoin bool
	isPassPeople bool
}

var pParamOnce sync.Once
var pParam *PublicParameter

func PublicParameterManager() *PublicParameter{
	pParamOnce.Do( func () {
		pParam = &PublicParameter{}
	} )

	return pParam
}

// 得到是否投入硬币
func (this *PublicParameter) GetIsInsertCoin () bool{
	return this.isInsertCoin
}

// 设置是否投入硬币
func (this *PublicParameter) SetIsInsertCoin (insertCoin bool) {
	this.isInsertCoin = insertCoin
}

// 得到是否有人通过
func (this *PublicParameter) GetIsPassPeople () bool{
	return this.isPassPeople
}

// 设置是否有人通过
func (this *PublicParameter) SetIsPassPeople (passPeople bool) {
	this.isPassPeople = passPeople
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 构造transition的参数
type TransConfig struct {
	FromName string
	ToName string
	OnTransitionCallback func() bool
	OnCheckCallback func() bool
}

// 构造state的参数
type StateConfig struct {
	Name string
	OnEnterCallBack func(self *State, pre SMInterface.IState)
	OnExitCallBack func(self *State, pre SMInterface.IState)
	OnUpdateCallBack func(self *State, dt int64)
}

// 构造stateMachine的参数
type StateMachineConfig struct {
	Name string
	defaultState SMInterface.IState
}