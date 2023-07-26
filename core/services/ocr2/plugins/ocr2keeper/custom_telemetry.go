package ocr2keeper

import (
	"context"
	"time"

	httypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker/types"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/synchronization/telem"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/libocr/commontypes"
	"google.golang.org/protobuf/proto"
)

type AutomationCustomTelemetryService struct {
	utils.StartStopOnce
	monitoringEndpoint  commontypes.MonitoringEndpoint
	headBroadcaster     httypes.HeadBroadcaster
	headCh              chan blockKey
	customTelemChanSize uint8
	unsubscribe         func()
	chDone              chan struct{}
	lggr                logger.Logger
	version             string
}

// NewAutomationCustomTelemetryService creates a telemetry service for new blocks and node version
func NewAutomationCustomTelemetryService(me commontypes.MonitoringEndpoint, hb httypes.HeadBroadcaster, lggr logger.Logger, vers string) *AutomationCustomTelemetryService {
	return &AutomationCustomTelemetryService{
		monitoringEndpoint:  me,
		headBroadcaster:     hb,
		headCh:              make(chan blockKey, 50),
		customTelemChanSize: 50,
		chDone:              make(chan struct{}),
		lggr:                lggr,
		version:             vers,
	}
}

// Start starts
func (e *AutomationCustomTelemetryService) Start(context.Context) error {
	return e.StartOnce("AutomationCustomTelemetryService", func() error {
		versionMsg := &telem.NodeVersion{
			Timestamp:   uint64(time.Now().UTC().UnixMilli()),
			NodeVersion: e.version,
		}
		wrappedMessage := &telem.AutomationTelemWrapper{
			Msg: &telem.AutomationTelemWrapper_NodeVersion{
				NodeVersion: versionMsg,
			},
		}
		bytes, err := proto.Marshal(wrappedMessage)
		if err != nil {
			e.lggr.Errorf("Error occured while marshalling the message: %v", err)
		}
		e.monitoringEndpoint.SendLog(bytes)
		e.lggr.Infof("BlockNumber Message Sent to Endpoint: %s", wrappedMessage.String())
		_, e.unsubscribe = e.headBroadcaster.Subscribe(&headWrapper{e.headCh})
		go func() {
			e.lggr.Infof("Started enhanced telemetry service")
			for {
				select {
				case blockKey := <-e.headCh:
					blockNumMsg := &telem.BlockNumber{
						Timestamp:   uint64(time.Now().UTC().UnixMilli()),
						BlockNumber: uint64(blockKey.block),
						BlockHash:   blockKey.hash,
					}
					wrappedMessage := &telem.AutomationTelemWrapper{
						Msg: &telem.AutomationTelemWrapper_BlockNumber{
							BlockNumber: blockNumMsg,
						},
					}
					bytes, err := proto.Marshal(wrappedMessage)
					if err != nil {
						e.lggr.Errorf("Error occured while marshalling the message: %v", err)
					}
					e.monitoringEndpoint.SendLog(bytes)
					e.lggr.Infof("BlockNumber Message Sent to Endpoint: %s", wrappedMessage.String())
				case <-e.chDone:
					return
				}
			}
		}()
		return nil
	})
}

// Close stops go routines and closes channels
func (e *AutomationCustomTelemetryService) Close() error {
	return e.StopOnce("AutomationCustomTelemetryService", func() error {
		e.chDone <- struct{}{}
		close(e.headCh)
		close(e.chDone)
		e.unsubscribe()
		e.lggr.Infof("Stopping custom telemetry service for job")
		return nil
	})
}

// blockKey contains block and hash info for BlockNumber telemetry message
type blockKey struct {
	block int64
	hash  string
}

// headWrapper is passed into HeadBroadcaster's subscribe() function, must implement OnNewLongestChain(_ context.Context, head *evmtypes.Head)
type headWrapper struct {
	headCh chan blockKey
}

// OnNewLongestChain sends block number and hash to head channel where message will be sent to monitoring endpoint
func (hw *headWrapper) OnNewLongestChain(_ context.Context, head *evmtypes.Head) {
	if head != nil {
		hw.headCh <- blockKey{
			block: head.Number,
			hash:  head.BlockHash().Hex(),
		}
	}
}
