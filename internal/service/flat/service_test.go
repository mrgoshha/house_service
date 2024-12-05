package flat

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repomocks "houseService/internal/adapter/dbs/mocks"
	"houseService/internal/model"
	servicemocks "houseService/internal/service/mocks"
	"testing"
)

func TestService_FlatCreate(t *testing.T) {
	// Init Test Table
	type mockBehaviorRepo func(r *repomocks.MockFlatRepository, flat *model.Flat)
	type mockBehaviorHouseServiceGet func(r *servicemocks.MockHouse, id int)
	type mockBehaviorHouseUpdate func(r *servicemocks.MockHouse, house *model.House)

	tests := []struct {
		name                        string
		inputFlat                   *model.Flat
		inputId                     int
		inputHouse                  *model.House
		mockBehaviorRepo            mockBehaviorRepo
		mockBehaviorHouseServiceGet mockBehaviorHouseServiceGet
		mockBehaviorHouseUpdate     mockBehaviorHouseUpdate
		expectedFlat                *model.Flat
		expectedError               error
	}{
		{
			name: "Ok",
			inputFlat: &model.Flat{
				HouseId: 1,
				Price:   20,
				Rooms:   2,
			},
			inputId: 1,
			inputHouse: &model.House{
				Id: 1,
			},
			mockBehaviorRepo: func(r *repomocks.MockFlatRepository, flat *model.Flat) {
				newFlat := &model.Flat{
					HouseId: 1,
					Price:   20,
					Rooms:   2,
					Status:  model.CREATED,
				}
				saveFlat := &model.Flat{
					Id:      1,
					HouseId: 1,
					Price:   20,
					Rooms:   2,
					Status:  model.CREATED,
				}
				r.EXPECT().FlatCreate(newFlat).Return(saveFlat, nil)
			},
			mockBehaviorHouseServiceGet: func(r *servicemocks.MockHouse, id int) {
				house := &model.House{
					Id: 1,
				}
				r.EXPECT().HouseGetById(1).Return(house, nil)
			},
			mockBehaviorHouseUpdate: func(r *servicemocks.MockHouse, h *model.House) {
				r.EXPECT().HouseUpdate(gomock.Any()).Return(nil)
			},
			expectedFlat: &model.Flat{
				Id:      1,
				HouseId: 1,
				Price:   20,
				Rooms:   2,
				Status:  model.CREATED,
			},
			expectedError: nil,
		},
		{
			name:                        "Invalid Flat",
			inputFlat:                   &model.Flat{},
			mockBehaviorRepo:            func(r *repomocks.MockFlatRepository, flat *model.Flat) {},
			mockBehaviorHouseServiceGet: func(r *servicemocks.MockHouse, id int) {},
			mockBehaviorHouseUpdate:     func(r *servicemocks.MockHouse, h *model.House) {},
			expectedFlat:                nil,
			expectedError:               errors.New("validate flat data is not valid"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := repomocks.NewMockFlatRepository(c)
			houseService := servicemocks.NewMockHouse(c)

			test.mockBehaviorRepo(repo, test.inputFlat)
			test.mockBehaviorHouseServiceGet(houseService, test.inputId)
			test.mockBehaviorHouseUpdate(houseService, test.inputHouse)

			service := &Service{
				repository:   repo,
				houseService: houseService,
			}

			// Call method
			create, err := service.FlatCreate(test.inputFlat)

			// Assert
			assert.Equal(t, create, test.expectedFlat)
			if test.expectedError == nil {
				assert.Equal(t, create, test.expectedFlat)
			} else {
				assert.Equal(t, test.expectedError.Error(), err.Error())
			}

		})
	}
}

func TestService_HouseFlatsGet(t *testing.T) {
	// Init Test Table
	type mockBehaviorRepo func(r *repomocks.MockFlatRepository, id int, status string)

	tests := []struct {
		name             string
		inputId          int
		inputUserType    string
		inputFlatType    string
		mockBehaviorRepo mockBehaviorRepo
		expectedFlats    []*model.Flat
		expectedError    error
	}{
		{
			name:          "Get approved flats",
			inputId:       1,
			inputUserType: "client",
			inputFlatType: "approved",
			mockBehaviorRepo: func(r *repomocks.MockFlatRepository, id int, status string) {
				flats := []*model.Flat{
					{
						Id:      1,
						HouseId: 1,
						Price:   20,
						Rooms:   2,
						Status:  model.APPROVED,
					},
				}
				r.EXPECT().HouseFlatsGet(1, "approved").Return(flats, nil)
			},
			expectedFlats: []*model.Flat{
				{
					Id:      1,
					HouseId: 1,
					Price:   20,
					Rooms:   2,
					Status:  model.APPROVED,
				},
			},
			expectedError: nil,
		},
		{
			name:          "Get all flats",
			inputId:       1,
			inputUserType: "moderator",
			inputFlatType: "",
			mockBehaviorRepo: func(r *repomocks.MockFlatRepository, id int, status string) {
				flats := []*model.Flat{
					{
						Id:      1,
						HouseId: 1,
						Price:   20,
						Rooms:   2,
						Status:  model.ON_MODERATION,
					},
				}
				r.EXPECT().HouseFlatsGet(1, "").Return(flats, nil)
			},
			expectedFlats: []*model.Flat{
				{
					Id:      1,
					HouseId: 1,
					Price:   20,
					Rooms:   2,
					Status:  model.ON_MODERATION,
				},
			},
			expectedError: nil,
		},
		{
			name:          "Storage error",
			inputId:       1,
			inputUserType: "moderator",
			inputFlatType: "",
			mockBehaviorRepo: func(r *repomocks.MockFlatRepository, id int, status string) {
				r.EXPECT().HouseFlatsGet(1, "").Return(nil, errors.New("storage error"))
			},
			expectedFlats: nil,
			expectedError: errors.New("storage error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := repomocks.NewMockFlatRepository(c)
			houseService := servicemocks.NewMockHouse(c)

			test.mockBehaviorRepo(repo, test.inputId, test.inputFlatType)

			service := &Service{
				repository:   repo,
				houseService: houseService,
			}

			// Call method
			flats, err := service.HouseFlatsGet(test.inputId, test.inputUserType)

			// Assert
			assert.Equal(t, flats, test.expectedFlats)
			if test.expectedError == nil {
				assert.Equal(t, err, test.expectedError)
			} else {
				assert.Equal(t, test.expectedError.Error(), err.Error())
			}

		})
	}
}
