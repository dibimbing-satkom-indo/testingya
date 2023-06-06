package main

import (
	"testing"
	"testingya/mocks"
)

func Test_user_Talk(t *testing.T) {
	type fields struct {
		isAngry bool
		logger  Logger
	}
	type args struct {
		sentence string
	}
	mockLogger := mocks.NewLogger(t)
	mockLogger.EXPECT().
		Log("YA UDAH").
		Return(nil).
		Once()
	mockLogger.EXPECT().
		Log("makasih").
		Return(nil).
		Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "is angry",
			fields: fields{
				isAngry: true,
				logger:  mockLogger,
			},
			args: args{
				sentence: "ya udah",
			},
			wantErr: false,
		},
		{
			name: "happy",
			fields: fields{
				isAngry: false,
				logger:  mockLogger,
			},
			args: args{
				sentence: "makasih",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := user{
				isAngry: tt.fields.isAngry,
				logger:  tt.fields.logger,
			}
			if err := u.Talk(tt.args.sentence); (err != nil) != tt.wantErr {
				t.Errorf("Talk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
