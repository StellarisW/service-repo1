package service1

import (
	"context"
	agent "github.com/StellarisW/service1/kitex_gen/agent"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddRepository(ctx context.Context, req *agent.AddRepositoryReq, callOptions ...callopt.Option) (resp *agent.AddRepositoryRes, err error) {
	resp, err = defaultClient.AddRepository(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddRepository call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteRepositories(ctx context.Context, req *agent.DeleteRepositoriesReq, callOptions ...callopt.Option) (resp *agent.DeleteRepositoriesRes, err error) {
	resp, err = defaultClient.DeleteRepositories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteRepositories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateRepository(ctx context.Context, req *agent.UpdateRepositoryReq, callOptions ...callopt.Option) (resp *agent.UpdateRepositoryRes, err error) {
	resp, err = defaultClient.UpdateRepository(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRepository call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetRepositories(ctx context.Context, req *agent.GetRepositoriesReq, callOptions ...callopt.Option) (resp *agent.GetRepositoriesRes, err error) {
	resp, err = defaultClient.GetRepositories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetRepositories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SyncRepositoryById(ctx context.Context, req *agent.SyncRepositoryByIdReq, callOptions ...callopt.Option) (resp *agent.SyncRepositoryByIdRes, err error) {
	resp, err = defaultClient.SyncRepositoryById(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SyncRepositoryById call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddIDL(ctx context.Context, req *agent.AddIDLReq, callOptions ...callopt.Option) (resp *agent.AddIDLRes, err error) {
	resp, err = defaultClient.AddIDL(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddIDL call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteIDL(ctx context.Context, req *agent.DeleteIDLsReq, callOptions ...callopt.Option) (resp *agent.DeleteIDLsRes, err error) {
	resp, err = defaultClient.DeleteIDL(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteIDL call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateIDL(ctx context.Context, req *agent.UpdateIDLReq, callOptions ...callopt.Option) (resp *agent.UpdateIDLRes, err error) {
	resp, err = defaultClient.UpdateIDL(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateIDL call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetIDLs(ctx context.Context, req *agent.GetIDLsReq, callOptions ...callopt.Option) (resp *agent.GetIDLsRes, err error) {
	resp, err = defaultClient.GetIDLs(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetIDLs call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SyncIDLsById(ctx context.Context, req *agent.SyncIDLsByIdReq, callOptions ...callopt.Option) (resp *agent.SyncIDLsByIdRes, err error) {
	resp, err = defaultClient.SyncIDLsById(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SyncIDLsById call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddTemplate(ctx context.Context, req *agent.AddTemplateReq, callOptions ...callopt.Option) (resp *agent.AddTemplateRes, err error) {
	resp, err = defaultClient.AddTemplate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddTemplate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteTemplate(ctx context.Context, req *agent.DeleteTemplateReq, callOptions ...callopt.Option) (resp *agent.DeleteTemplateRes, err error) {
	resp, err = defaultClient.DeleteTemplate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteTemplate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateTemplate(ctx context.Context, req *agent.UpdateTemplateReq, callOptions ...callopt.Option) (resp *agent.UpdateTemplateRes, err error) {
	resp, err = defaultClient.UpdateTemplate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateTemplate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetTemplates(ctx context.Context, req *agent.GetTemplatesReq, callOptions ...callopt.Option) (resp *agent.GetTemplatesRes, err error) {
	resp, err = defaultClient.GetTemplates(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetTemplates call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddTemplateItem(ctx context.Context, req *agent.AddTemplateItemReq, callOptions ...callopt.Option) (resp *agent.AddTemplateItemRes, err error) {
	resp, err = defaultClient.AddTemplateItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddTemplateItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteTemplateItem(ctx context.Context, req *agent.DeleteTemplateItemReq, callOptions ...callopt.Option) (resp *agent.DeleteTemplateItemRes, err error) {
	resp, err = defaultClient.DeleteTemplateItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteTemplateItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateTemplateItem(ctx context.Context, req *agent.UpdateTemplateItemReq, callOptions ...callopt.Option) (resp *agent.UpdateTemplateItemRes, err error) {
	resp, err = defaultClient.UpdateTemplateItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateTemplateItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetTemplateItems(ctx context.Context, req *agent.GetTemplateItemsReq, callOptions ...callopt.Option) (resp *agent.GetTemplateItemsRes, err error) {
	resp, err = defaultClient.GetTemplateItems(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetTemplateItems call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateTasks(ctx context.Context, req *agent.UpdateTasksReq, callOptions ...callopt.Option) (resp *agent.UpdateTasksRes, err error) {
	resp, err = defaultClient.UpdateTasks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateTasks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GenerateCode(ctx context.Context, req *agent.GenerateCodeReq, callOptions ...callopt.Option) (resp *agent.GenerateCodeRes, err error) {
	resp, err = defaultClient.GenerateCode(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GenerateCode call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
