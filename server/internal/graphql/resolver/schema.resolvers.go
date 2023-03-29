package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"fmt"
	"time"

	"github.com/layer5io/meshery/server/internal/graphql/generated"
	"github.com/layer5io/meshery/server/internal/graphql/model"
	"github.com/layer5io/meshery/server/models"
	"github.com/layer5io/meshkit/broker"
	"github.com/layer5io/meshkit/models/controllers"
)

// ChangeOperatorStatus is the resolver for the changeOperatorStatus field.
func (r *mutationResolver) ChangeOperatorStatus(ctx context.Context, input *model.OperatorStatusInput) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.changeOperatorStatus(ctx, provider, input.TargetStatus, input.ContextID)
}

// ChangeAdaptorStatus is the resolver for the changeAdaptorStatus field.
func (r *mutationResolver) ChangeAdaptorStatus(ctx context.Context, input *model.AdaptorStatusInput) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.changeAdaptorStatus(ctx, provider, input.TargetStatus, input.AdapterPort)
}

// GetAvailableAddons is the resolver for the getAvailableAddons field.
func (r *queryResolver) GetAvailableAddons(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.AddonList, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getAvailableAddons(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetControlPlanes is the resolver for the getControlPlanes field.
func (r *queryResolver) GetControlPlanes(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.ControlPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getControlPlanes(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetDataPlanes is the resolver for the getDataPlanes field.
func (r *queryResolver) GetDataPlanes(ctx context.Context, filter *model.ServiceMeshFilter) ([]*model.DataPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.getDataPlanes(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// GetOperatorStatus is the resolver for the getOperatorStatus field.
func (r *queryResolver) GetOperatorStatus(ctx context.Context, k8scontextID string) (*model.OperatorStatus, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getOperatorStatus(ctx, provider, k8scontextID)
}

// ResyncCluster is the resolver for the resyncCluster field.
func (r *queryResolver) ResyncCluster(ctx context.Context, selector *model.ReSyncActions, k8scontextID string) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.resyncCluster(ctx, provider, selector, k8scontextID)
}

// GetMeshsyncStatus is the resolver for the getMeshsyncStatus field.
func (r *queryResolver) GetMeshsyncStatus(ctx context.Context, k8scontextID string) (*model.OperatorControllerStatus, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getMeshsyncStatus(ctx, provider, k8scontextID)
}

// DeployMeshsync is the resolver for the deployMeshsync field.
func (r *queryResolver) DeployMeshsync(ctx context.Context, k8scontextID string) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.deployMeshsync(ctx, provider, k8scontextID)
}

// GetNatsStatus is the resolver for the getNatsStatus field.
func (r *queryResolver) GetNatsStatus(ctx context.Context, k8scontextID string) (*model.OperatorControllerStatus, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getNatsStatus(ctx, provider, k8scontextID)
}

// ConnectToNats is the resolver for the connectToNats field.
func (r *queryResolver) ConnectToNats(ctx context.Context, k8scontextID string) (model.Status, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.connectToNats(ctx, provider, k8scontextID)
}

// GetAvailableNamespaces is the resolver for the getAvailableNamespaces field.
func (r *queryResolver) GetAvailableNamespaces(ctx context.Context, k8sClusterIDs []string) ([]*model.NameSpace, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getAvailableNamespaces(ctx, provider, k8sClusterIDs)
}

// GetPerfResult is the resolver for the getPerfResult field.
func (r *queryResolver) GetPerfResult(ctx context.Context, id string) (*model.MesheryResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getPerfResult(ctx, provider, id)
	// panic(fmt.Errorf("not implemented"))
}

// FetchResults is the resolver for the fetchResults field.
func (r *queryResolver) FetchResults(ctx context.Context, selector model.PageFilter, profileID string) (*model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchResults(ctx, provider, selector, profileID)
}

// GetPerformanceProfiles is the resolver for the getPerformanceProfiles field.
func (r *queryResolver) GetPerformanceProfiles(ctx context.Context, selector model.PageFilter) (*model.PerfPageProfiles, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getPerformanceProfiles(ctx, provider, selector)
}

// FetchAllResults is the resolver for the fetchAllResults field.
func (r *queryResolver) FetchAllResults(ctx context.Context, selector model.PageFilter) (*model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchAllResults(ctx, provider, selector)
}

// FetchPatterns is the resolver for the fetchPatterns field.
func (r *queryResolver) FetchPatterns(ctx context.Context, selector model.PageFilter) (*model.PatternPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchPatterns(ctx, provider, selector)
}

// GetWorkloads is the resolver for the getWorkloads field.
func (r *queryResolver) GetWorkloads(ctx context.Context, name *string, id *string, trim *bool) ([]*model.OAMCapability, error) {
	return r.getWorkloads(ctx, name, id, trim)
}

// GetTraits is the resolver for the getTraits field.
func (r *queryResolver) GetTraits(ctx context.Context, name *string, id *string, trim *bool) ([]*model.OAMCapability, error) {
	return r.getTraits(ctx, name, id, trim)
}

// GetScopes is the resolver for the getScopes field.
func (r *queryResolver) GetScopes(ctx context.Context, name *string, id *string, trim *bool) ([]*model.OAMCapability, error) {
	return r.getScopes(ctx, name, id, trim)
}

// GetKubectlDescribe is the resolver for the getKubectlDescribe field.
func (r *queryResolver) GetKubectlDescribe(ctx context.Context, name string, kind string, namespace string) (*model.KctlDescribeDetails, error) {
	return r.getKubectlDescribe(ctx, name, kind, namespace)
}

// FetchPatternCatalogContent is the resolver for the fetchPatternCatalogContent field.
func (r *queryResolver) FetchPatternCatalogContent(ctx context.Context, selector *model.CatalogSelector) ([]*model.CatalogPattern, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchCatalogPattern(ctx, provider, selector)
}

// FetchFilterCatalogContent is the resolver for the fetchFilterCatalogContent field.
func (r *queryResolver) FetchFilterCatalogContent(ctx context.Context, selector *model.CatalogSelector) ([]*model.CatalogFilter, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.fetchCatalogFilter(ctx, provider, selector)
}

// GetClusterResources is the resolver for the getClusterResources field.
func (r *queryResolver) GetClusterResources(ctx context.Context, k8scontextIDs []string, namespace string) (*model.ClusterResources, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getClusterResources(ctx, provider, k8scontextIDs, namespace)
}

// GetMeshModelSummary is the resolver for the getMeshModelSummary field.
func (r *queryResolver) GetMeshModelSummary(ctx context.Context, selector model.MeshModelSummarySelector) (*model.MeshModelSummary, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getMeshModelSummary(ctx, provider, selector)
}

// FetchTelemetryComponents is the resolver for the fetchTelemetryComponents field.
func (r *queryResolver) FetchTelemetryComponents(ctx context.Context, contexts []string) ([]*model.TelemetryComp, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.getTelemetryComps(ctx, provider, contexts)
}

// ListenToAddonState is the resolver for the listenToAddonState field.
func (r *subscriptionResolver) ListenToAddonState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.AddonList, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToAddonState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// ListenToControlPlaneState is the resolver for the listenToControlPlaneState field.
func (r *subscriptionResolver) ListenToControlPlaneState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.ControlPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToControlPlaneState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// ListenToDataPlaneState is the resolver for the listenToDataPlaneState field.
func (r *subscriptionResolver) ListenToDataPlaneState(ctx context.Context, filter *model.ServiceMeshFilter) (<-chan []*model.DataPlane, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	if filter != nil {
		return r.listenToDataPlaneState(ctx, provider, filter)
	}

	return nil, ErrInvalidRequest
}

// ListenToOperatorState is the resolver for the listenToOperatorState field.
func (r *subscriptionResolver) ListenToOperatorState(ctx context.Context, k8scontextIDs []string) (<-chan *model.OperatorStatusPerK8sContext, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.listenToOperatorsState(ctx, provider, k8scontextIDs)
}

// ListenToMeshSyncEvents is the resolver for the listenToMeshSyncEvents field.
func (r *subscriptionResolver) ListenToMeshSyncEvents(ctx context.Context, k8scontextIDs []string) (<-chan *model.OperatorControllerStatusPerK8sContext, error) {
	// provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	// return r.listenToMeshSyncEvents(ctx, provider)
	return nil, nil
}

// SubscribePerfProfiles is the resolver for the subscribePerfProfiles field.
func (r *subscriptionResolver) SubscribePerfProfiles(ctx context.Context, selector model.PageFilter) (<-chan *model.PerfPageProfiles, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribePerfProfiles(ctx, provider, selector)
}

// SubscribePerfResults is the resolver for the subscribePerfResults field.
func (r *subscriptionResolver) SubscribePerfResults(ctx context.Context, selector model.PageFilter, profileID string) (<-chan *model.PerfPageResult, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribePerfResults(ctx, provider, selector, profileID)
}

// SubscribeBrokerConnection is the resolver for the subscribeBrokerConnection field.
func (r *subscriptionResolver) SubscribeBrokerConnection(ctx context.Context) (<-chan bool, error) {
	return r.subscribeBrokerConnection(ctx)
}

// SubscribeMesheryControllersStatus is the resolver for the subscribeMesheryControllersStatus field.
func (r *subscriptionResolver) SubscribeMesheryControllersStatus(ctx context.Context, k8scontextIDs []string) (<-chan []*model.MesheryControllersStatusListItem, error) {
	resChan := make(chan []*model.MesheryControllersStatusListItem)
	controllerHandlersPerContext, ok := ctx.Value(models.MesheryControllerHandlersKey).(map[string]map[models.MesheryController]controllers.IMesheryController)
	if !ok || len(controllerHandlersPerContext) == 0 || controllerHandlersPerContext == nil {
		er := model.ErrMesheryControllersStatusSubscription(fmt.Errorf("controller handlers are not configured for any of the contexts"))
		r.Log.Error(er)
		return nil, er
	}
	statusMapPerCtx := make(map[string]map[models.MesheryController]controllers.MesheryControllerStatus)
	// initialize the map
	for ctxID, ctrlHandlers := range controllerHandlersPerContext {
		for controller, handler := range ctrlHandlers {
			if _, ok := statusMapPerCtx[ctxID]; !ok {
				statusMapPerCtx[ctxID] = make(map[models.MesheryController]controllers.MesheryControllerStatus)
			}
			statusMapPerCtx[ctxID][controller] = handler.GetStatus()
		}
	}
	go func() {
		ctrlsStatusList := make([]*model.MesheryControllersStatusListItem, 0)
		// first send the initial status of the controllers
		for ctxID, statusMap := range statusMapPerCtx {
			for controller, status := range statusMap {
				ctrlsStatusList = append(ctrlsStatusList, &model.MesheryControllersStatusListItem{
					ContextID:  ctxID,
					Controller: model.GetInternalController(controller),
					Status:     model.GetInternalControllerStatus(status),
				})
			}
		}
		resChan <- ctrlsStatusList
		ctrlsStatusList = make([]*model.MesheryControllersStatusListItem, 0)
		// do this every 5 seconds
		for {
			for ctxID, ctrlHandlers := range controllerHandlersPerContext {
				for controller, handler := range ctrlHandlers {
					newStatus := handler.GetStatus()
					// if the status has changed, send that to the subscription
					if newStatus != statusMapPerCtx[ctxID][controller] {
						ctrlsStatusList = append(ctrlsStatusList, &model.MesheryControllersStatusListItem{
							ContextID:  ctxID,
							Controller: model.GetInternalController(controller),
							Status:     model.GetInternalControllerStatus(newStatus),
						})
						resChan <- ctrlsStatusList
					}
					// update the status list with newStatus
					statusMapPerCtx[ctxID][controller] = newStatus
					ctrlsStatusList = make([]*model.MesheryControllersStatusListItem, 0)
				}
			}
			// establish a watch connection to get updates, ideally in meshery-operator
			time.Sleep(time.Second * 5)
		}
	}()
	return resChan, nil
}

// SubscribeMeshSyncEvents is the resolver for the subscribeMeshSyncEvents field.
func (r *subscriptionResolver) SubscribeMeshSyncEvents(ctx context.Context, k8scontextIDs []string) (<-chan *model.MeshSyncEvent, error) {
	resChan := make(chan *model.MeshSyncEvent)
	// get handlers
	meshSyncDataHandlers, ok := ctx.Value(models.MeshSyncDataHandlersKey).(map[string]models.MeshsyncDataHandler)
	if !ok || len(meshSyncDataHandlers) == 0 || meshSyncDataHandlers == nil {
		er := model.ErrMeshSyncEventsSubscription(fmt.Errorf("meshsync data handlers are not configured for any of the contexts"))
		r.Log.Error(er)
		return nil, er
	}
	for ctxID, dataHandler := range meshSyncDataHandlers {
		brokerEventsChan := make(chan *broker.Message)
		err := dataHandler.ListenToMeshSyncEvents(brokerEventsChan)
		if err != nil {
			r.Log.Warn(err)
			r.Log.Info("skipping meshsync events subscription for contexId: %s", ctxID)
			continue
		}
		go func(ctxID string, brokerEventsChan chan *broker.Message) {
			for event := range brokerEventsChan {
				if event.EventType == broker.ErrorEvent {
					// TODO: Handle errors accordingly
					continue
				}
				// handle the events
				res := &model.MeshSyncEvent{
					ContextID: ctxID,
					Type:      string(event.EventType),
					Object:    event.Object,
				}
				resChan <- res
				go r.Config.DashboardK8sResourcesChan.PublishDashboardK8sResources()
			}
		}(ctxID, brokerEventsChan)
	}
	return resChan, nil
}

// SubscribeConfiguration is the resolver for the subscribeConfiguration field.
func (r *subscriptionResolver) SubscribeConfiguration(ctx context.Context, applicationSelector model.PageFilter, patternSelector model.PageFilter, filterSelector model.PageFilter) (<-chan *model.ConfigurationPage, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeConfiguration(ctx, provider, applicationSelector, patternSelector, filterSelector)
}

// SubscribeClusterResources is the resolver for the subscribeClusterResources field.
func (r *subscriptionResolver) SubscribeClusterResources(ctx context.Context, k8scontextIDs []string, namespace string) (<-chan *model.ClusterResources, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeClusterResources(ctx, provider, k8scontextIDs, namespace)
}

// SubscribeK8sContext is the resolver for the subscribeK8sContext field.
func (r *subscriptionResolver) SubscribeK8sContext(ctx context.Context, selector model.PageFilter) (<-chan *model.K8sContextsPage, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeK8sContexts(ctx, provider, selector)
}

// SubscribeMeshModelSummary is the resolver for the subscribeMeshModelSummary field.
func (r *subscriptionResolver) SubscribeMeshModelSummary(ctx context.Context, selector model.MeshModelSummarySelector) (<-chan *model.MeshModelSummary, error) {
	provider := ctx.Value(models.ProviderCtxKey).(models.Provider)
	return r.subscribeMeshModelSummary(ctx, provider, selector)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *subscriptionResolver) SubscribePerfResult(ctx context.Context, id string) (<-chan *model.MesheryResult, error) {
	panic(fmt.Errorf("not implemented"))
}
