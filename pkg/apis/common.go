package apis

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/utils"
)

type Params interface {
	SetParams(...interface{}) error
}

// for ctl
func SetParams(args []interface{}, ids ...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) != len(ids) {
		return fmt.Errorf("SetParams: args need %d items", len(ids))
	}
	for i := range ids {
		switch v := ids[i].(type) {
		case *int64:
			val, err := utils.ToInt64(args[i])
			if err != nil {
				return fmt.Errorf("failed to cast to int64 `%s`: %w", args[i], err)
			}
			*v = val
		case *string:
			val, err := utils.ToString(args[i])
			if err != nil {
				return fmt.Errorf("failed to cast to string `%s`: %w", args[i], err)
			}
			*v = val
		}
	}
	return nil
}

type Spec interface {
	api.Spec
	Params
}

type ListSpec interface {
	api.ListSpec
	Spec
}

type CountableListSpec interface {
	api.CountableListSpec
	Spec
}
