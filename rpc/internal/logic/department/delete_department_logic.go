package department

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/huuhoait/simple-admin-core/rpc/ent/department"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/simple-admin-common/i18n"
)

type DeleteDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(in *core.IDsReq) (*core.BaseResp, error) {
	exist, err := l.svcCtx.DB.Department.Query().Where(department.ParentIDIn(in.Ids...)).Exist(l.ctx)
	if exist {
		logx.Errorw("delete department failed, please check its children had been deleted",
			logx.Field("departmentId", in.Ids))
		return nil, errorx.NewInvalidArgumentError("department.deleteDepartmentChildrenFirst")
	}

	_, err = l.svcCtx.DB.Department.Delete().Where(department.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
