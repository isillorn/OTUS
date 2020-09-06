package main

import (
	"crypto/aes"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

/*
Параллельное исполнение
Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).

Функция принимает на вход:
- слайс с заданиями `[]func() error`;
- число заданий которые можно выполнять параллельно (`N`);
- максимальное число ошибок после которого нужно приостановить обработку.

Учесть что задания могут выполняться разное время
*/

// https://github.com/ivandavidov/minimal/releases/download/15-Dec-2019/minimal_linux_live_15-Dec-2019_64-bit_mixed.iso

const allJob = 50 // всего заданий
const simJob = 4  // одновременно выполняющихся заданий
const maxErr = 5  // допустимое количество ошибок

var chanErr chan error

func heavyLoader() {

	var rndBase, rndBase2, rndBase3, rndBase4 float64
	var rndBaseUint64, rndBase2Uint64, rndBase3Uint64, rndBase4Uint64 uint64
	var encoded, decoded, bytes []byte

	rndBase = rand.Float64()
	rndCount := rand.Intn(1000000)

	//fmt.Printf("%v %v\n", rndBase, rndCount)
	fmt.Printf("[%v] Thread started\n", rndCount)

	timer := time.NewTimer(3 * time.Second)

	for i := 0; i < rndCount; i++ {
		select {

		case <-timer.C:
			chanErr <- fmt.Errorf("[%v] FAILED (timeout)", rndCount)
			return
		default:
			rndBase = math.Gamma((math.Sqrt(math.Exp(rand.Float64()) * math.Acos(rand.Float64()) / rndBase)))
			rndBase2 = math.Sin(rndBase) * rand.Float64()
			rndBase3 = math.Log(math.Tan(rndBase2) * rand.Float64())
			rndBase4 = math.Sqrt(rndBase3) * rand.Float64()
			rndBaseUint64 = math.Float64bits(rndBase)
			rndBase2Uint64 = math.Float64bits(rndBase2)
			rndBase3Uint64 = math.Float64bits(rndBase3)
			rndBase4Uint64 = math.Float64bits(rndBase4)

			bytes = make([]byte, 32)
			encoded = make([]byte, 32)
			decoded = make([]byte, 32)

			for j := 0; j < 7; j++ {
				bytes[4*j] = byte(rndBaseUint64 >> j * 8)
				bytes[4*j+1] = byte(rndBase2Uint64 >> j * 8)
				bytes[4*j+2] = byte(rndBase3Uint64 >> j * 8)
				bytes[4*j+3] = byte(rndBase4Uint64 >> j * 8)
			}

			block, err := aes.NewCipher(bytes)
			if err == nil {
				block.Encrypt(encoded, []byte("This is a cool encrypted text!"))
				block.Decrypt(decoded, encoded)
				//fmt.Printf("[%v] %s\n", rndBase, string(decoded))
			}

		}
	}
	chanErr <- fmt.Errorf("[%v] OK", rndCount)
}

func multiThreader(f []func(), simJob int, maxErr int) {

	var first, last, okCnt, errCnt, totErr int
	var err error

	for first = 0; (first < len(f)) && (totErr < maxErr); first = first + simJob {
		okCnt, errCnt = 0, 0

		if first+simJob < len(f) {
			last = first + simJob
		} else {
			last = len(f)
		}

		fChunc := f[first:last]

		for i := range fChunc {
			go f[i]()
		}

		for (okCnt+errCnt < len(fChunc)) && (totErr < maxErr) {
			select {

			case err = <-chanErr:
				fmt.Printf("%v\n", err.Error())

				if strings.HasSuffix(err.Error(), "OK") {
					okCnt++
				} else {
					errCnt++
					totErr++
				}
			}
		}
	}
	fmt.Printf("\nFinished processing %d/%d tasks  (in %d threads) with %d errors", first-simJob+okCnt+errCnt, len(f), simJob, totErr)

}

func main() {
	var funcSlice []func()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < allJob; i++ {
		funcSlice = append(funcSlice, func() { heavyLoader() })
	}

	chanErr = make(chan error, simJob)

	multiThreader(funcSlice, simJob, maxErr)

}
