// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gorm

import (
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: gorm-datasource-config-default
	cominfobuilder.Next()
	cominfobuilder.ID("gorm-datasource-config-default").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfiguration{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gorm-driver-manager
	cominfobuilder.Next()
	cominfobuilder.ID("gorm-driver-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDriverManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gorm-datasource-default
	cominfobuilder.Next()
	cominfobuilder.ID("gorm-datasource-default").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFacade{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfiguration : the factory of component: gorm-datasource-config-default
type comFactory4pComConfiguration struct {

    mPrototype * datasource0x68a737.Configuration

	
	mNameSelector config.InjectionSelector
	mTableNamePrefixSelector config.InjectionSelector
	mTableNameSuffixSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mDriverSelector config.InjectionSelector
	mUsernameSelector config.InjectionSelector
	mPasswordSelector config.InjectionSelector
	mDatabaseSelector config.InjectionSelector

}

func (inst * comFactory4pComConfiguration) init() application.ComponentFactory {

	
	inst.mNameSelector = config.NewInjectionSelector("${datasource.default.name}",nil)
	inst.mTableNamePrefixSelector = config.NewInjectionSelector("${datasource.default.table-name-prefix}",nil)
	inst.mTableNameSuffixSelector = config.NewInjectionSelector("${datasource.default.table-name-suffix}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${datasource.default.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${datasource.default.port}",nil)
	inst.mDriverSelector = config.NewInjectionSelector("${datasource.default.driver}",nil)
	inst.mUsernameSelector = config.NewInjectionSelector("${datasource.default.username}",nil)
	inst.mPasswordSelector = config.NewInjectionSelector("${datasource.default.password}",nil)
	inst.mDatabaseSelector = config.NewInjectionSelector("${datasource.default.database}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComConfiguration) newObject() * datasource0x68a737.Configuration {
	return & datasource0x68a737.Configuration {}
}

func (inst * comFactory4pComConfiguration) castObject(instance application.ComponentInstance) * datasource0x68a737.Configuration {
	return instance.Get().(*datasource0x68a737.Configuration)
}

func (inst * comFactory4pComConfiguration) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComConfiguration) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComConfiguration) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComConfiguration) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfiguration) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComConfiguration) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.TableNamePrefix = inst.getterForFieldTableNamePrefixSelector(context)
	obj.TableNameSuffix = inst.getterForFieldTableNameSuffixSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.Driver = inst.getterForFieldDriverSelector(context)
	obj.Username = inst.getterForFieldUsernameSelector(context)
	obj.Password = inst.getterForFieldPasswordSelector(context)
	obj.Database = inst.getterForFieldDatabaseSelector(context)
	return context.LastError()
}

//getterForFieldNameSelector
func (inst * comFactory4pComConfiguration) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldTableNamePrefixSelector
func (inst * comFactory4pComConfiguration) getterForFieldTableNamePrefixSelector (context application.InstanceContext) string {
    return inst.mTableNamePrefixSelector.GetString(context)
}

//getterForFieldTableNameSuffixSelector
func (inst * comFactory4pComConfiguration) getterForFieldTableNameSuffixSelector (context application.InstanceContext) string {
    return inst.mTableNameSuffixSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComConfiguration) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComConfiguration) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldDriverSelector
func (inst * comFactory4pComConfiguration) getterForFieldDriverSelector (context application.InstanceContext) string {
    return inst.mDriverSelector.GetString(context)
}

//getterForFieldUsernameSelector
func (inst * comFactory4pComConfiguration) getterForFieldUsernameSelector (context application.InstanceContext) string {
    return inst.mUsernameSelector.GetString(context)
}

//getterForFieldPasswordSelector
func (inst * comFactory4pComConfiguration) getterForFieldPasswordSelector (context application.InstanceContext) string {
    return inst.mPasswordSelector.GetString(context)
}

//getterForFieldDatabaseSelector
func (inst * comFactory4pComConfiguration) getterForFieldDatabaseSelector (context application.InstanceContext) string {
    return inst.mDatabaseSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDriverManager : the factory of component: gorm-driver-manager
type comFactory4pComDriverManager struct {

    mPrototype * datasource0x68a737.DriverManager

	
	mDriversSelector config.InjectionSelector

}

func (inst * comFactory4pComDriverManager) init() application.ComponentFactory {

	
	inst.mDriversSelector = config.NewInjectionSelector(".gorm-datasource-driver",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDriverManager) newObject() * datasource0x68a737.DriverManager {
	return & datasource0x68a737.DriverManager {}
}

func (inst * comFactory4pComDriverManager) castObject(instance application.ComponentInstance) * datasource0x68a737.DriverManager {
	return instance.Get().(*datasource0x68a737.DriverManager)
}

func (inst * comFactory4pComDriverManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDriverManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDriverManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDriverManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDriverManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDriverManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Drivers = inst.getterForFieldDriversSelector(context)
	return context.LastError()
}

//getterForFieldDriversSelector
func (inst * comFactory4pComDriverManager) getterForFieldDriversSelector (context application.InstanceContext) []datasource0x68a737.Driver {
	list1 := inst.mDriversSelector.GetList(context)
	list2 := make([]datasource0x68a737.Driver, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(datasource0x68a737.Driver)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFacade : the factory of component: gorm-datasource-default
type comFactory4pComFacade struct {

    mPrototype * datasource0x68a737.Facade

	
	mDriversSelector config.InjectionSelector
	mConfigSelector config.InjectionSelector

}

func (inst * comFactory4pComFacade) init() application.ComponentFactory {

	
	inst.mDriversSelector = config.NewInjectionSelector("#gorm-driver-manager",nil)
	inst.mConfigSelector = config.NewInjectionSelector("#gorm-datasource-config-default",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFacade) newObject() * datasource0x68a737.Facade {
	return & datasource0x68a737.Facade {}
}

func (inst * comFactory4pComFacade) castObject(instance application.ComponentInstance) * datasource0x68a737.Facade {
	return instance.Get().(*datasource0x68a737.Facade)
}

func (inst * comFactory4pComFacade) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFacade) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFacade) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFacade) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Open()
}

func (inst * comFactory4pComFacade) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Close()
}

func (inst * comFactory4pComFacade) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Drivers = inst.getterForFieldDriversSelector(context)
	obj.Config = inst.getterForFieldConfigSelector(context)
	return context.LastError()
}

//getterForFieldDriversSelector
func (inst * comFactory4pComFacade) getterForFieldDriversSelector (context application.InstanceContext) datasource0x68a737.Drivers {

	o1 := inst.mDriversSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Drivers)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gorm-datasource-default")
		eb.Set("field", "Drivers")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Drivers")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldConfigSelector
func (inst * comFactory4pComFacade) getterForFieldConfigSelector (context application.InstanceContext) *datasource0x68a737.Configuration {

	o1 := inst.mConfigSelector.GetOne(context)
	o2, ok := o1.(*datasource0x68a737.Configuration)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gorm-datasource-default")
		eb.Set("field", "Config")
		eb.Set("type1", "?")
		eb.Set("type2", "*datasource0x68a737.Configuration")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




