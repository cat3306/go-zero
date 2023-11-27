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

func (m *{{.upperStartCamelObject}}) FindByPrimarys(db *gorm.DB,primarys []{{.primaryKeyFieldType}}) (list[]{{.upperStartCamelObject}},err error) {
	if len(primarys) == 0 {
		return
	}
    err = db.Table(m.TableName()).Where(" {{.primaryKeyField}} in (?)",primarys).Find(&list).Error
    return
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

func (m *{{.upperStartCamelObject}}) FindByPage(db *gorm.DB, page int, size int) (list []{{.upperStartCamelObject}}, total int64, err error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	db = db.Table(m.TableName())
	//conditions
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Offset((page - 1) * size).Limit(size).Find(&list).Error
	return
}