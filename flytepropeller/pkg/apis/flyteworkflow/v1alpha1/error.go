package v1alpha1

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

type ExecutionError struct {
	*core.ExecutionError
}

func (in *ExecutionError) UnmarshalJSON(b []byte) error {
	in.ExecutionError = &core.ExecutionError{}
	return jsonpb.Unmarshal(bytes.NewReader(b), in.ExecutionError)
}

func (in *ExecutionError) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, in.ExecutionError); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (in *ExecutionError) DeepCopyInto(out *ExecutionError) {
	*out = *in
}
