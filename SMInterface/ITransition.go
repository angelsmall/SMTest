package SMInterface

type ITransition interface {
	From() IState // 获取原状态

	To() IState // 获取目标状态

	OnTransition() bool // 切换状态

	ShouldBegin() bool // 是否能够切换
}
