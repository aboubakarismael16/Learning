## 2022-06-25

## GC in golang

Garbage collection is a mechanism Go developers use to find memory space that is allocated recently but is no longer needed, hence the need to  deallocate them to create a clean slate so that further allocation can  be done on the same space or so that memory can be reused. If this  process is done automatically, without any intervention of the  programmer, it is called *automatic garbage collection*. The term garbage essentially means *unused* or objects that are created in the memory and are no longer needed,  which can be seen as nothing more than garbage, hence a candidate for  wiping out from the memory.



Memory management is definitely easier in Go than it is in, say, C++.  But it’s also not an area we as Go developers can totally ignore,  either. Understanding how Go allocates and frees memory allows us to  write better, more efficient applications. The garbage collector is a  critical piece of that puzzle.



There are two major phases of garbage collection:

**Mark phase**: Identify and mark the objects that are no longer needed by the program.

**Sweep phase**: For every object marked “unreachable” by the mark phase, free up the memory to be used elsewhere



关于go1.8以后的gc

在 go1.15 版本后的 gc，按我的理解是分 4 个阶段：
1.stw 启动扫描 root 对象，然后启动插入屏障和删除屏障
2.start the world 然后用户逻辑和扫描并发执行，扫描 root 下面所有的引用对象改变颜色，然后屏障机制记录修改点
3.stw 重新扫描栈区把漏掉的插入屏障为录入的对象标记然后扫描屏障机制的记录把更改的对象置为灰色 - 黑色 直到没有灰色对象

4. 清除内存垃圾，归还内存
   然后 1.18 版本以后看到大部分文章说，混合写屏障几乎用不到 stw。
   我的问题是：
5. 请教下我对 1.15gc 的理解是否正确？如不对请各位指出一下。
6. 栈区没有 stw，那么堆是否需要 stw，那么堆栈整个过程是怎么样的？

————————————————
原文作者：duck9527
转自链接：https://learnku.com/go/t/65468
版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。

##More reference:

[Golang三色标记混合写屏障GC模式全分析 ](https://learnku.com/articles/68141)