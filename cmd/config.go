package main

import (
	"encoding/json"
	"io/ioutil"
)

type MTDData struct {
	Employer  Header   `json:"employer"`
	Employees []Detail `json:"employees"`
}

type Header struct {
	RecordType      string `json:"record_type"`
	HQNumber        string `json:"hq_number"`     // head quarter employer number
	BranchNumber    string `json:"branch_number"` // branch employer number
	Year            string `json:"year"`          // year of deduction
	Month           string `json:"month"`         // month of deduction
	TotalMTD        string `json:"total_mtd"`     // total MTD amount
	TotalMTDRecord  string `json:"total_mtd_record"`
	TotalCP38       string `json:"total_cp38"` // total CP38 amount
	TotalCP38Record string `json:"total_cp38_record"`
}

type Detail struct {
	RecordType   string `json:"record_type"`
	TaxReference string `json:"tax_reference"` // tax reference number
	WifeCode     string `json:"wife_code"`
	Name         string `json:"name"` // employee name
	OldIC        string `json:"old_ic"`
	NewIC        string `json:"new_ic"`
	Passport     string `json:"passport"`
	CountryCode  string `json:"country_code"`
	MTDAmount    string `json:"mtd_amount"`
	CP38Amount   string `json:"cp38_amount"`
	Number       string `json:"number"` // employee number
}

func parse(f string) (MTDData, error) {
	var mtd MTDData
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return mtd, err
	}
	err = json.Unmarshal(bs, &mtd)
	if err != nil {
		return mtd, err
	}
	return mtd, nil
}

func (d MTDData) String() string {
	s := d.Employer.RecordType + "," + d.Employer.HQNumber + "," + d.Employer.BranchNumber + "," + d.Employer.Year + "," + d.Employer.Month + "," + d.Employer.TotalMTD + "," + d.Employer.TotalMTDRecord + "," + d.Employer.TotalCP38 + "," + d.Employer.TotalCP38Record + "\n"

	if len(d.Employees) != 0 {
		for _, e := range d.Employees {
			s += e.RecordType + "," + e.TaxReference + "," + e.WifeCode + "," + e.Name + "," + e.OldIC + "," + e.NewIC + "," + e.Passport + "," + e.CountryCode + "," + e.MTDAmount + "," + e.CP38Amount + "," + e.Number + "\n"
		}
	}
	return s
}
