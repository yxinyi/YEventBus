### 动态参数事件系统

### 使用

##### 事件注册

```
Register("topic",func(){
	fmt.Println("Hello YEventBus")
})
```

##### 事件触发

```
Send("topic")
```

##### 动态参数注册

```
Register("topic",func(str_1_,str_2_ string){
	fmt.Println("Hello",str_1_,str_2_)
})
```

##### 动态参数事件触发

```
Send("topic")
```
会输出`"Hello"`

```
Send("topic","YEventBus")
```
会输出`"HelloYEventBus"`

```
Send("topic","YEventBus","Again")
```
会输出`"HelloYEventBus Again"`

### 后续

1.需要加上死循环检测

2.可能会增加异步功能