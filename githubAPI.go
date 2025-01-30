package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

type Config struct {
	DownloadURL string `json:"downloadURL"`
}

var (
	currentDirectory, _ = os.Getwd()
	exePath = filepath.Join(currentDirectory, "VRC-Avatar-Library.exe")
)

func DownloadAndReplace(a *App) error {
	time.Sleep(2 * time.Second)

	// 既存のアプリを削除
	fmt.Println("既存のアプリを削除しています...")
	err := os.Remove(exePath)
	if err != nil {
		fmt.Println("既存のアプリの削除に失敗しました:", err)
		return err
	}

	time.Sleep(1 * time.Second)

	// アプリのダウンロード
	fmt.Println("ダウンロードを開始します...")
	downloadURL := readJSON()
	resp, err := http.Get(downloadURL)
	if err != nil {
		fmt.Println("ダウンロードに失敗しました:", err)
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(exePath)
	if err != nil {
		fmt.Println("ファイルの作成に失敗しました:", err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("ファイルの書き込みに失敗しました:", err)
		return err
	}

	fmt.Println("ダウンロードが完了しました。")
	time.Sleep(1 * time.Second)
	return nil
}

func StartApp(a *App) error {
	cmd := exec.Command(exePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: false}
	err := cmd.Start()
	if err != nil {
		fmt.Println("アプリの起動に失敗しました:", err)
		return err
	}
	fmt.Println("アプリの起動に成功しました。")

	os.Exit(0)
	return nil
}

func readJSON() string {
	configJson := filepath.Join(currentDirectory, "Config", "config.json")

	file, err := os.Open(configJson)
	if err != nil {
		fmt.Println("設定ファイルの読み込みに失敗しました:", err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("設定ファイルのパースに失敗しました:", err)
	}
	return config.DownloadURL
}