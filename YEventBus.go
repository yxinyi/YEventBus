package  YEventBus

import (
	"fmt"
	"reflect"
)

var ErrorFunc = fmt.Errorf

type EventHandler struct{
	m_topic string
	m_fn reflect.Value
}

var g_event_handler_list = make(map[string][]*EventHandler)

func Register(topic_ string, fn_ interface{}){
	_,exists := g_event_handler_list[topic_]
	if !exists{
		g_event_handler_list[topic_] = make([]*EventHandler,0)
	}
	if reflect.TypeOf(fn_).Kind() != reflect.Func{
		panic("[YEventBus] [Register] register handler not is func type")
	}
	g_event_handler_list[topic_] = append(g_event_handler_list[topic_], &EventHandler{
		topic_,
		reflect.ValueOf(fn_),
	})
}


func getCallBackParameterList(call_back_type_ reflect.Type, args_ ... interface{})[]reflect.Value{
	_parameter_list := make([]reflect.Value,0,len(args_))
	for _idx := 0 ; _idx < call_back_type_.NumIn();_idx++{
		//如果没有传入值,则默认生成一个零值作为参数
		if _idx >= len(args_){
			_parameter_list = append(_parameter_list, reflect.New(call_back_type_.In(_idx)).Elem())
			continue
		}
		//如果参数不对,则生成一个零值
		if call_back_type_.In(_idx).Kind() !=  reflect.ValueOf(args_[_idx]).Kind(){
			_parameter_list = append(_parameter_list, reflect.New(call_back_type_.In(_idx)).Elem())
			continue
		}
			_parameter_list = append(_parameter_list, reflect.ValueOf(args_[_idx]))
	}
	return _parameter_list
}

func Send(topic_ string, args_ ... interface{}) error{
	_handler_list, _exists := g_event_handler_list[topic_]
	if !_exists {
		return ErrorFunc("[YEventBus] [Send] [%v] error not has this topic ",topic_)
	}
	for _,_it := range _handler_list{
		_it.m_fn.Call(getCallBackParameterList(_it.m_fn.Type(),args_...))
	}
	return nil
}
