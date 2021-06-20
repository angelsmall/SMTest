package main

import (
	"SMTest/SMInterface"
	"errors"
	"reflect"
	"sync"
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 状态构造器
var statePoolOnce sync.Once
var sPool *statePool

type statePool struct {
	sm *StateMachine
	m map[string]SMInterface.IState
}

func StatePoolManager() *statePool {
	statePoolOnce.Do(func() {
		sPool = &statePool{m :make(map[string]SMInterface.IState)}
	})
	return sPool
}

// 注册状态
// 这里的注册应该是同步的，不应该并发
func (this *statePool) RegisterState(state SMInterface.IState) error {
	if state == nil || reflect.ValueOf(state).IsNil(){
		return errors.New("the state is nil")
	}

	name := state.GetName()
	if _, ok := this.m[name]; ok {
		return errors.New("the state is registered")
	}

	this.m[name] = state
	return nil
}

// 注册转换器
func (this *statePool) RegisterTransition(config TransConfig) error {
	if config.FromName == "" {
		return errors.New("the parameter is invalid")
	}

	if _, ok := this.m[config.FromName]; !ok {
		return errors.New("the from State is not find")
	}

	fromState := this.m[config.FromName]
	var toState SMInterface.IState
	if _, ok := this.m[config.ToName]; ok && config.ToName == "" {
		toState = this.m[config.ToName]
	}

	transition := NewTransition(fromState, toState)
	transition.onCheckEvent = config.OnCheckCallback
	transition.onTransitionEvent = config.OnTransitionCallback
	fromState.AddTransition(transition)

	return nil
}

// 注册状态机
func (this *statePool) RegisterOnlyOneStateMachine(config StateMachineConfig) error{
	if config.Name == "" {
		return errors.New("the parameter is invalid")
	}

	this.sm = NewStateMachine(config.Name, config.defaultState)

	for _, it := range this.m {
		this.sm.AddState(it)
		it.SetFSM(this.sm)
	}

	return nil
}

// 得到状态机
func (this *statePool) GetOnlyOneStateMachine() *StateMachine{
	return this.sm
}