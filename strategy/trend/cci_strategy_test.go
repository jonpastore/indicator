// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/jonpastore/indicator

package trend_test

import (
	"os"
	"testing"

	"github.com/jonpastore/indicator/v2/asset"
	"github.com/jonpastore/indicator/v2/helper"
	"github.com/jonpastore/indicator/v2/strategy"
	"github.com/jonpastore/indicator/v2/strategy/trend"
)

func TestCciStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/cci_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	cci := trend.NewCciStrategy()
	actual := cci.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCciStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	cci := trend.NewCciStrategy()

	report := cci.Report(snapshots)

	fileName := "cci_strategy.html"
	defer os.Remove(fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
