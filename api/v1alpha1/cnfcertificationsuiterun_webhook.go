/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var logger = logf.Log.WithName("cnfcertificationsuiterun-resource")

var c client.Client

var (
	configMapLoggerKey       = "configMapName"
	preflightSecretLoggerKey = "preflightSecretName"
	logLevelLoggerKey        = "logLevel"
	namespaceLoggerKey       = "ns"
	cnfCertSuiteRunLoggerKey = "cnfCertificationSuiteRun"
)

func (r *CnfCertificationSuiteRun) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_, err := createClient()
	if err != nil {
		return err
	}
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

func createClient() (client.Client, error) {
	kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting OpenShift config: %v", err)
	}

	c, err = client.New(kubeconfig, client.Options{})
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	return c, nil
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//nolint:lll
//+kubebuilder:webhook:path=/validate-cnf-certifications-redhat-com-v1alpha1-cnfcertificationsuiterun,mutating=false,failurePolicy=fail,sideEffects=None,groups=cnf-certifications.redhat.com,resources=cnfcertificationsuiteruns,verbs=create;update,versions=v1alpha1,name=vcnfcertificationsuiterun.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CnfCertificationSuiteRun{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CnfCertificationSuiteRun) ValidateCreate() error {
	logger.Info("validate create", "name", r.Name)

	err := r.validateConfigMap()
	if err != nil {
		return err
	}

	err = r.validatePreflightSecret()
	if err != nil {
		return err
	}

	err = r.validateLogLevel()
	if err != nil {
		return err
	}

	return nil
}

func (r *CnfCertificationSuiteRun) validateConfigMap() error {
	configMap := &v1.ConfigMap{}

	if r.Spec.ConfigMapName == "" {
		err := fmt.Errorf("spec.configMapName must not be an empty string")
		logger.Error(err, "CnfCertificationSuiteRun's config map name is invalid",
			configMapLoggerKey, r.Spec.ConfigMapName, namespaceLoggerKey, r.Namespace)
		return err
	}

	// Return an error if config map is not found by name and ns, or field is empty
	err := c.Get(context.TODO(), types.NamespacedName{Name: r.Spec.ConfigMapName, Namespace: r.Namespace}, configMap)
	if err != nil {
		logger.Error(err, "CnfCertificationSuiteRun's config map name field is invalid",
			configMapLoggerKey, r.Spec.ConfigMapName, namespaceLoggerKey, r.Namespace)
		return err
	}

	// Verify required field exists and that it's not empty
	if value, exists := configMap.Data["tnf_config.yaml"]; !exists || value == "" {
		err := fmt.Errorf("config map's 'tnf_config.yaml' field must be set with a non-empty and valid configuration yaml for the CNF Certification Suite")
		logger.Error(err, "CnfCertificationSuiteRun's config map is invalid",
			configMapLoggerKey, r.Spec.ConfigMapName, namespaceLoggerKey, r.Namespace)
		return err
	}

	logger.Info("CnfCertificationSuiteRun's config map field is valid", configMapLoggerKey, configMap.Name, namespaceLoggerKey, r.Namespace)
	return nil
}

func (r *CnfCertificationSuiteRun) validatePreflightSecret() error {
	preflightSecret := &v1.Secret{}

	// Nil Preflight Secret is valid
	if r.Spec.PreflightSecretName == nil {
		logger.Info("Warning: No preflight secret was set.", cnfCertSuiteRunLoggerKey, r.Name, namespaceLoggerKey, r.Namespace)
		return nil
	}

	// Return an error if preflight secret is not found by name and ns, or field is empty
	err := c.Get(context.TODO(), types.NamespacedName{Name: *r.Spec.PreflightSecretName, Namespace: r.Namespace}, preflightSecret)
	if err != nil {
		logger.Error(err, "CnfCertificationSuiteRun's preflight secret name field is invalid",
			preflightSecretLoggerKey, r.Spec.PreflightSecretName, namespaceLoggerKey, r.Namespace)
		return err
	}

	// Verify required field exists and that it's not empty
	if value, exists := preflightSecret.Data["preflight_dockerconfig.json"]; !exists || len(value) == 0 {
		err := fmt.Errorf("preflight secret's 'preflight_dockerconfig.json' field must be set with a valid docker config json content")
		logger.Error(err, "CnfCertificationSuiteRun's preflight secret is invalid",
			configMapLoggerKey, r.Spec.ConfigMapName, namespaceLoggerKey, r.Namespace)
		return err
	}

	logger.Info("CnfCertificationSuiteRun's preflight secret field is valid", preflightSecretLoggerKey, preflightSecret.Name, namespaceLoggerKey, r.Namespace)
	return nil
}

func (r *CnfCertificationSuiteRun) validateLogLevel() error {
	logLevelLowerCase := strings.ToLower(r.Spec.LogLevel)
	switch logLevelLowerCase {
	case "info", "debug", "warn", "warning", "error":
		logger.Info("CnfCertificationSuiteRun's log level field is valid", logLevelLoggerKey, logLevelLowerCase)
	default:
		err := fmt.Errorf("not a valid slog Level: %q", logLevelLowerCase)
		logger.Error(err, "CnfCertificationSuiteRun's log level field is invalid",
			logLevelLoggerKey, logLevelLowerCase)
		return err
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
//
//nolint:revive
func (r *CnfCertificationSuiteRun) ValidateUpdate(old runtime.Object) error {
	logger.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CnfCertificationSuiteRun) ValidateDelete() error {
	logger.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
