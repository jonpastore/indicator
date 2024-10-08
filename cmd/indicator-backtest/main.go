// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/jonpastore/indicator

// main is the indicator backtest command line program.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jonpastore/indicator/v2/asset"
	"github.com/jonpastore/indicator/v2/backtest"
	"github.com/jonpastore/indicator/v2/strategy"
	"github.com/jonpastore/indicator/v2/strategy/compound"
	"github.com/jonpastore/indicator/v2/strategy/momentum"
	"github.com/jonpastore/indicator/v2/strategy/trend"
	"github.com/jonpastore/indicator/v2/strategy/volatility"
)

func main() {
	var repositoryName string
	var repositoryConfig string
	var reportName string
	var reportConfig string
	var workers int
	var lastDays int
	var addSplits bool
	var addAnds bool

	fmt.Fprintln(os.Stderr, "Indicator Backtest")
	fmt.Fprintln(os.Stderr, "Copyright (c) 2021-2024 Onur Cinar.")
	fmt.Fprintln(os.Stderr, "The source code is provided under GNU AGPLv3 License.")
	fmt.Fprintln(os.Stderr, "https://github.com/jonpastore/indicator")
	fmt.Fprintln(os.Stderr)

	flag.StringVar(&repositoryName, "repository-name", "filesystem", "repository name")
	flag.StringVar(&repositoryConfig, "repository-config", "", "repository config")
	flag.StringVar(&reportName, "report-name", "html", "report name")
	flag.StringVar(&reportConfig, "report-config", ".", "report type")
	flag.IntVar(&workers, "workers", backtest.DefaultBacktestWorkers, "number of concurrent workers")
	flag.IntVar(&lastDays, "last", backtest.DefaultLastDays, "number of days to do backtest")
	flag.BoolVar(&addSplits, "splits", false, "add the split strategies")
	flag.BoolVar(&addAnds, "ands", false, "add the and strategies")
	flag.Parse()

	source, err := asset.NewRepository(repositoryName, repositoryConfig)
	if err != nil {
		log.Fatalf("unable to initialize source: %v", err)
	}

	report, err := backtest.NewReport(repositoryName, repositoryConfig)
	if err != nil {
		log.Fatalf("unable to initialize report: %v", err)
	}

	backtester := backtest.NewBacktest(source, report)
	backtester.Workers = workers
	backtester.LastDays = lastDays
	backtester.Names = append(backtester.Names, flag.Args()...)
	backtester.Strategies = append(backtester.Strategies, compound.AllStrategies()...)
	backtester.Strategies = append(backtester.Strategies, momentum.AllStrategies()...)
	backtester.Strategies = append(backtester.Strategies, strategy.AllStrategies()...)
	backtester.Strategies = append(backtester.Strategies, trend.AllStrategies()...)
	backtester.Strategies = append(backtester.Strategies, volatility.AllStrategies()...)

	if addSplits {
		backtester.Strategies = append(backtester.Strategies, strategy.AllSplitStrategies(backtester.Strategies)...)
	}

	if addAnds {
		backtester.Strategies = append(backtester.Strategies, strategy.AllAndStrategies(backtester.Strategies)...)
	}

	err = backtester.Run()
	if err != nil {
		log.Fatalf("unable to run backtest: %v", err)
	}
}
