func (m *{{.upperStartCamelObject}}) FindBy{{.upperKeyField}}(db *gorm.DB,key {{.keyFieldType}}) error {
    return IgnoreRecordNotFound(db.Table(m.TableName()).Where(" {{.keyField}} = ?",key).Find(m).Error)
}