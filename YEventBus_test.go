package  YEventBus

import (
	"testing"
)

func TestRegisterSend(t *testing.T){
	_flag := 0

	Register("topic",func(){
		_flag++
	})
	Register("topic",func(){
		_flag++
	})
	Send("topic")
	if _flag != 2{
		t.Fatal()
	}
}

func TestRegisterSendWithParameter(t *testing.T){
	_flag := 0
	Register("topic",func(param_str_1_ string){
		if param_str_1_ != ""{
			_flag++
		}
	})
	Register("topic",func(param_str_1_ ,param_str_2_ string){
		if param_str_1_ != ""{
			_flag++
		}
		if param_str_2_ != ""{
			_flag++
		}
	})
	Send("topic","str_1","str_2")
	if _flag != 3{
		t.Fatal()
	}
	_flag = 0
	Send("topic",1,2)
	if _flag != 0{
		t.Fatal()
	}
}

func TestRegisterNonTopicSend(t *testing.T){
	err := Send("none")
	if err == nil {
		t.Fatal()
	}
}


func TestRegisterNonFuncType(t *testing.T){
	defer  func(){
		if r := recover(); r == nil {
			t.Fatal("")
		}
	}()

	Register("topic1",1)
}
