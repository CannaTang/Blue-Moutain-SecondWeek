# Blue-Moutain-SecondWeek
# 基础题第一题
#(1)一个结构体所占空间大小与(ABC)有关
#  A 成员本身大小；B 成员对齐系数；C 系统字长
#(2)占用内存最小的结构体
#type A struct{
#  num1 int32
#  byte1 byte
#  a struct{}
#  str string
#}
#(3) Go的map使用的数据结构是：
#  A B+树；B 开放寻址法的Hash表 ；C 拉链法Hash
#空结构体的地址在任何时候都是zerobase？
#  A 是；B 不是
#空接口就是nil接口？
