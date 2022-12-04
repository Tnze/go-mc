package maths

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"time"
)

type TpsCalculator struct {
	Started bool
	// TickRate is the number of ticks per second.
	TickRate   float64
	lastNTicks []float64
	// TimeLastUpdate is the time of the last update.
	TimeLastUpdate time.Time
	// This is the callback for the tick event. Channels are too slow to use unfortunately
	Callback func() basic.Error
}

func (t *TpsCalculator) Start() {
	t.Started = true
	t.TimeLastUpdate = time.Now()

	go func() {
		for {
			time.Sleep(time.Duration(50*t.TickAverage()) * time.Millisecond) // Synchronise with the server's TPS
			if t.Callback != nil {
				if err := t.Callback(); !err.Is(basic.NoError) {
					fmt.Println("Error in TPS callback:", err)
				}
			}
		}
	}() // Create a new thread for the tick event to avoid the delay from the main thread
}

func (t *TpsCalculator) Update() {
	if !t.Started {
		t.Start()
	}
	if t.lastNTicks == nil {
		t.lastNTicks = make([]float64, 5)
		for i := range t.lastNTicks {
			t.lastNTicks[i] = 0.99 // Meh, workaround for the first ticks.
		}
	}
	t.lastNTicks = append(t.lastNTicks, Clamp(time.Since(t.TimeLastUpdate).Seconds(), 0, 1))
	t.lastNTicks = t.lastNTicks[1:]
	t.TickRate = 1 / t.lastNTicks[0]
	t.TimeLastUpdate = time.Now()
}

/*
TickAverage

	@return float64 - The average tick rate from 0 to 1.
*/
func (t *TpsCalculator) TickAverage() float64 {
	var sum float64
	for _, v := range t.lastNTicks {
		sum += v
	}
	return sum / float64(len(nonZero(t.lastNTicks)))
}

func nonZero(arr []float64) []float64 {
	var out []float64
	for _, v := range arr {
		if v != 0 {
			out = append(out, v)
		}
	}
	return out
}

func Clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
