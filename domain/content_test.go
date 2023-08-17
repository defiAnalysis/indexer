package domain_test

import (
	"inscription/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoder(t *testing.T) {
	tt := []struct {
		input  string
		want   domain.Inscription
		hasErr bool
	}{
		{
			input: "data:text/plain,message",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:,message",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:,",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte(""),
			},
			hasErr: false,
		},
		{
			input: "data:text/plain;charset=UTF-8;random_key=random_value,message",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:text/plain;base64,bWVzc2FnZQ",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:text/plain;base64,bWVzc2FnZQ=",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:text/plain;base64,bWVzc2FnZQ==",
			want: domain.Inscription{
				ContentType: "text/plain",
				Data:        []byte("message"),
			},
			hasErr: false,
		},
		{
			input: "data:image/jpeg;base64,/9j/2wBDAAMCAgICAgMCAgIDAwMDBAYEBAQEBAgGBgUGCQgKCgkICQkKDA8MCgsOCwkJDRENDg8QEBEQCgwSExIQEw8QEBD/yQALCAABAAEBAREA/8wABgAQEAX/2gAIAQEAAD8A0s8g/9k=",
			want: domain.Inscription{
				ContentType: "image/jpeg",
				Data:        []byte("\xff\xd8\xff\xdb\x00\x43\x00\x03\x02\x02\x02\x02\x02\x03\x02\x02\x02\x03\x03\x03\x03\x04\x06\x04\x04\x04\x04\x04\b\x06\x06\x05\x06\t\b\n\n\t\b\t\t\n\f\x0f\f\n\x0b\x0e\x0b\t\t\r\x11\r\x0e\x0f\x10\x10\x11\x10\n\f\x12\x13\x12\x10\x13\x0f\x10\x10\x10\xff\xc9\x00\x0b\b\x00\x01\x00\x01\x01\x01\x11\x00\xff\xcc\x00\x06\x00\x10\x10\x05\xff\xda\x00\b\x01\x01\x00\x00\x3f\x00\xd2\xcf\x20\xff\xd9"),
			},
			hasErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			got, err := domain.DecodeContentURI(tc.input)
			if tc.hasErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Equal(t, tc.want.ContentType, got.ContentType)
			require.Equal(t, tc.want.Data, got.Data)
		})
	}
}

func FuzzDecoder(f *testing.F) {
	testcases := []string{
		"data:text/plain,message",
		"data:,message",
		"data:text/plain;charset=UTF-8;random_key=random_value,message",
		"data:text/plain;base64,bWVzc2FnZQ==",
		"data:text/plain;base64,hi",
		`data:,{"p":"brc-20","op":"mint","ticket":"t"}`,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		//check for panics only
		domain.DecodeContentURI(orig)
	})
}
