/*
Copyright the Velero contributors.

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

package velero

import (
	"k8s.io/apimachinery/pkg/runtime"

	api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
)

// RestoreItemActionExecuteInput contains the input parameters for the ItemAction's Execute function.
type RestoreItemActionExecuteInput struct {
	// Item is the item being restored. It is likely different from the pristine backed up version
	// (metadata reset, changed by various restore item action plugins, etc.).
	Item runtime.Unstructured
	// ItemFromBackup is the item taken from the pristine backed up version of resource.
	ItemFromBackup runtime.Unstructured
	// Restore is the representation of the restore resource processed by Velero.
	Restore *api.Restore
}

// RestoreItemActionExecuteOutput contains the output variables for the ItemAction's Execution function.
type RestoreItemActionExecuteOutput struct {
	// UpdatedItem is the item being restored mutated by ItemAction.
	UpdatedItem runtime.Unstructured

	// AdditionalItems is a list of additional related items that should
	// be restored.
	AdditionalItems []ResourceIdentifier

	// SkipRestore tells velero to stop executing further actions
	// on this item, and skip the restore step. When this field's
	// value is true, AdditionalItems will be ignored.
	SkipRestore bool
}

// NewRestoreItemActionExecuteOutput creates a new RestoreItemActionExecuteOutput
func NewRestoreItemActionExecuteOutput(item runtime.Unstructured) *RestoreItemActionExecuteOutput {
	return &RestoreItemActionExecuteOutput{
		UpdatedItem: item,
	}
}

// WithoutRestore returns SkipRestore for RestoreItemActionExecuteOutput
func (r *RestoreItemActionExecuteOutput) WithoutRestore() *RestoreItemActionExecuteOutput {
	r.SkipRestore = true
	return r
}
