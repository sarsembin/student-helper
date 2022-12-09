package commentsdb

func (c *repository) Add(e *Entity) (err error) {
	_, err = c.db.Model(e).Insert()
	return err
}

func (c *repository) GetAll(uniID int) (res []Entity, err error) {
	err = c.db.Model(&res).Where("university_id_id = ?0", uniID).Select()
	return res, err
}

func (c *repository) Get(id int) (res *Entity, err error) {
	res = &Entity{ID: id}

	err = c.db.Model(res).WherePK().Select()

	return res, err
}

func (c *repository) Put(e *Entity) (err error) {
	_, err = c.db.Model(e).WherePK().UpdateNotZero()
	return err
}

func (c *repository) Delete(id int) (err error) {
	res := &Entity{ID: id}

	_, err = c.db.Model(res).WherePK().Delete()

	return err
}
