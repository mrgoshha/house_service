package postrges

import (
	"houseService/internal/adapter/dbs/postrges/entity"
	"houseService/internal/model"
)

func ToFlatServiceModel(f *entity.Flat) *model.Flat {
	return &model.Flat{
		Id:      f.Id,
		HouseId: f.HouseId,
		Price:   f.Price,
		Rooms:   f.Rooms,
		Status:  f.Status,
	}
}

func ToHouseServiceModel(h *entity.House) *model.House {
	return &model.House{
		Id:        h.Id,
		Address:   h.Address,
		Year:      h.Year,
		Developer: h.Developer,
		CreatedAt: h.CreatedAt,
		UpdateAt:  h.UpdateAt,
	}
}

func ToUserServiceModel(u *entity.User) *model.User {
	return &model.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		UserType: u.UserType,
	}
}

func ToFlatsServiceModel(f []*entity.Flat) []*model.Flat {
	res := make([]*model.Flat, len(f))

	for i, flat := range f {
		res[i] = ToFlatServiceModel(flat)
	}
	return res
}
