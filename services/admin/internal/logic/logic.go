package logic

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserLogic, NewRoleLogic, NewRuleLogic)
