package role

import (
	"context"

	"github.com/huuhoait/simple-admin-common/utils/pointy"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIdLogic) GetRoleById(in *core.IDReq) (*core.RoleInfo, error) {
	result, err := l.svcCtx.DB.Role.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.RoleInfo{
		Id:            &result.ID,
		CreatedAt:     pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:     pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:        pointy.GetPointer(uint32(result.Status)),
		Name:          &result.Name,
		Code:          &result.Code,
		DefaultRouter: &result.DefaultRouter,
		Remark:        &result.Remark,
		Sort:          &result.Sort,
	}, nil
}
