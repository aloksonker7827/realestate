package impl

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	mockimpl "gitlab.com/umi7/DezHab_user_backend/dao/core/mock-impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"testing"
)

func TestCreate(t *testing.T) {
	testCase := []struct {
		name         string
		apiClient    models.ApiClient
		mockDBCreate func(object interface{}) error
		err          error
	}{
		{
			name:      "error",
			apiClient: models.ApiClient{},
			mockDBCreate: func(object interface{}) error {
				return errors.New("dummy error")
			},
			err: errors.New("dummy error"),
		},
		{
			name:      "success",
			apiClient: models.ApiClient{},
			mockDBCreate: func(object interface{}) error {
				return nil
			},
			err: nil,
		},
	}

	writer := impl.Write
	mockCtrl := gomock.NewController(t)
	defer func() {
		impl.Write = writer
		mockCtrl.Finish()
	}()

	mockWriter := mockimpl.NewMockWriter(mockCtrl)
	impl.Write = mockWriter

	//t.Parallel()
	ctx := context.Background()
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockWriter.EXPECT().Create(gomock.Any()).Return(tc.mockDBCreate(tc.apiClient))
			_, err := ApiClient.Create(ctx, tc.apiClient)
			assert.Equal(t, tc.err, err)
		})
	}
}

//TODO : find a way which modifies the parameters of the function in the runtime
// using mock. All mock packages provides the functionality to return whatever we want.
// But, there is no way to modify the parameters of the function using mock
// Either, we have to write our custom mock implementation in this case or skip testing the case.
func TestSearch(t *testing.T) {
	testCase := []struct {
		name        string
		apiKey      string
		mockDBFirst func(out interface{}, filters interface{}) error
		apiClient   models.ApiClient
		err         error
	}{
		{
			name:   "error-on-search",
			apiKey: "test",
			mockDBFirst: func(out interface{}, filters interface{}) error {
				return errors.New("dummy error")
			},
			apiClient: models.ApiClient{},
			err:       errors.New("dummy error"),
		},
		{
			name:   "error-on-length",
			apiKey: "test",
			mockDBFirst: func(out interface{}, filters interface{}) error {
				return errors.New(ErrNoData)
			},
			apiClient: models.ApiClient{},
			err:       errors.New(ErrNoData),
		},
		{ // update this test case once the above TODO is resolved
			name:   "success",
			apiKey: "test",
			mockDBFirst: func(out interface{}, filters interface{}) (err error) {
				return
			},
			apiClient: models.ApiClient{},
			err:       errors.New(ErrNoData),
		},
	}

	reader := impl.Read
	mockCtrl := gomock.NewController(t)
	defer func() {
		impl.Read = reader
		mockCtrl.Finish()
	}()

	mockReader := mockimpl.NewMockReader(mockCtrl)
	impl.Read = mockReader

	//t.Parallel()
	ctx := context.Background()
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockReader.EXPECT().First(gomock.Any(), gomock.Any()).Return(tc.mockDBFirst(gomock.Any(), gomock.Any()))
			apiClient, err := ApiClient.Search(ctx, tc.apiKey)
			assert.Equal(t, tc.apiClient, apiClient)
			assert.Equal(t, tc.err, err)
		})
	}
}
