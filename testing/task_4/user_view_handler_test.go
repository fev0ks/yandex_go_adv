package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		code           int
		expectedUser   User
		expectedErrMsg string
		contentType    string
	}

	tests := []struct {
		name string
		args map[string]User
		want want
	}{
		{
			name: "200",
			args: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				code: http.StatusOK,
				expectedUser: User{
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				contentType: "Application/json",
			},
		},
		{
			name: "400",
			args: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				code: http.StatusBadRequest,
				expectedUser: User{
					ID:        "",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				expectedErrMsg: "userId is empty",
				contentType:    "text/plain; charset=utf-8",
			},
		},
		{
			name: "404",
			args: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				code: http.StatusNotFound,
				expectedUser: User{
					ID:        "id2",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				expectedErrMsg: "user not found",
				contentType:    "text/plain; charset=utf-8",
			},
		},
		//{
		//	name: "500", //not possible
		//	args: map[string]User{
		//		"id1": {},
		//	},
		//	want: want{
		//		code: http.StatusInternalServerError,
		//		expectedUser: User{
		//			ID:        "id1",
		//			FirstName: "Misha",
		//			LastName:  "Popov",
		//		},
		//		expectedErrMsg: "can't provide a json. internal error",
		//		contentType: "text/plain; charset=utf-8",
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users?user_id=%s", tt.want.expectedUser.ID), nil)
			w := httptest.NewRecorder()
			h := UserViewHandler(tt.args)

			h.ServeHTTP(w, request)
			res := w.Result()
			assert.Equal(t, tt.want.code, w.Code)

			defer res.Body.Close()
			bytes, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			if w.Code != http.StatusOK {
				assert.Equal(t, tt.want.expectedErrMsg, strings.TrimSpace(string(bytes)))
			} else {
				var actualUser User
				err = json.Unmarshal(bytes, &actualUser)
				require.NoError(t, err)
				assert.Equal(t, tt.want.expectedUser, actualUser)
			}
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
