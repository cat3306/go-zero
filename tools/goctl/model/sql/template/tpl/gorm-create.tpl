func (m *{{.upperStartCamelObject}}) Create(db *gorm.DB) error {
        // m.CreateTime = time.Now()
        // m.UpdateTime = time.Now()
    	return db.Table(m.TableName()).Create(m).Error
}