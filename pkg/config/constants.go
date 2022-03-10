// Put crispr level constants here.
package config

// All standard sequence related paths.
const (
	SequenceLogsDir    = "logs"
	SequenceLogFile    = "logs.log"
	SequenceErrLogFile = "errors.log"
	SequenceOutDir     = "outputs"
	SequenceOutfile    = "output.json"
	SequenceInDir      = "inputs"
)

// AdminSequenceOutKeys is the list of keys that must be provided in the output.json by admin sequence.
var adminSequenceOutKeys = []string{"cell_automation_role", "cell_admin_role", "cell_aws_account_role", "cell_bucket", "encrypt", "cell_state_region", "cell_dynamodb_table", "cell_kms_key_id"}

func GetAdminSequenceOutKeys() []string {
	// keep adminSequenceOutKeys readonly
	return adminSequenceOutKeys
}
