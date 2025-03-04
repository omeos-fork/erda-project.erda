// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package addon

import (
	"github.com/erda-project/erda/apistructs"
)

const (
	LabelKeyPrefix = "annotations/"
)

func Create(op AddonOperator, sg *apistructs.ServiceGroup) error {
	if err := op.Validate(sg); err != nil {
		return err
	}
	k8syml, err := op.Convert(sg)
	if err != nil {
		return err
	}

	return op.Create(k8syml)
}

func Inspect(op AddonOperator, sg *apistructs.ServiceGroup) (*apistructs.ServiceGroup, error) {
	if err := op.Validate(sg); err != nil {
		return nil, err
	}
	return op.Inspect(sg)
}

func Remove(op AddonOperator, sg *apistructs.ServiceGroup) error {
	return op.Remove(sg)
}

func Update(op AddonOperator, sg *apistructs.ServiceGroup) error {
	if err := op.Validate(sg); err != nil {
		return err
	}
	k8syml, err := op.Convert(sg)
	if err != nil {
		return err
	}

	return op.Update(k8syml)
}
