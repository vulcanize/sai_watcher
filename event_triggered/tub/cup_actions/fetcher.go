// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cup_actions

import (
	"github.com/8thlight/sai_watcher/event_triggered/tub"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type CupFetcher struct {
	Blockchain core.Blockchain
}

var CupsContractMethod = "cups"

func (cupDataFetcher CupFetcher) FetchCupData(methodArg interface{}, blockNumber int64) (*Cup, error) {
	abiJSON := tub.TubContractABI
	address := tub.TubContractAddress
	method := CupsContractMethod
	result := &Cup{}
	err := cupDataFetcher.Blockchain.FetchContractData(abiJSON, address, method, methodArg, result, blockNumber)
	return result, err
}
