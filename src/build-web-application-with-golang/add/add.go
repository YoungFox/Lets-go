package add

//Add1 简单的一个函数，实现了参数+1的操作
func Add1(a *int) *int {
	*a = *a + 1 // 我们改变了a的值
	return a    //返回一个新值
}
