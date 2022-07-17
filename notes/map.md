# <center>map</center>
### **map要求**
- 唯一、无序，不能是引用类型(切片，函数，以及包含切片的结构类)
- map不能使用cap()函数 能够自动扩容

### **创建map**
- var m1 map[int]string 不能存储数据
- m2 := map[int]string 可以存储数据 
- m3 := make(map[int]string) 可以存储数据 
- m4 := make(map[int]string, 5) 可以存储数据