package context_topick

import (
	"context"
	"fmt"
	"time"
)

const (
	waitDur    = 1 * time.Second
	cancelDur  = 2000 * time.Millisecond
	timeoutDur = 500 * time.Millisecond
)

type DB struct {
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email string) (User, error) {
	ctx2, cancel := context.WithDeadline(ctx, time.Now().Add(timeoutDur)) /* 1 — допишите здесь создание контекста с тайм-аутом */
	defer cancel()
	timer := time.NewTimer(waitDur)
	select {
	case <-timer.C:
		return User{Name: "Gosha"}, nil
	case <-ctx2.Done(): // 1 - допишите получение сигнала отмены контекста
		return User{}, ctx2.Err()
	}
}

type Handler struct {
	db *DB
}

type Request struct {
	Email string
}

type Response struct {
	User User
}

func (h *Handler) HandleAPI(ctx context.Context, req Request) (Response, error) {
	u, err := h.db.SelectUser(ctx, req.Email)
	if err != nil {
		return Response{}, err
	}

	return Response{User: u}, nil
}

func CtxTask1() {
	db := DB{}
	handler := Handler{db: &db}
	ctx, cancel := context.WithCancel(context.Background())
	//go func() {
	//	timer := time.NewTimer(time.Millisecond * 1500)
	//	<-timer.C
	//	cancel()
	//}()
	time.AfterFunc(cancelDur, cancel)
	// 2 - допишите код, который отменяет контекст через 500 миллисекунд

	// когда запустите код и он отработает успешно,
	// попробуйте заменить длительность на 2000 миллисекунд

	req := Request{Email: "test@yandex.ru"}
	resp, err := handler.HandleAPI(ctx, req)
	fmt.Println(resp, err)
}
