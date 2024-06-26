package position

import (
	"context"

	"github.com/huuhoait/simple-admin-common/utils/pointy"

	"github.com/huuhoait/simple-admin-core/rpc/ent/position"
	"github.com/huuhoait/simple-admin-core/rpc/ent/predicate"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionListLogic) GetPositionList(in *core.PositionListReq) (*core.PositionListResp, error) {
	var predicates []predicate.Position
	if in.Name != nil {
		predicates = append(predicates, position.NameContains(*in.Name))
	}
	if in.Code != nil {
		predicates = append(predicates, position.CodeContains(*in.Code))
	}
	if in.Remark != nil {
		predicates = append(predicates, position.RemarkContains(*in.Remark))
	}
	result, err := l.svcCtx.DB.Position.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.PositionListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.PositionInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Sort:      &v.Sort,
			Name:      &v.Name,
			Code:      &v.Code,
			Remark:    &v.Remark,
		})
	}

	return resp, nil
}
