package token

import (
	"context"

	"github.com/huuhoait/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/simple-admin-common/i18n"

	"github.com/huuhoait/simple-admin-core/rpc/ent/token"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Token.Delete().Where(token.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
