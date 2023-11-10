package statetracker

import (
	"sync"
	"time"

	"github.com/lavanet/lava/utils"
	downtimev1 "github.com/lavanet/lava/x/downtime/v1"
)

type ConsumerEmergencyTrackerInf interface {
	GetLatestVirtualEpoch() uint64
}

type EmergencyTrackerMetrics interface {
	SetVirtualEpoch(virtualEpoch uint64)
}

type EmergencyTracker struct {
	lock             sync.RWMutex
	downtimeParams   downtimev1.Params
	virtualEpochsMap map[uint64]uint64
	latestEpoch      uint64
	latestEpochTime  time.Time
	metrics          EmergencyTrackerMetrics
}

func (cs *EmergencyTracker) GetVirtualEpoch(epoch uint64) uint64 {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.virtualEpochsMap[epoch]
}

func (cs *EmergencyTracker) GetLatestVirtualEpoch() uint64 {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.virtualEpochsMap[cs.latestEpoch]
}

func (cs *EmergencyTracker) SetDowntimeParams(params downtimev1.Params) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	cs.downtimeParams = params
}

func (cs *EmergencyTracker) GetDowntimeParams() downtimev1.Params {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.downtimeParams
}

func (cs *EmergencyTracker) UpdateEpoch(epoch uint64) {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	if epoch < cs.latestEpoch {
		return
	}

	cs.latestEpoch = epoch
	cs.latestEpochTime = time.Now()
}

func (cs *EmergencyTracker) blockNotFound(latestBlockTime time.Time) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	if time.Since(latestBlockTime) > cs.downtimeParams.DowntimeDuration {
		epochDuration := cs.downtimeParams.EpochDuration.Milliseconds()
		if epochDuration == 0 {
			utils.LavaFormatWarning("Emergency Tracker: downtime params are not initialized", nil)
			return
		}

		latestEpochTime := cs.latestEpochTime
		if cs.latestEpochTime.IsZero() {
			latestEpochTime = latestBlockTime
		}

		// division without rounding up to skip normal epoch,
		// if time since latestEpochTime > epochDuration => virtual epoch has started
		virtualEpoch := uint64(time.Since(latestEpochTime).Milliseconds() / epochDuration)
		if virtualEpoch < cs.virtualEpochsMap[cs.latestEpoch] {
			utils.LavaFormatError("Current virtual epoch is lower than stored", nil)
			return
		}
		// check if the new virtual epoch has started
		if virtualEpoch > 0 && cs.virtualEpochsMap[cs.latestEpoch] != virtualEpoch {
			utils.LavaFormatDebug("Emergency Tracker: emergency mode enabled",
				utils.Attribute{Key: "virtual_epoch", Value: virtualEpoch},
				utils.Attribute{Key: "epoch", Value: cs.latestEpoch},
			)
			if cs.metrics != nil {
				cs.metrics.SetVirtualEpoch(virtualEpoch)
			}
			cs.virtualEpochsMap[cs.latestEpoch] = virtualEpoch
		}
	}
}

func NewEmergencyTracker(metrics EmergencyTrackerMetrics) (emergencyTracker *EmergencyTracker, emergencyCallback func(latestBlockTime time.Time)) {
	emergencyTracker = &EmergencyTracker{
		virtualEpochsMap: map[uint64]uint64{},
		latestEpoch:      0,
		metrics:          metrics,
	}

	return emergencyTracker, emergencyTracker.blockNotFound
}