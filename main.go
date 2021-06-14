package main

/*
* (网络搜索有关状态机的东西，找到一个题目，简单做一下)
* 假设地铁的门有2个状态，开和关。
* 在关的情况下，投入一个硬币，门就开了；
* 在关的状态下，强行进入，就发出警告。
* 在开的情况下，人一进入，门就关了；
* 在开的情况下，继续投币，就说：“thank you”。
*/

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	machine := initMachine()

	go stdin()

	for {
		time.Sleep(time.Second)
		machine.Update(time.Now().Unix())
	}

	fmt.Println("hello world!")
}

func stdin() {
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	// 逐行扫描
	for input.Scan() {
		line := input.Text()

		// 输入bye时 结束
		if line == "bye" {
			break
		}

		if line == "insertCoin" {
			PublicParameterManager().SetIsInsertCoin(true)
		} else if line == "pass" {
			PublicParameterManager().SetIsPassPeople(true)
		}

	}

}


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


// 初始化状态
func initMachine() *StateMachine{
	// 门开启
	openState := NewState("open")
	openState.OnEnterEvent(openOnEnterEvent)
	openState.OnExitEvent(openOnExitEvent)
	openState.OnUpdateEvent(openOnUpdateEvent)

	// 门关闭
	closeState := NewState("close")
	closeState.OnEnterEvent(closeOnEnterEvent)
	closeState.OnExitEvent(closeOnExitEvent)
	closeState.OnUpdateEvent(closeOnUpdateEvent)

	// 从状态开到关
	openToCloseTrans := NewTransition(openState, closeState)
	openToCloseTrans.OnTransitionEvent(openToClose)
	openToCloseTrans.OnCheckEvent(openToCloseCheck)
	openState.AddTransition(openToCloseTrans)


	// 从状态关到开
	closeToOpenTrans := NewTransition(closeState, openState)
	closeToOpenTrans.OnTransitionEvent(closeToOpen)
	closeToOpenTrans.OnCheckEvent(closeToOpenCheck)
	closeState.AddTransition(closeToOpenTrans)

	machine := NewStateMachine("subway", closeState)
	machine.AddState(openState)
	machine.AddState(closeState)

	return machine
}


