1. 怎么判断goroutine是否泄漏
2. goleak是怎么实现呢？
   在函数最后defer调用,
   通过runtime.Stack获取goroutine的跟踪栈
   解析goId, gStatus, 调用函数,等
   再进行filter不是(
       isTestStack
       isSyscallStack
       isStdLibStack
       isTraceStack
   )