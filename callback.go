package main

// 从开到关的状态切换
func openToClose() bool {
	PublicParameterManager().SetIsInsertCoin(false)
	PublicParameterManager().SetIsPassPeople(false)

	return true
}

// 从开到关的状态检查
func openToCloseCheck() bool{
	if PublicParameterManager().GetIsPassPeople() {
		return true
	}

	return false
}

// 从关到开的状态切换
func closeToOpen() bool {
	PublicParameterManager().SetIsInsertCoin(true)
	PublicParameterManager().SetIsPassPeople(false)

	return true
}

// 从关到开的状态检查
func closeToOpenCheck() bool {
	if PublicParameterManager().GetIsInsertCoin() {
		return true
	}

	return false
}