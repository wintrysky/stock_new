package orm

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"ww/stocks/orm/internal"
)

var defaultKey = "default_db_key"
var db *gorm.DB
// 连接实例池
var dbInstances = struct {
	sync.RWMutex
	m map[string]*gorm.DB
}{m: make(map[string]*gorm.DB)}

func defaultDb() *gorm.DB {
	dbInstances.RLock()
	defer dbInstances.RUnlock()

	dbInst, ok := dbInstances.m[defaultKey]

	if ok == true {
		return dbInst
	} else {
		panic("没有发现DB实例" + defaultKey)
	}
}

func setDbInstance(key string, db *gorm.DB) {
	dbInstances.Lock()
	defer dbInstances.Unlock()

	if _, ok := dbInstances.m[key]; ok == false {
		dbInstances.m[key] = db
	}
}

// InitDB 注册多个数据库连接
func InitDB(conf *DbSettings) error {
	pwd, err := internal.Decrypt(conf.Password, conf.Key)
	if err != nil {
		return err
	}

	maxIdel := 5
	if conf.MaxIdleConns > 0 {
		maxIdel = conf.MaxIdleConns
	}
	maxConn := 20
	if conf.MaxOpenConns > 0 {
		maxConn = conf.MaxOpenConns
	}
	lifeTime := 5
	if conf.ConnMaxLifetime > 0 {
		lifeTime = conf.ConnMaxLifetime
	}

	p := strconv.Itoa(conf.Port)
	var source string
	source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=true&loc=Local", conf.User, pwd, conf.Host, p, conf.DbName)
	db, err = gorm.Open(mysql.Open(source), &gorm.Config{
		Logger: DefaultLogger,
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdel)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(maxConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(lifeTime))

	// 保存DB实例
	setDbInstance(defaultKey, db)

	return nil
}

type FiDB struct {
	DB               *gorm.DB // 局部
	RowsAffected int64
	Error          error
}

// FiTX 数据表操作类
type FiTX struct {
	DB             *gorm.DB // 局部
	Error          error
	rollbackCalled bool // 校验是否调用了Rollback函数
	commitCalled   bool // 校验是否调用了Commit函数
	endTransaction bool //
}

// NewQuery 获取一个查询会话（只作用于一次查询）
func NewQuery() *FiDB {
	var da = new(FiDB)
	da.DB = defaultDb()
	return da
}

// BeginTransaction 开始事务
func BeginTransaction() *FiTX {
	var tx = new(FiTX)
	tx.DB = defaultDb().Begin()
	return tx
}

// EndTransaction 结束事务
func (t *FiTX) EndTransaction() {
	if r := recover(); r != nil {
		fmt.Println("recover...", r)
		t.Error = errors.New(fmt.Sprint(r))
		t.DB.Rollback()
		panic(r)
	}

	if t.Error != nil && t.Error != gorm.ErrRecordNotFound {
		fmt.Println("rollback...", t.Error)
		t.DB.Rollback()
	} else {
		fmt.Println("commit...")
		t.DB.Commit()
	}
}

func (t *FiTX) NewQuery() *FiDB {
	var da = new(FiDB)
	da.DB = t.DB
	return da
}


// UpdateItem 更新一行数据
// entityPtr 数据库实体 引用类型
// cols 待更新的数据库字段 e.g []string{"email", "address"} "
// UPDATE `mpv_order_filter_log` SET `order_mode` = 'xxxx', `uuid` = 'xxx'  WHERE `id` = 1
// qa.UpdateItem(&item1,[]string{"uuid","order_mode"})
func (t *FiDB) UpdateItem(entityPtr interface{}, cols []string) {
	t.DB = t.DB.Model(entityPtr).Select(cols).UpdateColumns(entityPtr)
	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

// InsertItem 插入一条记录
// entityPtr 数据实体 引用类型
// qa := orm.NewQuery()
// qa.InsertItem(&user)
func (t *FiDB) InsertItem(entityPtr interface{}) {
	t.DB = t.DB.Create(entityPtr)
	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

// GetItemWhereFirst 根据条件查询一条数据
// entityPrt 待返回的数据实体 引用类型
// qa.GetItemWhereFirst(&item1,"category_id = ? and address = ?",2,"shenzhen")
func (t *FiDB) GetItemWhereFirst(entityPrt interface{}, query interface{}, args ...interface{}) {
	t.DB = t.DB.Where(query, args...).First(entityPrt)
	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

// qa.Where(1=1,"name = ? and addr = ?","xx","")
func (t *FiDB) Where(check bool,query interface{}, args ...interface{}) *FiDB {
	if check == false {
		return t
	}

	t.DB = t.DB.Where(query, args...)
	return t
}

// GetItemWhere 查询符合条件的多条记录
// SELECT * FROM `mpv_order_filter_data`  WHERE (category_id = 2)
// qa.GetItemWhere(&items,"category_id = ?",2)
func (t *FiDB) GetItemWhere(entitiesPrt interface{}, query string, args ...interface{}) {
	if query == "" {
		t.DB = t.DB.Find(entitiesPrt)
	}else{
		t.DB = t.DB.Where(query, args...).Find(entitiesPrt)
	}

	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

// BatchInsert 批量插入
// entities 待更新的实体数组slice, ID不能为空
// 例子：var users []model.User
// qa.BatchInsert(users)
func (t *FiDB) BatchInsert(entities interface{})  {
	t.DB = t.DB.CreateInBatches(entities, 1000)
	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

func (t *FiDB)UpdateOrInserts(entitiesPrt interface{},cols []string){
	// 在`id`冲突时，将列更新为新值
	if len(cols) >0 {
		t.DB = t.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "symbol"}},
			DoUpdates: clause.AssignmentColumns(cols),
		}).Create(entitiesPrt)
	}else{
		t.DB = t.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "symbol"}},
			UpdateAll: true,
		}).Create(entitiesPrt)
	}


	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

func (t *FiDB)UpdateOrInsertHistory(entitiesPrt interface{},cols []string){
	// 在`id`冲突时，将列更新为新值
	var keys []clause.Column
	var c1 clause.Column
	c1.Name = "symbol"
	var c2 clause.Column
	c2.Table = "trade_time"
	keys = append(keys,c1)
	keys = append(keys,c2)

	if len(cols) >0 {
		t.DB = t.DB.Clauses(clause.OnConflict{
			Columns:   keys,
			DoUpdates: clause.AssignmentColumns(cols),
		}).Create(entitiesPrt)
	}else{
		t.DB = t.DB.Clauses(clause.OnConflict{
			Columns:   keys,
			UpdateAll: true,
		}).Create(entitiesPrt)
	}


	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

// ExecuteTextQuery 原生语法查询
// entitiesPrt 实体数组 引用类型
// sql 原生SQL
// values SQL占位符中的变量值
// 例子：var userList []model.UserView
// da.ExecuteTextQuery(&userList,
//    "SELECT u.id,u.name, dept_name FROM department t ,user u WHERE u.dept_id=t.id AND u.ID >=?",12)
func (t *FiDB) ExecuteTextQuery(entitiesPrt interface{}, sql string, values ...interface{}) {
	d := t.DB.Raw(sql, values...)
	if d.Error != nil && d.Error != gorm.ErrRecordNotFound {
		t.Error = d.Error
	}

	t.DB = d.Scan(entitiesPrt)
	if t.DB.Error != nil && t.DB.Error != gorm.ErrRecordNotFound {
		t.Error = t.DB.Error
	}

	t.RowsAffected = t.DB.RowsAffected
}

// Exec execute raw sql
// sql := "update mpv_order_filter_log set order_mode='xxx' where id =3"
//// update mpv_order_filter_log set order_mode='xxx' where id =3
// qa.Exec(sql)
func (t *FiDB) Exec(sql string, values ...interface{}) {
	t.DB = t.DB.Exec(sql, values...)

	if t.DB.Error != nil && t.DB.Error != gorm.ErrRecordNotFound {
		t.Error = t.DB.Error
	}

	t.RowsAffected = t.DB.RowsAffected
}

// 删除多条数据
// 必须要有where语句
// DELETE FROM `mpv_order_filter_log`  WHERE (category_id = 33)
// qa.Where("category_id = ?",33).DeleteItem(items)
func (t *FiDB) DeleteItem(entityPtrOrSlice interface{}) {
	t.DB = t.DB.Delete(entityPtrOrSlice)
	if t.DB.Error != nil {
		t.Error = t.DB.Error
	}
	t.RowsAffected = t.DB.RowsAffected
}

// db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
func (t *FiDB) UpdateTradeRate(symbol string, cols map[string]interface{}) {
	t.DB = t.DB.Table("stock_basic").Where("symbol = ?",symbol).
		Updates(cols)
	t.Error = t.DB.Error
	t.RowsAffected = t.DB.RowsAffected
}

func(t *FiDB)DeleteAll(sql string){
	t.DB = t.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec(sql)
	if t.DB.Error != nil {
		t.Error = t.DB.Error
	}
	t.RowsAffected = t.DB.RowsAffected
}


//qa.BatchUpdate(items,[]string{"uuid","filter_mode"})
func (t *FiTX) BatchUpdate(rows interface{}, cols []string) (err error) {
	// 获取表结构信息 DryRun模式
	stmt := t.DB.Session(&gorm.Session{DryRun: true}).First(&rows, 1).Statement
	if stmt.Schema == nil {
		return errors.New("没有发现表结构信息:"+stmt.Table)
	}

	// 判断主键是否为ID
	primaryKey, err := t.checkPrimaryKey(stmt)
	if err != nil {
		return
	}

	// 获取结构名称和数据库字段名称对应关系
	fieldMap := getTableFieldMap(stmt) // map[dept_name]=DeptName
	// 获取待更新字段的映射关系
	colMap := make(map[string]string) // e.g map[DeptName]=dept_name
	for _,col := range cols {
		if v,ok := fieldMap[col];ok {
			colMap[v] = v
		}
	}

	// 判断cols字段是否存在
	err = checkColumnExists(stmt,cols)
	if err != nil {
		return
	}

	// 分析结构，构建带参SQL
	raws := reflect.ValueOf(rows)
	total := raws.Len()
	if total == 0 {
		return
	}

	var sql string
	batchSize := 1000
	if batchSize <= total {
		sql,err = t.generateUpdateSQL(stmt, cols, primaryKey,batchSize)
		if err != nil {
			return
		}
	}

	var subRows []interface{}
	for i := 0; i < total; i++ {
		subRows = append(subRows, raws.Index(i).Interface())
		if len(subRows) == batchSize {
			t.batchUpdateInternal(subRows,cols,sql,primaryKey,batchSize,fieldMap,colMap)
			subRows = nil
		}
	}
	if subRows != nil && len(subRows) > 0 {
		sql,err = t.generateUpdateSQL(stmt, cols, primaryKey,len(subRows))
		if err != nil {
			return
		}
		t.batchUpdateInternal(subRows,cols,sql,primaryKey,len(subRows),fieldMap,colMap)
	}
	return
}


func (t *FiTX) batchUpdateInternal(subRows interface{}, cols []string,sql string,
	primaryKey string,total int,fieldMap map[string]string,colMap map[string]string) (err error) {
	raws := reflect.ValueOf(subRows)

	// 将传入的行转换为数组
	var rowsMap []map[string]interface{}
	var ids []interface{}
	for i:=0;i<total;i++{
		val := raws.Index(i).Interface()
		m := make(map[string]interface{})

		vv := reflect.ValueOf(val)
		tp := reflect.Indirect(vv).Type()

		for i := 0; i < vv.NumField(); i++ {
			value := vv.Field(i).Interface()
			fieldName := tp.Field(i).Name
			_,ok := colMap[fieldName]
			if ok || strings.ToLower(fieldName) == "id" {
				m[fieldName] = value
			}
		}

		rowsMap = append(rowsMap,m)
		// 保存ID集合
		id,_ := getValueByName(m,primaryKey,fieldMap)
		ids = append(ids,id)
	}

	// 分批获取变量数组
	var values []interface{}
	for _,col := range cols {
		for _, rowMap := range rowsMap {
			// 添加ID值
			id,err := getValueByName(rowMap,primaryKey,fieldMap)
			if err != nil {
				return err
			}
			values = append(values,id)

			// 添加ID对应的更新字段值
			v, err := getValueByName(rowMap,col,fieldMap)
			values = append(values,v)
		}
	}
	values = append(values,ids...)

	//分批执行预编译语句
	t.DB = t.DB.Exec(sql,values...)
	if t.DB.Error != nil {
		t.Error = t.DB.Error
	}
	return
}


func getValueByName(val map[string]interface{},col string,fieldMap map[string]string)(data interface{},err error){
	colName,_ := fieldMap[col]
	v,ok := val[colName]
	if !ok {
		return "",errors.New("没有发现字段:"+col)
	}
	return v, nil
}

// 判断更新字段是否存在
func checkColumnExists(stmt *gorm.Statement,cols []string) (err error) {
	cm := make(map[string]string)
	for _, field := range stmt.Schema.Fields {
		colName := strings.ToLower(field.DBName)
		cm[colName] = field.DBName
	}

	for _,c := range cols {
		key := strings.ToLower(c)
		if _,ok := cm[key]; !ok {
			return errors.New("不存在字段:"+c)
		}
	}

	return
}

// 获取结构名称和数据库字段名称对应关系
func getTableFieldMap(stmt *gorm.Statement)(fieldMap map[string]string){
	fieldMap = make(map[string]string)
	for _, field := range stmt.Schema.Fields {
		fieldMap[field.DBName] = field.Name
	}
	return
}

// 判断主键是否为ID
func (t *FiTX) checkPrimaryKey(stmt *gorm.Statement) (id string, err error) {

	// 获取主键集合
	var primaryKeys []string
	var isID bool
	for _, field := range stmt.Schema.Fields {
		if field.PrimaryKey {
			primaryKeys = append(primaryKeys,field.DBName)
			if strings.ToLower(field.Name) == "id" {
				isID = true
			}
		}
	}
	if len(primaryKeys) == 0 {
		return "",errors.New("没有发现主键,表:"+stmt.Table)
	}
	if len(primaryKeys) == 1 {
		return primaryKeys[0], nil
	}
	if isID == true {
		return "id", nil
	}else{
		return "",errors.New("批量更新不支持复合主键,表:"+stmt.Table)
	}
}

func (t *FiTX) generateUpdateSQL(stmt *gorm.Statement, cols []string, primaryKey string,
	rowCount int) (sql string,err error) {
	sql = "UPDATE " + stmt.Table + " SET "

	for _, col := range cols {
		line := " `" + col + "` = CASE " + primaryKey + " "
		// 根据需要更新的行数构建参数SQL
		for i := 0; i < rowCount; i++ {
			line += "when ? then ? "
		}
		line += "END,"
		sql += line
	}
	sql = strings.TrimSuffix(sql, ",")

	ids := " WHERE " + primaryKey + " IN ("
	for i := 0; i < rowCount; i++ {
		ids += "?,"
	}
	ids = strings.TrimSuffix(ids, ",")
	sql += ids + ")"

	return
}
