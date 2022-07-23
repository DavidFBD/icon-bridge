/*
 * Copyright 2021 ICON Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bsc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain"
	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/bsc/binding"

	"math/big"

	"github.com/icon-project/icon-bridge/common/errors"
	"github.com/icon-project/icon-bridge/common/log"
)

type receiver struct {
	cl  *client
	src chain.BTPAddress
	dst chain.BTPAddress
	log log.Logger
	opt struct {
	}
}

func NewReceiver(src, dst chain.BTPAddress, endpoints []string, opt map[string]interface{}, l log.Logger) (chain.Receiver, error) {
	r := &receiver{
		src: src,
		dst: dst,
		log: l,
	}
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("empty urls: %v", endpoints)
	}
	b, err := json.Marshal(opt)
	if err != nil {
		l.Panicf("fail to marshal opt:%#v err:%+v", opt, err)
	}
	if err = json.Unmarshal(b, &r.opt); err != nil {
		l.Panicf("fail to unmarshal opt:%#v err:%+v", opt, err)
	}
	r.cl, err = NewClient(endpoints, src.ContractAddress(), l)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *receiver) newBTPMessage(v *BlockNotification) ([]*chain.Receipt, error) {
	var receipts []*chain.Receipt
	var events []*chain.Event
	srcContractAddress := HexToAddress(r.src.ContractAddress())

	var index, BlockNumber uint64
	events = events[:0]
	for _, vLog := range v.Logs {
		if bmcMsg, err := binding.UnpackEventLog(vLog.Data); err == nil {
			events = append(events, &chain.Event{
				Message:  bmcMsg.Msg,
				Next:     chain.BTPAddress(bmcMsg.Next),
				Sequence: bmcMsg.Seq.Uint64(),
			})
			index = uint64(vLog.Index)
			BlockNumber = vLog.BlockNumber
		}
	}
	if len(events) > 0 {
		rp := &chain.Receipt{}
		rp.Index = uint64(index)
		rp.Height = BlockNumber
		rp.Events = append(rp.Events, events...)
		receipts = append(receipts, rp)
		r.log.Debugf("event found for height %v & address: %v", rp.Height, srcContractAddress)
	}

	return receipts, nil
}

func (r *receiver) receiveLoop(ctx context.Context, height int64, callback func(v *BlockNotification) error) error {
	r.log.Debugf("ReceiveLoop connected")
	br := &BlockRequest{
		Height:             big.NewInt(height),
		SrcContractAddress: HexToAddress(r.src.ContractAddress()),
	}
	r.cl.MonitorBlock(ctx, br,
		func(v *BlockNotification) error {
			if err := callback(v); err != nil {
				return errors.Wrapf(err, "receiveLoop: callback: %v", err)
			}
			return nil
		},
	)
	return nil
}

func (r *receiver) StopReceiveLoop() {
	r.cl.CloseAllMonitor()
}

func (r *receiver) Subscribe(
	ctx context.Context, msgCh chan<- *chain.Message,
	opts chain.SubscribeOptions) (errCh <-chan error, err error) {

	opts.Seq++

	_errCh := make(chan error)

	go func() {
		defer close(_errCh)
		lastHeight := opts.Height - 1
		if err := r.receiveLoop(ctx, int64(opts.Height),
			func(v *BlockNotification) error {
				r.log.WithFields(log.Fields{"height": v.Height}).Debug("block notification")

				if v.Height.Uint64() != lastHeight+1 {
					r.log.Errorf("expected v.Height == %d, got %d", lastHeight+1, v.Height.Uint64())
					return fmt.Errorf(
						"block notification: expected=%d, got=%d",
						lastHeight+1, v.Height.Uint64())
				}

				receipts, err := r.newBTPMessage(v)
				if err != nil {
					return fmt.Errorf("Error creating BTP message from block notification: %v", err)
				}

				for _, receipt := range receipts {
					events := receipt.Events[:0]
					for _, event := range receipt.Events {
						r.log.Infof("evt seq %v seq no %v", event.Sequence, opts.Seq)
						switch {
						case event.Sequence == opts.Seq:
							events = append(events, event)
							opts.Seq++
						case event.Sequence > opts.Seq:
							r.log.WithFields(log.Fields{
								"seq": log.Fields{"got": event.Sequence, "expected": opts.Seq},
							}).Error("invalid event seq")
							return fmt.Errorf("invalid event seq")
						}
					}
					receipt.Events = events
				}

				msgCh <- &chain.Message{Receipts: receipts}
				lastHeight++
				return nil
			}); err != nil {
			// TODO decide whether to ignore or handle err
			r.log.Errorf("receiveLoop terminated: %v", err)
			_errCh <- err
		}
	}()

	return _errCh, nil
}
