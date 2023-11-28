func (m *{{.upperStartCamelObject}}) TableName() string {
	return {{.upperStartCamelObject}}TName
}
func (m *{{.upperStartCamelObject}}) Create(db *gorm.DB) error {
        // m.CreateTime = time.Now()
        // m.UpdateTime = time.Now()
    	return db.Table(m.TableName()).Create(m).Error
}

func (m *{{.upperStartCamelObject}}) FindByPrimary(db *gorm.DB,primary {{.primaryKeyFieldType}}) error {
    return IgnoreRecordNotFound(db.Table(m.TableName()).Where(" {{.primaryKeyField}} = ?",primary).Find(m).Error)
}


func (m *{{.upperStartCamelObject}}) UpdateByPrimary(db *gorm.DB, primary {{.primaryKeyFieldType}}) error {
	return db.Table(m.TableName()).Where("{{.primaryKeyField}} = ?", primary).Updates(m).Error
}

func (m *{{.upperStartCamelObject}}) UpdateFieldsByPrimary(db *gorm.DB, primary {{.primaryKeyFieldType}}, fields map[string]interface{}) error {
	return db.Table(m.TableName()).Where("{{.primaryKeyField}} = ?", primary).Updates(fields).Error
}
func (m *{{.upperStartCamelObject}}) DeleteByPrimary(db *gorm.DB, primary {{.primaryKeyFieldType}}) error {
	return db.Table(m.TableName()).Where("{{.primaryKeyField}} = ?", primary).Delete(m).Error
}

type {{.upperStartCamelObject}}List []{{.upperStartCamelObject}}

func (l*{{.upperStartCamelObject}}List) FindByPrimarys(db *gorm.DB,primarys []{{.primaryKeyFieldType}}) (err error) {
	if len(primarys) == 0 {
		return
	}
    err = db.Table({{.upperStartCamelObject}}TName).Where(" {{.primaryKeyField}} in (?)",primarys).Find(l).Error
    return
}

func (l*{{.upperStartCamelObject}}List)FindByPage(db *gorm.DB, page int, size int)(total int64, err error){
    	if page <= 0 {
    		page = 1
    	}
    	if size <= 0 {
    		size = 10
    	}
    	db = db.Table({{.upperStartCamelObject}}TName)
        //conditions
        err = db.Count(&total).Error
        if err != nil {
        	return
        }
        err = db.Offset((page - 1) * size).Limit(size).Find(&l).Error
        return
}

func (l *{{.upperStartCamelObject}}List) Create(db *gorm.DB, batchSize int) error {
	return db.CreateInBatches(l, batchSize).Error
}