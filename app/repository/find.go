package repository

/*FindOne is for getting only one data
 * @method
 * struct repo with value mysql database connection
 *
 * @parameter
 * ctx - context from http.Request
 * i - struct with pointer
 * where interface - is condition query to database
 * field - show only field
 * whereValue - any value from where interface
 *
 * @return
 * error
 */
func (r *repo) FindOne(table string, i, where interface{}, field string, whereValue ...interface{}) error {
	err := r.db.Table(table).Where(where, whereValue...).Select(field).First(i).Error
	if err != nil {
		return err
	}

	return nil
}
