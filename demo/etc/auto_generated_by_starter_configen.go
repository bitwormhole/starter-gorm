// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	data_5df53416 "github.com/bitwormhole/starter-gorm/data"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // _gorm_db_drivers
    cb.AddComponent(&config.ComInfo{
		ID: "_gorm_db_drivers",
		Class: "gorm-db-drivers",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &data_5df53416.DBDriverRegistrar{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*data_5df53416.DBDriverRegistrar)
		    return _gorm_db_drivers(target,context)
		},
    })

    // _gorm_starter
    cb.AddComponent(&config.ComInfo{
		ID: "_gorm_starter",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &data_5df53416.GormStarter{}
		},
		OnInit: func(obj lang.Object) error {
		    target := obj.(*data_5df53416.GormStarter)
		    return target.Open()
		},
		OnDestroy: func(obj lang.Object) error {
		    target := obj.(*data_5df53416.GormStarter)
		    return target.Close()
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*data_5df53416.GormStarter)
		    return _gorm_starter(target,context)
		},
    })

    return nil
}

