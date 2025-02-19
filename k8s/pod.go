package k8s

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type Pod struct {
	Name   string
}

func GetPods(context Context, namespace Namespace) ([]Pod, error) {
	cmd := exec.Command("kubectl", "get", "pod", "--context", context.Name, "--namespace", namespace.Name);

	var stdout,stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil,errors.New(stderr.String())
	}

	lines := strings.Split(strings.Trim(stdout.String(), "\n"), "\n")

	var pods[]Pod
	for _, line := range lines {
		var name = strings.Split(line, " ")[0]
		if name != "NAME" {
			pods = append(pods, Pod{Name: name})
		}
	}

	return pods,nil
}
