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

func EncryptBackup(file, recipient string) error {
	g := NewGPG()
	// gpg --encrypt --recipient config.gpg.recipient config.backup.file
	cmd := exec.Command(g.GPGPath, "--encrypt", "--recipient", recipient, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Decrypt decrypts a file using GPG
func DecryptBackup(file, recipient string) error {
	g := NewGPG()
	// gpg --decrypt --recipient config.gpg.recipient config.backup.file
	cmd := exec.Command(g.GPGPath, "--decrypt", "--recipient", recipient, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
