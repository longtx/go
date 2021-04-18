package cmd

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:    "sql",
	Short:  "sql转换和处理",
	Long:   "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use:     "struct",
	Short:   "sql 转换",
	Long:    "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:     dbType,
			Host:       host,
			UserName:   username,
			Password:   password,
			Charset:    charset,
		}
		dbModel  := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err:%v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "192.168.10.5:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}

/**
验证方法：go run main.go sql struct --username 数据库的账号  --password 数据库的密码  --db数据库名称  --table 需要转换的表名
 */

