package main

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

type newCommand struct {
	*editorCommand
}

func newCommandFromContext(c *cli.Context) (*newCommand, error) {
	ecmd, err := editorCommandFromContext(c)
	if err != nil {
		return nil, err
	}

	if exists(ecmd.filename) {
		return nil, fmt.Errorf("%s already exists", ecmd.filename)
	}

	return &newCommand{ecmd}, nil
}

func (c *newCommand) run() error {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "new-secret",
		},
		Data: map[string][]byte{"Key": []byte("Value")},
		Type: "Opaque",
	}

	scheme := runtime.NewScheme()
	serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme, scheme)

	buf := new(bytes.Buffer)

	if err := serializer.Encode(secret, buf); err != nil {
		return err
	}

	yaml := buf.Bytes()

	if c.decode {
		decoded, err := decode(yaml)
		if err != nil {
			return nil
		}
		yaml = decoded
	}

	updatedPlainText, err := c.editText(yaml)
	if err != nil {
		return err
	}

	if c.decode {
		encoded, err := encode(updatedPlainText)
		if err != nil {
			return err
		}
		updatedPlainText = encoded
	}

	if err := c.validator.validate(updatedPlainText); err != nil {
		return err
	}

	return c.create(updatedPlainText)
}
