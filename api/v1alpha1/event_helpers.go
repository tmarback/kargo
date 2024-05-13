package v1alpha1

import (
	"encoding/json"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SetEventPayloads encodes given payloads in JSON and set it to the object annotation.
func SetEventPayloads(obj client.Object, payloads ...EventPayload) error {
	data, err := json.Marshal(payloads)
	if err != nil {
		return fmt.Errorf("marshal event payloads: %w", err)
	}
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string, 1)
	}
	annotations[AnnotationKeyEventPayloads] = string(data)
	obj.SetAnnotations(annotations)
	return nil
}

// ClearEventPayloads clears event payloads annotation from the given object.
func ClearEventPayloads(obj client.Object) {
	annotations := obj.GetAnnotations()
	if annotations != nil {
		delete(annotations, AnnotationKeyEventPayloads)
	}
}

// ExtractEventPayloads extracts event payloads from the given object.
// Returns event payloads and bool value, which indicates whether the object
// has payloads or not.
func ExtractEventPayloads(obj client.Object) ([]EventPayload, bool, error) {
	raw, ok := obj.GetAnnotations()[AnnotationKeyEventPayloads]
	if !ok {
		return nil, false, nil
	}
	var payloads []EventPayload
	if err := json.Unmarshal([]byte(raw), &payloads); err != nil {
		return nil, true, fmt.Errorf("unmarshal event payloads: %w", err)
	}
	return payloads, true, nil
}
