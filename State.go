package main

import (
	"SMTest/SMInterface"
)

type State struct {
	_name       string
	_fsm        SMInterface.IStateMachine
	_transition []SMInterface.ITransition

	onEnterEvent func(self *State, pre SMInterface.IState) // 进入前执行的事件

	onExitEvent func(self *State, next SMInterface.IState) // 推出后执行的事件

	onUpdateEvent func(self *State, dt int64) // 每次更新执行的事件
}

// 创建一个状态对象
func NewState(config StateConfig) *State {
	if config.Name == "" {
		return nil
	}
	state := &State{_name: config.Name, _transition: make([]SMInterface.ITransition, 0)}
	state.registerEnterEvent(config.OnEnterCallBack)
	state.registerExitEvent(config.OnExitCallBack)
	state.registerUpdateEvent(config.OnUpdateCallBack)

	return &State{_name: config.Name, _transition: make([]SMInterface.ITransition, 0)}
}

func (this *State) SetName(name string) {
	this._name = name
}

func (this *State) GetName() string {
	return this._name
}

func (this *State) GetFSM() SMInterface.IStateMachine {
	return this._fsm
}

func (this *State) SetFSM(sm SMInterface.IStateMachine) {
	this._fsm = sm
}

func (this *State) GetTransition() []SMInterface.ITransition {
	return this._transition
}

func (this *State) OnEnter(pre SMInterface.IState) {
	if this.onEnterEvent != nil {
		this.onEnterEvent(this, pre)
	}
}

func (this *State) OnExit(next SMInterface.IState) {
	if this.onExitEvent != nil {
		this.onExitEvent(this, next)
	}
}

func (this *State) Update(dt int64) {
	if this.onUpdateEvent != nil {
		this.onUpdateEvent(this, dt)
	}
}

func (this *State) AddTransition(t SMInterface.ITransition) {
	if this._transition == nil {
		this._transition = make([]SMInterface.ITransition, 0)
	}

	this._transition = append(this._transition, t)
}

func (this *State) registerEnterEvent(callBack func(self *State, pre SMInterface.IState)) {
	this.onEnterEvent = callBack
}

func (this *State) registerExitEvent(callBack func(self *State, next SMInterface.IState)) {
	this.onExitEvent = callBack
}

func (this *State) registerUpdateEvent(callBack func(self *State, dt int64)) {
	this.onUpdateEvent = callBack
}
