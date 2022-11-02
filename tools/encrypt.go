package tools

// This file is part of rabbit-backup will get a file and encrypt it using GPG
//

import (
	"os"
	"os/exec"
)

// GPG is a struct to hold GPG commands
type GPG struct {
	// GPGPath is the path to GPG executable
	GPGPath string
}

// NewGPG creates a new GPG struct
func NewGPG() *GPG {
	return &GPG{
		GPGPath: "/usr/bin/gpg",
	}
}

// Encrypt encrypts a file using GPG
func (g *GPG) Encrypt(file, recipient string) error {
	// gpg --encrypt --recipient config.gpg.recipient config.backup.file
	cmd := exec.Command(g.GPGPath, "--encrypt", "--recipient", recipient, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Decrypt decrypts a file using GPG
func (g *GPG) Decrypt(file, recipient string) error {
	// gpg --decrypt --recipient config.gpg.recipient config.backup.file
	cmd := exec.Command(g.GPGPath, "--decrypt", "--recipient", recipient, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
