package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

const TERRAFORM_TEMPLATES_PATH = "/terraform"

func RunInstaller() {
	CreateDirectories()
	InstallTerraform()
	CreateTemplates()
}

func CreateDirectories() {
	homeDirectory := filepath.Join(os.Getenv("HOME"), ".workspaces/terraform")

	if err := os.MkdirAll(homeDirectory, os.ModePerm); err != nil {
		log.Fatalf("error creating Terraform directory: %s", err)
	}
}

func InstallTerraform() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	_, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}
}

func CreateTemplates() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting work path: %s", err)
	}
	coderVscodeTemplate := filepath.Join(currentPath, TERRAFORM_TEMPLATES_PATH)
	coderVsCodeDir := filepath.Join(os.Getenv("HOME"), ".workspaces/terraform/coder-vscode")

	if err := os.MkdirAll(coderVsCodeDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating Coder vscode directory: %s", err)
	}

	if err := CopyFile(coderVscodeTemplate+"/coder-vscode.tf", coderVsCodeDir+"/coder-vscode.tf"); err != nil {
		log.Fatalf("Error creating Coder vscode template: %s", err)
	}
}

func RunTemplate(workingDir string) (string, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		return "", fmt.Errorf("error installing Terraform: %s", err)
	}

	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return "", fmt.Errorf("error running NewTerraform: %s", err)
	}

	if err := tf.Init(context.Background(), tfexec.Upgrade(true)); err != nil {
		return "", fmt.Errorf("error running Init: %s", err)
	}

	planned, err := tf.Plan(context.Background())
	if err != nil {
		return "", fmt.Errorf("error running Plan: %s", err)
	}
	if !planned {
		return "", fmt.Errorf("error running Plan, returned false: %s", err)
	}

	if err := tf.Apply(context.Background()); err != nil {
		return "", fmt.Errorf("error running Apply: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		return "", fmt.Errorf("error running Show: %s", err)
	}

	return state.FormatVersion, nil
}
