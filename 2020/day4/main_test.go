package main

import (
	"io"
	"strings"
	"testing"
)

func Test_day4(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name             string
		args             args
		wantValidCount   int
		wantInvalidCount int
	}{
		{"example1", args{reader: strings.NewReader(
			`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`)}, 1, 0},
		{"example2", args{reader: strings.NewReader(
			`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`)}, 0, 1},
		{"example3", args{reader: strings.NewReader(
			`hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`)}, 1, 0},
		{"example4", args{reader: strings.NewReader(
			`hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)}, 0, 1},
		{"combined examples", args{reader: strings.NewReader(
			`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValidCount, gotInvalidCount := day4(tt.args.reader)
			if gotValidCount != tt.wantValidCount {
				t.Errorf("day4() gotValidCount = %v, want %v", gotValidCount, tt.wantValidCount)
			}
			if gotInvalidCount != tt.wantInvalidCount {
				t.Errorf("day4() gotInvalidCount = %v, want %v", gotInvalidCount, tt.wantInvalidCount)
			}
		})
	}
}

func Test_passport_isValid(t *testing.T) {
	type fields struct {
		byr int
		iyr int
		eyr int
		hgt string
		hcl string
		ecl string
		pid string
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"valid", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "193cm", pid: "123456789", iyr: 2018, byr: 1926}, true},
		{"valid2", fields{pid: "087499704", hgt: "74in", ecl: "grn", iyr: 2012, eyr: 2030, byr: 1980, hcl: "#623a2f"}, true},
		{"valid3", fields{eyr: 2029, ecl: "blu", cid: "129", byr: 1989, iyr: 2014, pid: "896056539", hcl: "#a97842", hgt: "165cm"}, true},
		{"valid4", fields{eyr: 2029, ecl: "blu", cid: "129", byr: 1989, iyr: 2014, pid: "896056539", hcl: "#a97842", hgt: "165cm"}, true},
		{"valid5", fields{hcl: "#888785", hgt: "164cm", byr: 2001, iyr: 2015, cid: "88", pid: "545766238", ecl: "hzl", eyr: 2022}, true},
		{"valid6", fields{iyr: 2010, hgt: "158cm", hcl: "#b6652a", ecl: "blu", byr: 1944, eyr: 2021, pid: "093154719"}, true},
		{"178cm tall", fields{iyr: 2010, hgt: "178cm", hcl: "#b6652a", ecl: "blu", byr: 1944, eyr: 2021, pid: "093154719"}, true},
		{"189cm tall", fields{iyr: 2010, hgt: "178cm", hcl: "#b6652a", ecl: "blu", byr: 1944, eyr: 2021, pid: "093154719"}, true},
		{"cid ignored", fields{hcl: "#123abc", iyr: 2012, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1992}, true},

		{"too tall cm", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "194cm", pid: "123456789", iyr: 2018, byr: 1926}, false},
		{"missing unit", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "160", pid: "123456789", iyr: 2018, byr: 1926}, false},
		{"too short cm", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "149cm", pid: "123456789", iyr: 2018, byr: 1926}, false},
		{"too tall in", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "77in", pid: "123456789", iyr: 2018, byr: 1926}, false},
		{"too short in", fields{eyr: 2024, cid: "100", hcl: "#18171d", ecl: "amb", hgt: "58in", pid: "123456789", iyr: 2018, byr: 1926}, false},
		{"missing # in hcl", fields{hcl: "dab227", iyr: 2012, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1992, cid: "277"}, false},
		{"non-hex in hcl", fields{hcl: "#123abz", iyr: 2012, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1992, cid: "277"}, false},

		{"birth year too early", fields{hcl: "#dab227", iyr: 2012, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1919, cid: "277"}, false},
		{"birth year too late", fields{hcl: "#dab227", iyr: 2012, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 2003, cid: "277"}, false},
		{"issue year too early", fields{hcl: "#dab227", iyr: 2009, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1967, cid: "277"}, false},
		{"birth year too late", fields{hcl: "#dab227", iyr: 2021, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1967, cid: "277"}, false},
		{"issue year too early", fields{hcl: "#dab227", iyr: 2009, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1967, cid: "277"}, false},
		{"birth year too late", fields{hcl: "#dab227", iyr: 2021, ecl: "brn", hgt: "182cm", pid: "021572410", eyr: 2020, byr: 1967, cid: "277"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &passport{
				byr: tt.fields.byr,
				iyr: tt.fields.iyr,
				eyr: tt.fields.eyr,
				hgt: tt.fields.hgt,
				hcl: tt.fields.hcl,
				ecl: tt.fields.ecl,
				pid: tt.fields.pid,
				cid: tt.fields.cid,
			}
			if got := p.isValid(); got != tt.want {
				t.Errorf("passport.checkPassport() = %v, want %v", got, tt.want)
			}
		})
	}
}
