package mysql

//type Repositories struct {
//	UserRepo     repository.IUserRepo
//	BlogRepo     repository.IBlogRepo
//	ShopRepo     repository.IShopRepo
//	ShopTypeRepo repository.IShopTypeRepo
//	VoucherRepo  repository.IVoucherRepo
//
//	DB *gorm.DB
//}

//func NewRepositories() (*Repositories, error) {
//	// 从配置文件中获取 MySQL 相关信息
//	host := viper.GetString("mysql.host")
//	port := viper.GetString("mysql.port")
//	user := viper.GetString("mysql.user")
//	password := viper.GetString("mysql.password")
//	database := viper.GetString("mysql.database")
//
//	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, database)
//	// 初始化GORM日志配置
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Second, // Slow SQL threshold
//			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
//			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
//			Colorful:                  false,       // Disable color
//		},
//	)
//	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
//		Logger: newLogger,
//	})
//	// Error
//	if connString == "" || err != nil {
//		logger2.Logger.Error(fmt.Sprintf("mysql lost: %v", err))
//		panic(err)
//	}
//	return &Repositories{
//		UserRepo:     NewUserRepo(db),
//		BlogRepo:     NewBlogRepo(db),
//		ShopTypeRepo: NewShopTypeRepo(db),
//		ShopRepo:     NewShopRepo(db),
//		VoucherRepo:  NewVoucherRepo(db),
//		DB:           db,
//	}, nil
//}

// AutoMigrate This migrate all tables
//func (s *Repositories) AutoMigrate() error {
//	return s.DB.AutoMigrate(
//		&entity.User{},
//		&entity.ShowType{},
//		&entity.Blog{},
//		&entity.Shop{},
//		&entity.Voucher{},
//	)
//}
