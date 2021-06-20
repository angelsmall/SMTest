package SMInterface

type IState interface {
	GetName() string // 获得状态的名字

	GetFSM() IStateMachine // 获得当前状态机
	SetFSM(sm IStateMachine)	// 设置当前状态机

	GetTransition() []ITransition // 获得转换器列表

	OnEnter(state IState) // 状态开始时

	OnExit(state IState) // 状态结束时

	Update(dt int64) // 更新

	AddTransition(t ITransition) // 加入转换器
}
