package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/models"
	_ "github.com/jackc/pgx/stdlib"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, params string) error {
	db, err := sql.Open("pgx", params)
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	s.db = db

	logger.UseLogger().Info("service connected to database")

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) CreateEvent(ctx context.Context, event models.Event) (int, error) {
	args := []interface{}{event.Title, event.DateTimeStart, event.DateTimeEnd, event.Description, event.UserId, event.NotificationDuration}
	query := `INSERT INTO event (title, date_time_start, date_time_end, description, user_id, notification_duration)
			  VALUES($1, $2, $3, $4, $5, $6)
			  RETURNING id`

	var id int
	err := s.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) UpdateEvent(ctx context.Context, id int, event models.Event) error {
	args := []interface{}{event.Title, event.DateTimeStart, event.DateTimeEnd, event.Description, event.UserId, event.NotificationDuration, id}
	query := `UPDATE event
			  SET title = $1,
				  date_time_start = $2,
				  date_time_end = $3,
				  description = $4,
				  user_id = $5,
				  notification_duration = $6
			  WHERE id = $7;`

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, id int) error {
	query := `DELETE FROM event	WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteEventsOldThenLastYear(ctx context.Context) error {
	query := `DELETE FROM event	WHERE date_time_end < date_trunc('minute', now() - interval '1 year')`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetEvents(ctx context.Context) ([]models.Event, error) {
	query := `SELECT id, title, date_time_start, date_time_end, description, user_id, notification_duration
			  FROM event
		      ORDER BY date_time_start`
	return s.prepareEvents(ctx, query)
}

func (s *Storage) GetEventsByLastDay(ctx context.Context, date time.Time) ([]models.Event, error) {
	query := `SELECT id, title, date_time_start, date_time_end, description, user_id, notification_duration
			  FROM event
			  WHERE date_time_end > date_trunc('minute', now() - interval '1 day')
		      ORDER BY date_time_start`
	return s.prepareEvents(ctx, query)
}

func (s *Storage) GetEventsByLastWeek(ctx context.Context, date time.Time) ([]models.Event, error) {
	query := `SELECT id, title, date_time_start, date_time_end, description, user_id, notification_duration
			  FROM event
			  WHERE date_time_end > date_trunc('minute', now() - interval '1 week')
		      ORDER BY date_time_start`
	return s.prepareEvents(ctx, query)
}

func (s *Storage) GetEventsByLastMonth(ctx context.Context, date time.Time) ([]models.Event, error) {
	query := `SELECT id, title, date_time_start, date_time_end, description, user_id, notification_duration
			  FROM event
			  WHERE date_time_end > date_trunc('minute', now() - interval '1 month')
		      ORDER BY date_time_start`
	return s.prepareEvents(ctx, query)
}

func (s *Storage) prepareEvents(ctx context.Context, query string) ([]models.Event, error) {
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]models.Event, 0)
	for rows.Next() {
		var event models.Event
		var notification sql.NullInt64
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.Description,
			&event.UserId,
			&notification,
		)
		if err != nil {
			return nil, err
		}

		if notification.Valid {
			event.NotificationDuration = (*time.Duration)(&notification.Int64)
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
