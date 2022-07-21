# Blue-Moutain-SecondWeek  
#(1)一个结构体所占空间大小与(ABC)有关  
  A 成员本身大小；B 成员对齐系数；C 系统字长  

#(2)占用内存最小的结构体  
在32位系统中  
type A struct{  
  str string  
  num1 int32  
  byte1 byte  
  a struct{}  
}  
16+4+1+1+(2) = 24  

#(3)Go字符串中，每个字符占用多少字节:(AB)  
A 1；B 3；C 1-4  
英文字符为一个字节，中文字符为3个字节  

#(4) Go的map使用的数据结构是：(C)  
  A B+树；B 开放寻址法的Hash表 ；C 拉链法Hash  
hash之后放到对应指针对应的桶中  

#(5)空结构体的地址在任何时候都是zerobase？(是)  
  A 是；B 不是  
  size为0的内存分配都是zerobase  
#(6)空接口就是nil接口？  
  空接口内部没有方法  
  nil接口的值是nil但类型不是nil  
  
