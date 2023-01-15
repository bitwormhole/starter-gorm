// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgstartergorm

import (
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	dsimpl0x33f250 "github.com/bitwormhole/starter-gorm/datasource/dsimpl"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-dsimpl0x33f250.DefaultDataSource
	cominfobuilder.Next()
	cominfobuilder.ID("com0-dsimpl0x33f250.DefaultDataSource").Class("starter-gorm-source-registry life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultDataSource{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: starter-gorm-driver-manager
	cominfobuilder.Next()
	cominfobuilder.ID("starter-gorm-driver-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDriverManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: starter-gorm-source-manager
	cominfobuilder.Next()
	cominfobuilder.ID("starter-gorm-source-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSourceManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultDataSource : the factory of component: com0-dsimpl0x33f250.DefaultDataSource
type comFactory4pComDefaultDataSource struct {

    mPrototype * dsimpl0x33f250.DefaultDataSource

	
	mDriversSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mDriverSelector config.InjectionSelector
	mUsernameSelector config.InjectionSelector
	mPasswordSelector config.InjectionSelector
	mDatabaseSelector config.InjectionSelector
	mTableNamePrefixSelector config.InjectionSelector
	mTableNameSuffixSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mLazyLoadSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultDataSource) init() application.ComponentFactory {

	
	inst.mDriversSelector = config.NewInjectionSelector("#starter-gorm-driver-manager",nil)
	inst.mNameSelector = config.NewInjectionSelector("${datasource.default.name}",nil)
	inst.mDriverSelector = config.NewInjectionSelector("${datasource.default.driver}",nil)
	inst.mUsernameSelector = config.NewInjectionSelector("${datasource.default.username}",nil)
	inst.mPasswordSelector = config.NewInjectionSelector("${datasource.default.password}",nil)
	inst.mDatabaseSelector = config.NewInjectionSelector("${datasource.default.database}",nil)
	inst.mTableNamePrefixSelector = config.NewInjectionSelector("${datasource.default.table-name-prefix}",nil)
	inst.mTableNameSuffixSelector = config.NewInjectionSelector("${datasource.default.table-name-suffix}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${datasource.default.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${datasource.default.port}",nil)
	inst.mLazyLoadSelector = config.NewInjectionSelector("${datasource.default.lazyload}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultDataSource) newObject() * dsimpl0x33f250.DefaultDataSource {
	return & dsimpl0x33f250.DefaultDataSource {}
}

func (inst * comFactory4pComDefaultDataSource) castObject(instance application.ComponentInstance) * dsimpl0x33f250.DefaultDataSource {
	return instance.Get().(*dsimpl0x33f250.DefaultDataSource)
}

func (inst * comFactory4pComDefaultDataSource) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultDataSource) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultDataSource) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultDataSource) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultDataSource) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultDataSource) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Drivers = inst.getterForFieldDriversSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Driver = inst.getterForFieldDriverSelector(context)
	obj.Username = inst.getterForFieldUsernameSelector(context)
	obj.Password = inst.getterForFieldPasswordSelector(context)
	obj.Database = inst.getterForFieldDatabaseSelector(context)
	obj.TableNamePrefix = inst.getterForFieldTableNamePrefixSelector(context)
	obj.TableNameSuffix = inst.getterForFieldTableNameSuffixSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.LazyLoad = inst.getterForFieldLazyLoadSelector(context)
	return context.LastError()
}

//getterForFieldDriversSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldDriversSelector (context application.InstanceContext) datasource0x68a737.DriverManager {

	o1 := inst.mDriversSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.DriverManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-dsimpl0x33f250.DefaultDataSource")
		eb.Set("field", "Drivers")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.DriverManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldNameSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldDriverSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldDriverSelector (context application.InstanceContext) string {
    return inst.mDriverSelector.GetString(context)
}

//getterForFieldUsernameSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldUsernameSelector (context application.InstanceContext) string {
    return inst.mUsernameSelector.GetString(context)
}

//getterForFieldPasswordSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldPasswordSelector (context application.InstanceContext) string {
    return inst.mPasswordSelector.GetString(context)
}

//getterForFieldDatabaseSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldDatabaseSelector (context application.InstanceContext) string {
    return inst.mDatabaseSelector.GetString(context)
}

//getterForFieldTableNamePrefixSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldTableNamePrefixSelector (context application.InstanceContext) string {
    return inst.mTableNamePrefixSelector.GetString(context)
}

//getterForFieldTableNameSuffixSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldTableNameSuffixSelector (context application.InstanceContext) string {
    return inst.mTableNameSuffixSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldLazyLoadSelector
func (inst * comFactory4pComDefaultDataSource) getterForFieldLazyLoadSelector (context application.InstanceContext) bool {
    return inst.mLazyLoadSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDriverManagerImpl : the factory of component: starter-gorm-driver-manager
type comFactory4pComDriverManagerImpl struct {

    mPrototype * dsimpl0x33f250.DriverManagerImpl

	
	mDriversSelector config.InjectionSelector

}

func (inst * comFactory4pComDriverManagerImpl) init() application.ComponentFactory {

	
	inst.mDriversSelector = config.NewInjectionSelector(".starter-gorm-driver-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDriverManagerImpl) newObject() * dsimpl0x33f250.DriverManagerImpl {
	return & dsimpl0x33f250.DriverManagerImpl {}
}

func (inst * comFactory4pComDriverManagerImpl) castObject(instance application.ComponentInstance) * dsimpl0x33f250.DriverManagerImpl {
	return instance.Get().(*dsimpl0x33f250.DriverManagerImpl)
}

func (inst * comFactory4pComDriverManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDriverManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDriverManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDriverManagerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDriverManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDriverManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Drivers = inst.getterForFieldDriversSelector(context)
	return context.LastError()
}

//getterForFieldDriversSelector
func (inst * comFactory4pComDriverManagerImpl) getterForFieldDriversSelector (context application.InstanceContext) []datasource0x68a737.DriverRegistry {
	list1 := inst.mDriversSelector.GetList(context)
	list2 := make([]datasource0x68a737.DriverRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(datasource0x68a737.DriverRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSourceManagerImpl : the factory of component: starter-gorm-source-manager
type comFactory4pComSourceManagerImpl struct {

    mPrototype * dsimpl0x33f250.SourceManagerImpl

	
	mSourcesSelector config.InjectionSelector

}

func (inst * comFactory4pComSourceManagerImpl) init() application.ComponentFactory {

	
	inst.mSourcesSelector = config.NewInjectionSelector(".starter-gorm-source-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSourceManagerImpl) newObject() * dsimpl0x33f250.SourceManagerImpl {
	return & dsimpl0x33f250.SourceManagerImpl {}
}

func (inst * comFactory4pComSourceManagerImpl) castObject(instance application.ComponentInstance) * dsimpl0x33f250.SourceManagerImpl {
	return instance.Get().(*dsimpl0x33f250.SourceManagerImpl)
}

func (inst * comFactory4pComSourceManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSourceManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSourceManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSourceManagerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSourceManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSourceManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Sources = inst.getterForFieldSourcesSelector(context)
	return context.LastError()
}

//getterForFieldSourcesSelector
func (inst * comFactory4pComSourceManagerImpl) getterForFieldSourcesSelector (context application.InstanceContext) []datasource0x68a737.SourceRegistry {
	list1 := inst.mSourcesSelector.GetList(context)
	list2 := make([]datasource0x68a737.SourceRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(datasource0x68a737.SourceRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}




