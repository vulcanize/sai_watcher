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

package cup_actions_test

import (
	"math/big"

	"strings"

	"github.com/8thlight/sai_watcher/event_triggered/tub/cup_actions"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/libraries/shared"
)

var _ = Describe("Cup Data Converter", func() {
	It("converts cup entity to cup model", func() {
		cupAction := cup_actions.CupActionEntity{
			ID:              "0x0000000000000000000000000000000000000000000000000000000000000001",
			TransactionHash: "0x53c89dade0a03228ad7312d7f682018b58ad4410df2414410ff3b66993344c54",
			Act:             "lock",
			Arg:             "0x0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			Lad:             "0x000000000000000000000000cd5f8fa45e0ca0937f86006b9ee8fe1eedee5fc4",
			Ink:             big.NewFloat(123000000000000000000).String(),
			Art:             big.NewFloat(456000000000000000000).String(),
			Ire:             big.NewFloat(789000000000000000000).String(),
			Block:           4754490,
			Deleted:         false,
		}

		model := cup_actions.ConvertToModel(cupAction)

		Expect(model.ID).To(Equal(shared.HexToInt64(cupAction.ID)))
		Expect(model.TransactionHash).To(Equal(cupAction.TransactionHash))
		Expect(model.Act).To(Equal(cupAction.Act))
		Expect(model.Arg).To(Equal("1"))
		Expect(model.Lad).To(Equal(strings.ToLower(common.HexToAddress(cupAction.Lad).Hex())))
		Expect(model.Ink).To(Equal("123"))
		Expect(model.Art).To(Equal("456"))
		Expect(model.Ire).To(Equal("789"))
		Expect(model.Block).To(Equal(int64(4754490)))
		Expect(model.Deleted).To(Equal(cupAction.Deleted))
	})

	It("does not convert cup arg if act is 'give'", func() {
		cupAction := cup_actions.CupActionEntity{
			ID:              "0x0000000000000000000000000000000000000000000000000000000000000001",
			TransactionHash: "0x53c89dade0a03228ad7312d7f682018b58ad4410df2414410ff3b66993344c54",
			Act:             "give",
			Arg:             "0x12345",
			Lad:             "0x000000000000000000000000cd5f8fa45e0ca0937f86006b9ee8fe1eedee5fc4",
			Ink:             big.NewFloat(123000000000000000000).String(),
			Art:             big.NewFloat(456000000000000000000).String(),
			Ire:             big.NewFloat(789000000000000000000).String(),
			Block:           4754490,
			Deleted:         false,
		}

		model := cup_actions.ConvertToModel(cupAction)

		Expect(model.ID).To(Equal(shared.HexToInt64(cupAction.ID)))
		Expect(model.TransactionHash).To(Equal(cupAction.TransactionHash))
		Expect(model.Act).To(Equal(cupAction.Act))
		Expect(model.Arg).To(Equal(cupAction.Arg))
		Expect(model.Lad).To(Equal(strings.ToLower(common.HexToAddress(cupAction.Lad).Hex())))
		Expect(model.Ink).To(Equal("123"))
		Expect(model.Art).To(Equal("456"))
		Expect(model.Ire).To(Equal("789"))
		Expect(model.Block).To(Equal(int64(4754490)))
		Expect(model.Deleted).To(Equal(cupAction.Deleted))
	})

	It("converts topic2 (foo) from hex to an int ", func() {
		id := shared.HexToInt64("0x0000000000000000000000000000000000000000000000000000000000000001")
		Expect(id).To(Equal(int64(1)))
	})

	It("converts topic3 (bar) ", func() {
		bar := "0x00000000000000000000000000000000000000000000002b020ba44da84e6cae"
		convertedBar := cup_actions.Arg(bar, "lock")
		Expect(convertedBar).To(Equal("793.3573872357736"))
	})
})
