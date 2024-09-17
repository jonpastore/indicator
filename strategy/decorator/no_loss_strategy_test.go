// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/jonpastore/indicator

package decorator_test

import (
	"os"
	"testing"

	"github.com/jonpastore/indicator/v2/asset"
	"github.com/jonpastore/indicator/v2/helper"
	"github.com/jonpastore/indicator/v2/strategy"
	"github.com/jonpastore/indicator/v2/strategy/decorator"
	"github.com/jonpastore/indicator/v2/strategy/trend"
)

func TestNoLossStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/no_loss_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	innerStrategy := trend.NewAroonStrategy()
	strategy := decorator.NewNoLossStrategy(innerStrategy)

	actual := strategy.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoLossStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	innerStrategy := trend.NewAroonStrategy()
	strategy := decorator.NewNoLossStrategy(innerStrategy)

	report := strategy.Report(snapshots)

	fileName := "no_loss_strategy.html"
	defer os.Remove(fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
