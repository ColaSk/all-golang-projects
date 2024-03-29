# go 语言基础

## 特殊问题

### import、包名、目录名的关系

- 结论
  - import后面的是目录 **import "xxx"**
  - import重命名后的名称为包名 **import m "xxx"**
  - 包名和目录名没有关系，但是包名最好等于目录名
  - 同一个目录下只能有一种包名

> 如果包名与文件夹名称不同，按照原理来看执行不存在问题，假设存在问题 **imported and not used: "go-basis/basis" as basiss**
> 这是由于安装的一些规范插件造成的，为了避免不必要的麻烦我们还是需要保持 **包名与文件夹名称保持一致**

### 函数, 结构体的大小写的影响

> **类名大写**表示其他包可以被访问
> **属性名大写**表示该属性名对外能够访问
> **方法名大写**表示该方法其他包可访问

### go map 数据结构与扩容

> go map是基本数据结构是hash数组+桶内的key-value数组+溢出的桶链表 当hash表超过阈值需要扩容增长时，会分配一个新的数组，新数组的大小一般是旧数组的2倍
