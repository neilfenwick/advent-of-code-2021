package main

type DiagnosticRegisters struct {
	gamma         []int
	epsilon       []int
	oxygen        []int
	co2           []int
	readingCounts []int
	width         int
}

func NewDiagnosticRegisters(width int) *DiagnosticRegisters {
	d := DiagnosticRegisters{}
	d.gamma = make([]int, width)
	d.epsilon = make([]int, width)
	d.oxygen = make([]int, 0)
	d.co2 = make([]int, 0)
	d.readingCounts = make([]int, width)
	d.width = width
	return &d
}

func (r *DiagnosticRegisters) AddReading(reading int) {
	for i := 0; i < r.width; i++ {
		bitMask := (1 << (r.width - i - 1)) // left shift to test bit at current position
		if reading&bitMask > 0 {
			r.gamma[i] = r.gamma[i] + 1
		} else {
			r.epsilon[i] = r.epsilon[i] + 1
		}
		r.readingCounts[i] = r.readingCounts[i] + 1
	}
	r.oxygen = append(r.oxygen, reading)
	r.co2 = append(r.co2, reading)
}

func (r *DiagnosticRegisters) PowerConsumption() *PowerConsumption {
	var (
		gamma, epsilon int64
	)

	for i, count := range r.readingCounts {
		if float32(r.gamma[i])/float32(count) >= .5 {
			gamma = gamma | int64(1<<(r.width-i-1))
		} else {
			epsilon = epsilon | int64(1<<(r.width-i-1))
		}
	}

	return &PowerConsumption{gammaRate: int(gamma), epsilonRate: int(epsilon)}
}

func (r *DiagnosticRegisters) LifeSupport() *LifeSupport {
	if len(r.oxygen) > 1 {
		oxygen := r.oxygenReading(r.oxygen, 0)
		r.oxygen = []int{oxygen}
	}

	if len(r.co2) > 1 {
		co2 := r.co2Reading(r.co2, 0)
		r.co2 = []int{co2}
	}

	return &LifeSupport{oxygenGenerator: r.oxygen[0], co2Scrubber: r.co2[0]}
}

func (r *DiagnosticRegisters) oxygenReading(readings []int, bitIndexToCheck int) int {
	oxygenSubRegisters := NewDiagnosticRegisters(r.width)
	for _, r := range readings {
		oxygenSubRegisters.AddReading(r)
	}

	mostFrequentBits := oxygenSubRegisters.PowerConsumption().gammaRate
	for i := len(oxygenSubRegisters.oxygen) - 1; i >= 0; i-- {
		bitsNotMatching := oxygenSubRegisters.oxygen[i] ^ mostFrequentBits
		if int64(bitsNotMatching)&int64(1<<(oxygenSubRegisters.width-bitIndexToCheck-1)) > 0 {
			oxygenSubRegisters.oxygen = append(oxygenSubRegisters.oxygen[:i], oxygenSubRegisters.oxygen[i+1:]...)
		}
	}

	if len(oxygenSubRegisters.oxygen) == 1 {
		return oxygenSubRegisters.oxygen[0]
	}

	return oxygenSubRegisters.oxygenReading(oxygenSubRegisters.oxygen, bitIndexToCheck+1)
}

func (r *DiagnosticRegisters) co2Reading(readings []int, bitIndexToCheck int) int {
	co2SubRegisters := NewDiagnosticRegisters(r.width)
	for _, r := range readings {
		co2SubRegisters.AddReading(r)
	}

	mostFrequentBits := co2SubRegisters.PowerConsumption().gammaRate
	for i := len(co2SubRegisters.co2) - 1; i >= 0; i-- {
		bitsNotMatching := co2SubRegisters.co2[i] ^ mostFrequentBits
		if int64(bitsNotMatching)&int64(1<<(co2SubRegisters.width-bitIndexToCheck-1)) == 0 {
			co2SubRegisters.co2 = append(co2SubRegisters.co2[:i], co2SubRegisters.co2[i+1:]...)
		}
	}

	if len(co2SubRegisters.co2) == 1 {
		return co2SubRegisters.co2[0]
	}

	return co2SubRegisters.co2Reading(co2SubRegisters.co2, bitIndexToCheck+1)
}

type PowerConsumption struct {
	gammaRate   int
	epsilonRate int
}

type LifeSupport struct {
	oxygenGenerator int
	co2Scrubber     int
}
