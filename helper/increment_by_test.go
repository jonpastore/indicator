// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/jonpastore/indicator

package helper_test

import (
	"testing"

	"github.com/jonpastore/indicator/v2/helper"
)

func TestIncrementBy(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := helper.SliceToChan([]int{2, 3, 4, 5})

	actual := helper.IncrementBy(helper.SliceToChan(input), 1)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
