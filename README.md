# go_reflect
学习go语言的反射机制

1.只要反射对象要修改他们表示的对象，就必须获取他们表示的对象的地址

2.Go语言中的类型名称对应的反射获取方法是reflect.Type中的Name()方法

3.reflect.Type的Field()方法返回StructField结构