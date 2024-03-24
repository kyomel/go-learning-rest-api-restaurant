package resto

import (
	"reflect"
	"rest-api-restaurant/internal/model"
	"rest-api-restaurant/internal/repository/menu"
	"rest-api-restaurant/internal/repository/order"
	"rest-api-restaurant/internal/repository/user"
	"testing"

	"golang.org/x/net/context"
)

func Test_restoUsecase_GetMenuList(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restoUsecase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}
			got, err := r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}
