package repository

/*Update to database
 * @paremeter
 * i - struct to saving into database
 *
 * @return
 * uint - id after insert into database
 * error
 */
func (r *repo) Update(i interface{}, data map[string]interface{}) error {
	query := r.db.Model(i).Updates(data)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
