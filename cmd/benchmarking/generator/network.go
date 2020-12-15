package generator

import (
	"os/exec"
)

func Clean() {
	sifnodedArgs := []string{"rf", "~/.sifnoded"}
	sifnodecliArgs := []string{"rf", "~/.sifnodecli"}

	exec.Command("rm", sifnodedArgs...)
	exec.Command("rm", sifnodecliArgs...)
}

func InitialiseConfigurationFiles(chainIdName string) {
	sifnodedArgs := []string{"init", "test", "--chain-id=" + chainIdName}
	exec.Command("sifnoded", sifnodedArgs...)
}

func ConfigureSifnodeCli(output string, indent string, trustNode string, chainId string, keyringBackend string) {
	outputArgs := []string{"config", "output", output}
	indentArgs := []string{"config", "indent", indent}
	trustNodeArgs := []string{"config", "trust-node", trustNode}
	chainIdArgs := []string{"config", "chain-id", chainId}
	keyringBackendArgs := []string{"config", "keyring-backend", keyringBackend}

	exec.Command("sifnodecli", outputArgs...)
	exec.Command("sifnodecli", indentArgs...)
	exec.Command("sifnodecli", trustNodeArgs...)
	exec.Command("sifnodecli", chainIdArgs...)
	exec.Command("sifnodecli", keyringBackendArgs...)
}
