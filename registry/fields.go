package registry

import (
	"github.com/nkarakotova/lim-core/services"
)

type AppServiceFields struct {
	ClientService       services.ClientService
	CoachService        services.CoachService
	HallService         services.HallService
	TrainingService     services.TrainingService
}
