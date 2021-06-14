package main

import (
	"SMTest/SMInterface"
	"fmt"
)

/////////////////////////////////////////////////////////////////
// 打开的事件
func openOnEnterEvent(self *State, pre SMInterface.IState) {
	if self.GetName() != "open" {
		return
	}

	if pre == nil {
		// TODO 自动关门
		fmt.Println("系统已启动，正在关门中，请勿通过，以免夹伤")
	}

	if PublicParameterManager().GetIsInsertCoin() {
		// 打开之后需要付钱
		PublicParameterManager().SetIsInsertCoin(false)
		fmt.Println("验证成功")
	}
}

func openOnExitEvent(self *State, next SMInterface.IState) {
	if self.GetName() != "open" {
		return
	}

	fmt.Println("已通过")
}


func openOnUpdateEvent(self *State, dt int64) {
	if self.GetName() != "open" {
		return
	}

	if PublicParameterManager().GetIsInsertCoin() {
		fmt.Println("thank you")
		PublicParameterManager().SetIsInsertCoin(false)
	} else {
		fmt.Println("门已打开，请快速通过")
	}

	// TODO 这里设置一个10秒的定时器，用来自动关门
}

///////////////////////////////////////////////////////////////
// 关闭的事件
func closeOnEnterEvent(self *State, pre SMInterface.IState) {
	if self.GetName() != "close" {
		return
	}

	if pre == nil {
		fmt.Println("系统已启动，欢迎乘坐地铁")
	}
}

func closeOnExitEvent(self *State, next SMInterface.IState) {
	if self.GetName() != "close" {
		return
	}

	if next == nil {
		fmt.Println("出现故障，请联系管理人员")
	} else {
		fmt.Println("请通过")
	}
}


func closeOnUpdateEvent(self *State, dt int64) {
	if self.GetName() != "close" {
		return
	}

	if PublicParameterManager().GetIsInsertCoin() {
		fmt.Println("请稍等，正在检验中...")
	} else if PublicParameterManager().GetIsPassPeople() {
		fmt.Println("警告！有人逃票！有人逃票！")
		PublicParameterManager().SetIsPassPeople(false)
	} else {
		fmt.Println("地铁欢迎您！请投币！")
	}
}





