package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	mockDB := NewMockDB(ctrl)
	mockDB.EXPECT().Get(gomock.Eq("Bob")).Return(100, errors.New("not exist"))
	mockDB.EXPECT().Get(gomock.Eq("Cindy")).Return(200, nil)
	mockDB.EXPECT().Get(gomock.Eq("David")).DoAndReturn(func(key string) (int, error) {
		return 300, nil
	})

	type args struct {
		db  DB
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_no_exist",
			args: args{
				db:  mockDB,
				key: "Bob",
			},
			want: -1,
		},
		{
			name: "test_exist01",
			args: args{
				db:  mockDB,
				key: "Cindy",
			},
			want: 200,
		},
		{
			name: "test_exist02",
			args: args{
				db:  mockDB,
				key: "David",
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromDB(tt.args.db, tt.args.key); got != tt.want {
				t.Errorf("GetFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 参数(Eq, Any, Not, Nil)
func TestMockParameters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	mockDB.EXPECT().Get(gomock.Eq("Ellen")).Return(0, errors.New("not exist"))
	mockDB.EXPECT().Get(gomock.Any()).Return(1, nil)
	mockDB.EXPECT().Get(gomock.Not("Bob")).Return(0, nil)
	mockDB.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))
}

// 返回值(Return, DoAndReturn)
func TestMockReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	mockDB.EXPECT().Get(gomock.Not("Bob")).Return(0, nil)
	mockDB.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	mockDB.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		switch key {
		case "Allen":
			return 100, nil
		case "Bob":
			return 200, nil
		default:
			return 0, errors.New("not exist")
		}
	})
}

// 调用次数(Times)
func TestMockInvokeTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	mockDB.EXPECT().Get(gomock.Any()).Return(0, nil).Times(2)
	GetFromDB(mockDB, "aaa")
	GetFromDB(mockDB, "bbb")
}

// 调用顺序(InOrder)
func TestMockInvokeOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	o1 := mockDB.EXPECT().Get(gomock.Eq("Allen")).Return(0, errors.New("not exist"))
	o2 := mockDB.EXPECT().Get(gomock.Eq("Bob")).Return(10, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(mockDB, "Allen")
	GetFromDB(mockDB, "Bob")
}
