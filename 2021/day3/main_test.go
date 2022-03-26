package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_PowerConsumption(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *PowerConsumption
	}{
		{
			"Example test",
			args{strings.NewReader(`00100
				11110
				10110
				10111
				10101
				01111
				00111
				11100
				10000
				11001
				00010
				01010`)},
			&PowerConsumption{gammaRate: 22, epsilonRate: 9},
		},
		{
			"64-bit input test",
			args{strings.NewReader(`101010000100
			100001010100
			111100000101`)},
			&PowerConsumption{gammaRate: 2564, epsilonRate: 1531},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initDiagnosticRegisters(tt.args.r).PowerConsumption(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LifeSupport(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *LifeSupport
	}{
		{
			"Example test",
			args{strings.NewReader(`00100
				11110
				10110
				10111
				10101
				01111
				00111
				11100
				10000
				11001
				00010
				01010`)},
			&LifeSupport{oxygenGenerator: 23, co2Scrubber: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initDiagnosticRegisters(tt.args.r).LifeSupport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LifeSupport() = %v, want %v", got, tt.want)
			}
		})
	}
}
