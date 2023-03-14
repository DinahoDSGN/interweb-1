package postgres

import (
	"context"
	"database/sql"
	"interweb/telegram-bot-service/pkg/database"
	"interweb/telegram-bot-service/pkg/domain"
	"interweb/telegram-bot-service/pkg/logger"
	"time"
)

type Postgres struct {
	*database.Postgres
}

func NewPostgres(postgres *database.Postgres) *Postgres {
	return &Postgres{Postgres: postgres}
}

func (p *Postgres) InsertUserRequest(ctx context.Context, req domain.UserRequest) (int64, error) {
	var id int64
	err := p.Conn.QueryRowContext(ctx, "INSERT INTO user_requests(chat_id, request, result_json) VALUES($1, $2, $3) RETURNING id",
		req.ChatID, req.Request, req.Result).Scan(&id)

	return id, err
}

func (p *Postgres) GetDateFirstRequest(ctx context.Context, id int64) (datetime time.Time, err error) {
	err = p.Conn.QueryRowContext(ctx, "SELECT request_date FROM user_requests WHERE chat_id = $1 ORDER BY request_date ASC limit 1",
		id).Scan(&datetime)

	return
}

func (p *Postgres) AggregateTotalRequests(ctx context.Context, id int64) ([]domain.TotalUserRequests, error) {
	rows, err := p.Conn.QueryContext(ctx, "SELECT request, count(id) FROM user_requests WHERE chat_id = $1 GROUP BY request ORDER BY request DESC", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(rows)

	var totals []domain.TotalUserRequests
	for rows.Next() {
		var total domain.TotalUserRequests
		if err = rows.Scan(&total.Request, &total.Count); err != nil {
			logger.Error(err.Error())
			continue
		}

		totals = append(totals, total)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return totals, nil
}

func (p *Postgres) ListRequests(ctx context.Context, id int64, userRequests chan<- domain.UserRequest) error {
	defer close(userRequests)
	rows, err := p.Conn.QueryContext(ctx, "SELECT id, request, request_date, chat_id, result_json FROM user_requests WHERE chat_id = $1 ORDER BY request_date DESC", id)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(rows)

	for rows.Next() {
		var req domain.UserRequest
		if err = rows.Scan(&req.ID, &req.Request, &req.RequestDate, &req.ChatID, &req.Result); err != nil {
			logger.Error(err.Error())
			continue
		}

		userRequests <- req
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
