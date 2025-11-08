package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("🔨 Building static site...")
	if err := buildSite(); err != nil {
		log.Fatalf("Build failed: %v", err)
	}
	fmt.Println("✨ Build complete!")

	fmt.Println("\n🚀 Deploying to Cloudflare Pages...")
	if err := deploy(); err != nil {
		log.Fatalf("Deployment failed: %v", err)
	}
	fmt.Println("✅ Deployment complete!")
}

func buildSite() error {
	cmd := exec.Command("go", "run", "cmd/build/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func deploy() error {
	if _, err := exec.LookPath("wrangler"); err != nil {
		return fmt.Errorf("wrangler not found. Install it with: npm install -g wrangler")
	}

	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	distPath := filepath.Join(projectRoot, "dist")

	cmd := exec.Command("wrangler", "pages", "deploy", distPath, "--project-name=aacevski-com", "--branch=production")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
