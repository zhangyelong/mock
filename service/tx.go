package service

import (
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/key"
	"github.com/kaifei-bianjie/mock/sign"
	"github.com/kaifei-bianjie/mock/types"
	"log"
)

func BatchGenSignedTxData(num int, subFaucets []conf.SubFaucet) []string {
	var (
		method       = "BatchGenSignedTx"
		signedTxData []string
	)
	resChan := make(chan types.GenSignedTxDataRes)

	senderInfos, err := key.NewAccount(num, subFaucets)
	if err != nil {
		// TODO: handle err
	}

	lens := len(senderInfos)
	if lens > 0 {
		log.Printf("%v: now use %v goroutine to gen signed data\n",
			method, lens)

		for i, senderInfo := range senderInfos {
			go sign.GenSignedTxData(senderInfo, conf.DefaultReceiverAddr, resChan, i)
		}

		counter := 0
		for {
			res := <-resChan
			counter++
			if res.Res != "" {
				log.Printf("%v: successed, goroutine %v gen signed tx data. now left %v goroutine\n",
					method, res.ChanNum, lens-counter)
				signedTxData = append(signedTxData, res.Res)
			} else {
				log.Printf("%v: failed, goroutine %v gen signed tx data. now left %v goroutine\n",
					method, res.ChanNum, lens-counter)
			}

			if counter == lens {
				log.Printf("%v: all sign tx goroutine over\n", method)
				break
			}
		}
	} else {
		log.Printf("%v: no signed tx data\n", method)
	}

	return signedTxData
}
