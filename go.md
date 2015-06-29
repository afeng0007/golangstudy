### go语言特性的缺失：
[the way to go](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook)
1. 为了简化设计，不支持函数重载和运算符重载。
2. 不支持隐式转换
3. 放弃了类和类型的继承。
4. 不支持动态加载代码。
5. 不支持动态链接库。
6. 在接口使用方面可以支持变体类型，但是本身不支持变体类型。
7. 不支持动态加载代码。
8. 不支持泛型。
9. 不支持静态变量。
10. 通过`recover`和`panic`来替代异常处理机制。
11. 不支持断言。
12. 不支持静态变量。

#### runtime：
go编译器产生的是本地可执行代码，但是这些代码仍然运行在GO的runtime中，这个runtime类似于Java和.Net语言运行用到的虚拟机。负责管理包括内存分配、垃圾回收，堆栈处理，goroutine，channel，slice，map和reflection机制。
