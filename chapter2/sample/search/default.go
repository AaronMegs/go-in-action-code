package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// Search implements the behavior for the default matcher.
// 因为 defaultMatcher 是空结构，不需要分配内存，且不需要维护其状态，所以此处方式接收 defaultMatcher 类型的值而不是指针
// defaultMatcher 	接收：defaultMatcher类型的值和指针
// *defaultMatcher	接收：defaultMatcher类型的值和指针
// 普通函数与方法的区别
// 	1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
// 	2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。
// 最佳实践是：将方法的接收者声明为指针
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// init registers the default matcher with the program. - init 将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}
