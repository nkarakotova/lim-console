package registry

import (
	"github.com/nkarakotova/lim-core/services"
)

type AppServiceFields struct {
	ClientService       services.ClientService
	CoachService        services.CoachService
	DirectionService    services.DirectionService
	HallService         services.HallService
	SubscriptionService services.SubscriptionService
	TrainingService     services.TrainingService
}
