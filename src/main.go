package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"gopkg.in/ini.v1"
)

func get_aws_config() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return cfg
}

func get_aws_config_path() string {
	HOME, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s/.aws/credentials", HOME)
	return path
}

func load_aws_credentials() *ini.File {
	path := get_aws_config_path()
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}

func get_profile_name() string {

	DEFAULT_PROFILE_NAME := "default"

	profile_name := os.Getenv("AWS_PROFILE")
	if profile_name == "" {
		profile_name = DEFAULT_PROFILE_NAME
	}

	// Check if '${profile}-long-term' exists
	// This is used for aws-mfa
	aws_credentials := load_aws_credentials()
	long_term_name := fmt.Sprintf("%s-long-term", profile_name)

	for _, section := range aws_credentials.Sections() {
		if section.Name() == long_term_name {
			profile_name = long_term_name
		}
	}

	return profile_name
}

func save_key(key *types.AccessKey, profile_name string) error {

	// Load ~/.aws/credentials
	cfg := load_aws_credentials()
	path := get_aws_config_path()

	fmt.Printf("Saving Access Key (%s) to %s. (profile = [%s])\n",
		*key.AccessKeyId,
		path,
		profile_name)

	cfg.Section(profile_name).Key("aws_access_key_id").SetValue(*key.AccessKeyId)
	cfg.Section(profile_name).Key("aws_secret_access_key").SetValue(*key.SecretAccessKey)

	return cfg.SaveTo(path)
}

func main() {

	cfg := get_aws_config()
	client := iam.NewFromConfig(cfg)

	aws_credentials := load_aws_credentials()
	profile_name := get_profile_name()

	current_key_id := aws_credentials.Section(profile_name).Key("aws_access_key_id").String()
	// fmt.Printf("Current Access Key ID: %s\n", current_key_id)

	// Create New Access Key
	resp, err := client.CreateAccessKey(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New Access Key Created: %s\n", *resp.AccessKey.AccessKeyId)

	// Delete old access key
	fmt.Printf("Deleting old access key: %s\n", current_key_id)
	_, err2 := client.DeleteAccessKey(context.TODO(), &iam.DeleteAccessKeyInput{
		AccessKeyId: &current_key_id,
	})
	if err2 != nil {
		log.Fatal(err)
	} else {
		save_key(resp.AccessKey, profile_name)
	}
}
