package scripts

import _ "embed"

//go:embed blockscout_trigger.sql
var BlockscoutTriggerSQL string

//go:embed blockscout_balances_trigger.sql
var BlockscoutBalancesTriggerSQL string
