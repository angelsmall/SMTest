package SMInterface

type IStateMachine interface {
	GetCurState() IState // 得到当前状态

	GetDefState() IState // 得到默认状态

	AddState(state IState) // 添加状态

	RemoveState(state IState) // 移除状态
}
