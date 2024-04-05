// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`UserID`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`UserID`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsersUserIDPrefix  = "cache:users:userID:"
	cacheUsersAccountPrefix = "cache:users:account:"
	cacheUsersEmailPrefix   = "cache:users:email:"
	cacheUsersPhonePrefix   = "cache:users:phone:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, userID int64) (*Users, error)
		FindOneByAccount(ctx context.Context, account string) (*Users, error)
		FindOneByEmail(ctx context.Context, email sql.NullString) (*Users, error)
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, userID int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		UserID           int64          `db:"UserID"`
		Nickname         string         `db:"Nickname"`
		Account          string         `db:"Account"`
		Email            sql.NullString `db:"Email"`
		Phone            string         `db:"Phone"`
		Password         string         `db:"Password"`
		RegistrationTime time.Time      `db:"RegistrationTime"`
		LastLoginTime    time.Time      `db:"LastLoginTime"`
		Status           string         `db:"Status"`
		Role             string         `db:"Role"`
		BackgroundImage  sql.NullString `db:"BackgroundImage"`
		AvatarBackground sql.NullString `db:"AvatarBackground"`
		PostCount        int64          `db:"PostCount"`
		FollowCount      int64          `db:"FollowCount"`
		FansCount        int64          `db:"FansCount"`
		LikeCount        int64          `db:"LikeCount"`
		PointCount       int64          `db:"PointCount"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`Users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, userID int64) error {
	data, err := m.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersUserIDKey := fmt.Sprintf("%s%v", cacheUsersUserIDPrefix, userID)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `UserID` = ?", m.table)
		return conn.ExecCtx(ctx, query, userID)
	}, usersAccountKey, usersEmailKey, usersPhoneKey, usersUserIDKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, userID int64) (*Users, error) {
	usersUserIDKey := fmt.Sprintf("%s%v", cacheUsersUserIDPrefix, userID)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersUserIDKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `UserID` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userID)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByAccount(ctx context.Context, account string) (*Users, error) {
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, account)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersAccountKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `Account` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, account); err != nil {
			return nil, err
		}
		return resp.UserID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByEmail(ctx context.Context, email sql.NullString) (*Users, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, email)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `Email` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.UserID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, phone)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `Phone` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.UserID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersUserIDKey := fmt.Sprintf("%s%v", cacheUsersUserIDPrefix, data.UserID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Account, data.Email, data.Phone, data.Password, data.RegistrationTime, data.LastLoginTime, data.Status, data.Role, data.BackgroundImage, data.AvatarBackground, data.PostCount, data.FollowCount, data.FansCount, data.LikeCount, data.PointCount)
	}, usersAccountKey, usersEmailKey, usersPhoneKey, usersUserIDKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.UserID)
	if err != nil {
		return err
	}

	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersUserIDKey := fmt.Sprintf("%s%v", cacheUsersUserIDPrefix, data.UserID)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `UserID` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Nickname, newData.Account, newData.Email, newData.Phone, newData.Password, newData.RegistrationTime, newData.LastLoginTime, newData.Status, newData.Role, newData.BackgroundImage, newData.AvatarBackground, newData.PostCount, newData.FollowCount, newData.FansCount, newData.LikeCount, newData.PointCount, newData.UserID)
	}, usersAccountKey, usersEmailKey, usersPhoneKey, usersUserIDKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsersUserIDPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `UserID` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
