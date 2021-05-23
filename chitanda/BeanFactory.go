package chitanda

import (
	"reflect"
)

type BeanFactory struct {
	beans []interface{}
}
func NewBeanFactory() *BeanFactory {
	bf:= &BeanFactory{beans: make([]interface{},0)}
	bf.beans=append(bf.beans,bf)
	return bf
}

func(this *BeanFactory) setBean(beans ...interface{}){
	this.beans=append(this.beans,beans...)
}

func(this *BeanFactory) GetBean(bean interface{}) interface{}{
	return this.getBean(reflect.TypeOf(bean))
}

func(this *BeanFactory) getBean(t reflect.Type) interface{} {
	for _,p:=range this.beans{
		if t==reflect.TypeOf(p){
			return p
		}
	}
	return nil
}
func(this *BeanFactory) Inject(object interface{}){
	vObject:=reflect.ValueOf(object)
	if vObject.Kind()==reflect.Ptr{
		vObject=vObject.Elem()
	}
	for i:=0;i<vObject.NumField();i++{
		f:=vObject.Field(i)
		if f.Kind()!=reflect.Ptr || !f.IsNil()   {
			continue
		}

		if p:=this.getBean(f.Type());p!=nil && f.CanInterface(){
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

func(this *BeanFactory) inject(class IClass){
	vClass:=reflect.ValueOf(class).Elem()
	vClassT:=reflect.TypeOf(class).Elem()
	for i:=0;i<vClass.NumField();i++{
		f:=vClass.Field(i)
		if f.Kind()!=reflect.Ptr || !f.IsNil()  {
			continue
		}

		if IsAnnotation(f.Type()){
			f.Set(reflect.New(f.Type().Elem()))
			f.Interface().(Annotation).SetTag(vClassT.Field(i).Tag)
			this.Inject(f.Interface())
			continue
		}
		if p:=this.getBean(f.Type());p!=nil{
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
