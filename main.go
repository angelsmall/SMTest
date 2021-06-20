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

// 初始化状态
func initMachine() *StateMachine {
	// 门开启
	openState := NewState(StateConfig{
		Name:             State_Open,
		OnEnterCallBack:  openOnEnterEvent,
		OnExitCallBack:   openOnExitEvent,
		OnUpdateCallBack: openOnUpdateEvent,
	})

	// 门关闭
	closeState := NewState(StateConfig{
		Name:             State_Close,
		OnEnterCallBack:  closeOnEnterEvent,
		OnExitCallBack:   closeOnExitEvent,
		OnUpdateCallBack: closeOnUpdateEvent,
	})

	StatePoolManager().RegisterState(openState)
	StatePoolManager().RegisterState(closeState)

	StatePoolManager().RegisterTransition(TransConfig{
		FromName: State_Open,
		ToName: State_Close,
		OnTransitionCallback: openToClose,
		OnCheckCallback: openToCloseCheck,
	})

	StatePoolManager().RegisterTransition(TransConfig{
		FromName: State_Close,
		ToName: State_Open,
		OnTransitionCallback: closeToOpen,
		OnCheckCallback: closeToOpenCheck,
	})

	StatePoolManager().RegisterOnlyOneStateMachine(StateMachineConfig{
		Name: StateMachine_Subway,
		defaultState: closeState,
	})

	return StatePoolManager().GetOnlyOneStateMachine()
}
