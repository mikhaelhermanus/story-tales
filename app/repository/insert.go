package repository

import (
	"time"
)

/*Insert to database
 * @paremeter
 * i - struct to saving into database
 *
 * @return
 * uint - id after insert into database
 * error
 */
func (r *repo) Insert(table string, i interface{}) error {
	query := r.db.Table(table).Create(i)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *repo) SetRedis(key string, value interface{}, exp time.Duration) error {
	err := r.redis.Set(key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}
