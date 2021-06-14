package main

import (
	"SMTest/SMInterface"
	"reflect"
)

type Transition struct {
	_from SMInterface.IState
	_to   SMInterface.IState

	onTransitionEvent func() bool
	onCheckEvent func() bool
}

func NewTransition(from, to SMInterface.IState) *Transition{
	transition := &Transition{}

	transition._from = from
	transition._to = to
	return transition
}

func (this *Transition) From() SMInterface.IState {
	return this._from
}

func (this *Transition) To() SMInterface.IState {
	return this._to
}

func (this *Transition) OnTransitionEvent(callback func() bool) {
	this.onTransitionEvent = callback
}

func (this *Transition) OnCheckEvent(callback func() bool) {
	this.onCheckEvent = callback
}

func (this *Transition) OnTransition() bool{
	if !reflect.ValueOf(this.onTransitionEvent).IsNil() {
		return this.onTransitionEvent()
	}
	return false
}

func (this *Transition) ShouldBegin() bool {
	if !reflect.ValueOf(this.onCheckEvent).IsNil() {
		return this.onCheckEvent()
	}

	return false
}
