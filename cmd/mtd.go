package main

import (
	"encoding/json"
	"fmt"
	"github.com/choonsiong/golib/format"
	"io/ioutil"
	"strconv"
)

type MTDData struct {
	Employer  *Header   `json:"employer"`
	Employees []*Detail `json:"employees"`
}

type Header struct {
	RecordType      string  `json:"record_type"`
	HQNumber        string  `json:"hq_number"`     // head quarter employer number
	BranchNumber    string  `json:"branch_number"` // branch employer number
	Year            string  `json:"year"`          // year of deduction
	Month           string  `json:"month"`         // month of deduction
	TotalMTD        float64 `json:"total_mtd"`     // total MTD amount
	TotalMTDRecord  int     `json:"total_mtd_record"`
	TotalCP38       float64 `json:"total_cp38"` // total CP38 amount
	TotalCP38Record int     `json:"total_cp38_record"`
}

type Detail struct {
	RecordType   string  `json:"record_type"`
	TaxReference string  `json:"tax_reference"` // tax reference number
	WifeCode     string  `json:"wife_code"`
	Name         string  `json:"name"` // employee name
	OldIC        string  `json:"old_ic"`
	NewIC        string  `json:"new_ic"`
	Passport     string  `json:"passport"`
	CountryCode  string  `json:"country_code"`
	MTDAmount    float64 `json:"mtd_amount"`
	CP38Amount   float64 `json:"cp38_amount"`
	Number       string  `json:"number"` // employee number
}

func parse(f string) (*MTDData, error) {
	mtd := new(MTDData)
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, mtd)
	if err != nil {
		return nil, err
	}
	return mtd, nil
}

func (d *MTDData) String() string {
	s := d.Employer.RecordType + "," + d.Employer.HQNumber + "," + d.Employer.BranchNumber + "," + d.Employer.Year + "," + d.Employer.Month + "," + fmt.Sprintf("%f", d.Employer.TotalMTD) + "," + strconv.Itoa(d.Employer.TotalMTDRecord) + "," + fmt.Sprintf("%f", d.Employer.TotalCP38) + "," + strconv.Itoa(d.Employer.TotalCP38Record) + "\n"

	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			s += e.RecordType + "," + e.TaxReference + "," + e.WifeCode + "," + e.Name + "," + e.OldIC + "," + e.NewIC + "," + e.Passport + "," + e.CountryCode + "," + fmt.Sprintf("%f", e.MTDAmount) + "," + fmt.Sprintf("%f", e.CP38Amount) + "," + e.Number + "\n"
		}
	}
	return s
}

func (d *MTDData) TotalMTDAmount() {
	var total float64
	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			total += e.MTDAmount
		}
	}
	d.Employer.TotalMTD = total
}

func (d *MTDData) TotalMTDRecord() {
	var count int
	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			if e.MTDAmount > 0 {
				count++
			}
		}
	}
	d.Employer.TotalMTDRecord = count
}

func (d *MTDData) TotalCP38Amount() {
	var total float64
	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			total += e.CP38Amount
		}
	}
	d.Employer.TotalCP38 = total
}

func (d *MTDData) TotalCP38Record() {
	var count int
	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			if e.CP38Amount > 0 {
				count++
			}
		}
	}
	d.Employer.TotalCP38Record = count
}

func (d *MTDData) Normalize() {
	d.Employer.HQNumber = format.LeftPaddingWithSize(10, d.Employer.HQNumber, "0")
	d.Employer.BranchNumber = format.LeftPaddingWithSize(10, d.Employer.BranchNumber, "0")
	d.Employer.Month = format.LeftPaddingWithSize(2, d.Employer.Month, "0")

	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			e.TaxReference = format.LeftPaddingWithSize(10, e.TaxReference, "0")
			e.Name = format.RightPaddingWithSize(60, e.Name, " ")
			e.OldIC = format.RightPaddingWithSize(12, e.OldIC, " ")
			e.NewIC = format.RightPaddingWithSize(12, e.NewIC, " ")
			e.Passport = format.RightPaddingWithSize(12, e.Passport, " ")
			e.Number = format.RightPaddingWithSize(10, e.Number, " ")
		}
	}
}

func (d *MTDData) Generate() (string, error) {
	var output string

	d.Normalize()
	d.TotalMTDAmount()
	d.TotalMTDRecord()
	d.TotalCP38Amount()
	d.TotalCP38Record()

	// Header row

	output += d.Employer.RecordType
	output += d.Employer.HQNumber
	output += d.Employer.BranchNumber
	output += d.Employer.Year
	output += d.Employer.Month

	totalMTDAmount := format.LeftPaddingWithSize(10, fmt.Sprintf("%.0f", d.Employer.TotalMTD*100), "0")
	output += totalMTDAmount

	totalMTDRecord := format.LeftPaddingWithSize(5, strconv.Itoa(d.Employer.TotalMTDRecord), "0")
	output += totalMTDRecord

	totalCP38Amount := format.LeftPaddingWithSize(10, fmt.Sprintf("%.0f", d.Employer.TotalCP38*100), "0")
	output += totalCP38Amount

	totalCP38Record := format.LeftPaddingWithSize(5, strconv.Itoa(d.Employer.TotalCP38Record), "0")
	output += totalCP38Record

	output += "\n"

	// Details row

	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			output += e.RecordType
			output += e.TaxReference
			output += e.WifeCode
			output += e.Name
			output += e.OldIC
			output += e.NewIC
			output += e.Passport
			output += format.LeftPaddingWithSize(2, e.CountryCode, " ")
			mtdAmount := format.LeftPaddingWithSize(8, fmt.Sprintf("%.0f", e.MTDAmount*100), "0")
			output += mtdAmount

			cp38Amount := format.LeftPaddingWithSize(8, fmt.Sprintf("%.0f", e.CP38Amount*100), "0")
			output += cp38Amount

			output += e.Number
			output += "\n"
		}
	}

	return output, nil
}

func (d *MTDData) out() error {
	s, err := d.Generate()
	if err != nil {
		return err
	}

	filename := "PCB_" + d.Employer.HQNumber + "_" + d.Employer.BranchNumber + "_" + d.Employer.Year + d.Employer.Month + ".txt"
	return ioutil.WriteFile(filename, []byte(s), 0644)
}
