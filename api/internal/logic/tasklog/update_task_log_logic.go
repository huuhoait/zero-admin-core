package tasklog

import (
	"context"

	"github.com/huuhoait/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/huuhoait/simple-admin-core/api/internal/svc"
	"github.com/huuhoait/simple-admin-core/api/internal/types"
	"github.com/huuhoait/simple-admin-job/types/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogLogic {
	return &UpdateTaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTaskLogLogic) UpdateTaskLog(req *types.TaskLogInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.JobRpc.UpdateTaskLog(l.ctx,
		&job.TaskLogInfo{
			Id:         req.Id,
			StartedAt:  req.StartedAt,
			FinishedAt: req.FinishedAt,
			Result:     req.Result,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
