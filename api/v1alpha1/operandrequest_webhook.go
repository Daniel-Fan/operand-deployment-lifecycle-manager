//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package v1alpha1

import (
	"strings"

	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/IBM/operand-deployment-lifecycle-manager/controllers/util"
)

// log is for logging in this package.
var operandrequestlog = logf.Log.WithName("operandrequest-resource")

func (r *OperandRequest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-operator-ibm-com-v1alpha1-operandrequest,mutating=true,failurePolicy=fail,sideEffects=None,groups=operator.ibm.com,resources=operandrequests,verbs=create;update,versions=v1alpha1,name=moperandrequest.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &OperandRequest{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OperandRequest) Default() {
	operandrequestlog.Info("default", "name", r.Name)
	for i, req := range r.Spec.Requests {
		regNs := req.RegistryNamespace
		if regNs == "" {
			regNs = r.Namespace
		}
		watchNamespace := util.GetWatchNamespace()
		if !util.Contains(strings.Split(watchNamespace, ","), regNs) {
			regNs = util.GetOperatorNamespace()
		}
		r.Spec.Requests[i].RegistryNamespace = regNs
	}
	// TODO(user): fill in your defaulting logic.
}
