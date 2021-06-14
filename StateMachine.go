package main

import (
	"SMTest/SMInterface"
	"reflect"
)

type StateMachine struct {
	State

	_curState      SMInterface.IState
	_defState      SMInterface.IState
	_states        []SMInterface.IState
	_isTransiting  bool
	_curTransition SMInterface.ITransition
}

// 创建一个状态机
func NewStateMachine(name string, defState SMInterface.IState) *StateMachine {
	sm := &StateMachine{}

	sm.SetName(name)
	sm._states = make([]SMInterface.IState, 0)

	if defState != nil && !reflect.ValueOf(defState).IsNil() {
		sm._defState = defState
		sm._states = append(sm._states, defState)
	}

	return sm
}

// 获得当前状态
func (this *StateMachine) GetCurState() SMInterface.IState {
	return this._curState
}

// 获得默认状态
func (this *StateMachine) GetDefState() SMInterface.IState {
	return this._curState
}

func (this *StateMachine) AddState(state SMInterface.IState) {
	if !reflect.ValueOf(state).IsNil() {
		this._states = append(this._states, state)
	}
}

func (this *StateMachine) RemoveState(state SMInterface.IState) {
	if reflect.ValueOf(state).IsNil() {
		return
	}

	for i := 0; i < len(this._states); i++ {
		if this._states[i].GetName() == state.GetName() {
			this._states = append(this._states[:i], this._states[i+1:]...)
			break
		}
	}
}

func (this *StateMachine) Update(dt int64) {
	if this._isTransiting {
		if this._curTransition.OnTransition() {
			this.doTransition(this._curTransition)
			this._isTransiting = false
		}
		return
	}

	this.State.Update(dt)

	if this._curState == nil || reflect.ValueOf(this._curState).IsNil() {
		this._curState = this._defState
		this._curState.OnEnter(nil)
	}

	transition := this._curState.GetTransition()
	for i := 0; i < len(transition); i++ {
		if transition[i].ShouldBegin() {
			this._curTransition = transition[i]
			this._isTransiting = true
			return
		}
	}

	// 如果不做切换，则跟新当前状态
	this._curState.Update(dt)
}

func (this *StateMachine) doTransition(transition SMInterface.ITransition) {
	if transition.To() == nil {
		return
	}
	form := transition.From()
	to := transition.To()

	to.OnEnter(form)    // 预加载
	this._curState = to // 切换状态
	form.OnExit(to)     // 结束上一个状态
}
