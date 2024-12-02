package http

import (
	"fmt"
	apimodel "houseService/internal/handler/http/model"
	"houseService/internal/model"
)

func ToHouseServiceCreateModel(h *apimodel.HouseCreateBody) *model.HouseCreate {
	return &model.HouseCreate{
		Address:   h.Address,
		Year:      h.Year,
		Developer: h.Developer,
	}
}

func ToHouseApiModel(h *model.House) *apimodel.House {
	return &apimodel.House{
		Id:        h.Id,
		Address:   h.Address,
		Year:      h.Year,
		Developer: h.Developer,
		CreatedAt: h.CreatedAt,
		UpdateAt:  h.UpdateAt,
	}
}

func ToFlatServiceModel(f *apimodel.FlatCreateBody) *model.Flat {
	return &model.Flat{
		HouseId: f.HouseId,
		Price:   f.Price,
		Rooms:   f.Rooms,
	}
}

func ToFlatUpdateServiceModel(f *apimodel.FlatUpdateBody) (*model.FlatUpdate, error) {
	s, err := ToStatusServiceModel(f.Status)
	if err != nil {
		return nil, err
	}
	return &model.FlatUpdate{
		Id:     f.Id,
		Status: s,
	}, nil
}

func ToStatusServiceModel(status apimodel.Status) (model.Status, error) {
	switch status {
	case apimodel.CREATED:
		return model.CREATED, nil
	case apimodel.APPROVED:
		return model.APPROVED, nil
	case apimodel.DECLINED:
		return model.DECLINED, nil
	case apimodel.ON_MODERATION:
		return model.ON_MODERATION, nil
	default:
		return "", fmt.Errorf("invalid status: %s", status)
	}
}

func ToFlatApiModel(f *model.Flat) *apimodel.Flat {
	return &apimodel.Flat{
		Id:      f.Id,
		HouseId: f.HouseId,
		Price:   f.Price,
		Rooms:   f.Rooms,
		Status:  apimodel.Status(f.Status),
	}
}

func ToFlatsApiModel(f []*model.Flat) *apimodel.InlineResponse2002 {
	res := make([]*apimodel.Flat, len(f))

	for i, flat := range f {
		res[i] = ToFlatApiModel(flat)
	}
	return &apimodel.InlineResponse2002{
		Flats: res,
	}
}

func ToUserServiceModel(u *apimodel.RegisterBody) (*model.User, error) {
	t, err := ToUserTypeServiceModel(u.UserType)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
		UserType: t,
	}, nil
}

func ToUserLoginServiceModel(u *apimodel.LoginBody) *model.UserLogin {
	return &model.UserLogin{
		Id:       u.Id,
		Password: u.Password,
	}
}

func ToUserTypeServiceModel(userType apimodel.UserType) (model.UserType, error) {
	switch userType {
	case apimodel.CLIENT:
		return model.CLIENT, nil
	case apimodel.MODERATOR:
		return model.MODERATOR, nil
	default:
		return "", fmt.Errorf("invalid user type: %s", userType)
	}
}

func ToUUIDApiModel(uuid string) *apimodel.InlineResponse2001 {
	return &apimodel.InlineResponse2001{
		UserId: uuid,
	}
}

func ToTokenApiModel(token string) *apimodel.InlineResponse200 {
	return &apimodel.InlineResponse200{
		Token: token,
	}
}
