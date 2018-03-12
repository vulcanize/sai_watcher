package pep

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/config"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type MockPepsRepository struct {
	values       []string
	blockNumbers []int64
	logIds       []int64
}

func (mpr *MockPepsRepository) CheckNewPep() ([]*core.WatchedEvent, error) {
	panic("implement me")
}

func (mpr *MockPepsRepository) CreatePep(value string, blockNumber int64, logId int64) error {
	mpr.values = append(mpr.values, value)
	mpr.blockNumbers = append(mpr.blockNumbers, blockNumber)
	mpr.logIds = append(mpr.logIds, logId)
	return nil
}

type MockWatchedEventsRepository struct {
}

var logId = int64(67890)
var blockNumber = int64(12345)
var eventDataValue = "899009000000000000000"

func (MockWatchedEventsRepository) GetWatchedEvents(name string) ([]*core.WatchedEvent, error) {
	data := big.NewInt(0)
	data.SetString(eventDataValue, 10)
	dataEncoded := common.BigToHash(data).Hex()
	watchedEvent := core.WatchedEvent{
		LogID:       logId,
		Name:        "",
		BlockNumber: blockNumber,
		Address:     "",
		TxHash:      "",
		Index:       0,
		Topic0:      "",
		Topic1:      "",
		Topic2:      "",
		Topic3:      "",
		Data:        dataEncoded,
	}
	return []*core.WatchedEvent{&watchedEvent}, nil
}

type FakePepFetcher struct{}

func (FakePepFetcher) GetBlockByNumber(blockNumber int64) core.Block {
	panic("implement me")
}

func (FakePepFetcher) GetLogs(contract core.Contract, startingBlockNumber *big.Int, endingBlockNumber *big.Int) ([]core.Log, error) {
	panic("implement me")
}

func (FakePepFetcher) Node() core.Node {
	panic("implement me")
}

func (FakePepFetcher) LastBlock() *big.Int {
	panic("implement me")
}

func (FakePepFetcher) GetContractOutput(address string, input []byte, blockNumber int64) ([]byte, error) {
	panic("implement me")
}

func (FakePepFetcher) FetchContractData(abiJSON string, address string, method string, methodArg interface{}, result interface{}, blockNumber int64) error {
	panic("implement me")
}

var _ = Describe("pep updater", func() {
	It("Updates a pep", func() {
		db, err := postgres.NewDB(config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_private",
			Port:     5432,
		}, core.Node{})
		Expect(err).NotTo(HaveOccurred())
		db.Query(`DELETE FROM cups`)
		db.Query(`DELETE FROM logs`)
		db.Query(`DELETE FROM log_filters`)

		pepUpdater := NewPepHandler(db, &FakePepFetcher{})
		pepsRepository := &MockPepsRepository{}
		watchedEventsRepository := MockWatchedEventsRepository{}
		pepUpdater = &Handler{IPepsRepository: pepsRepository, WatchedEventRepository: watchedEventsRepository}

		pepUpdater.Execute()

		Expect(len(pepsRepository.values)).To(Equal(1))
		Expect(pepsRepository.values[0]).To(Equal(eventDataValue))
		Expect(len(pepsRepository.blockNumbers)).To(Equal(1))
		Expect(pepsRepository.blockNumbers[0]).To(Equal(blockNumber))
		Expect(len(pepsRepository.logIds)).To(Equal(1))
		Expect(pepsRepository.logIds[0]).To(Equal(logId))
	})
})
