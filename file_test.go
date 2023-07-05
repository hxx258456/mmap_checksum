package mmap_checksum

import (
	"reflect"
	"testing"
)

func Test_getFileSize(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "success", args: args{
			filePath: "E:\\download\\SW_DVD9_WIN_ENT_LTSC_2021_64BIT_ChnSimp_MLF_X22-84402.iso",
		}, want: 5_044_211_712, wantErr: false},
		{name: "test", args: args{
			filePath: "D:\\steam\\steamapps\\common\\鬼谷八荒\\guigubahuang_Data\\resources.assets.resS",
		}, want: 18103267008, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFileSize(tt.args.filePath)

			if (err != nil) != tt.wantErr {
				t.Errorf("getFileSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFileSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFileMMAP(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "success", args: args{
			filePath: "D:\\steam\\steamapps\\common\\鬼谷八荒\\guigubahuang_Data\\resources.assets.resS",
		}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFileMMAP(tt.args.filePath)

			if (err != nil) != tt.wantErr {
				t.Errorf("readFileMMAP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFileMMAP() got = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
