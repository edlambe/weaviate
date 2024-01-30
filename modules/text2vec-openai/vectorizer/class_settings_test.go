//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package vectorizer

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/moduletools"
	"github.com/weaviate/weaviate/entities/schema"
)

func Test_classSettings_Validate(t *testing.T) {
	class := &models.Class{
		Class: "test",
		Properties: []*models.Property{
			{
				DataType: []string{schema.DataTypeText.String()},
				Name:     "test",
			},
		},
	}
	tests := []struct {
		name    string
		cfg     moduletools.ClassConfig
		wantErr error
	}{
		{
			name: "text-embedding-3-small",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model": "text-embedding-3-small",
				},
			},
		},
		{
			name: "text-embedding-3-small, 512 dimensions",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":      "text-embedding-3-small",
					"dimensions": 512,
				},
			},
		},
		{
			name: "text-embedding-3-small, wrong dimensions",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":      "text-embedding-3-small",
					"dimensions": 1,
				},
			},
			wantErr: errors.New("wrong dimensions setting for text-embedding-3-small model, available dimensions are: [512 1536]"),
		},
		{
			name: "text-embedding-3-large",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model": "text-embedding-3-large",
				},
			},
		},
		{
			name: "text-embedding-3-large, 512 dimensions",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":      "text-embedding-3-large",
					"dimensions": 1024,
				},
			},
		},
		{
			name: "text-embedding-3-large, wrong dimensions",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":      "text-embedding-3-large",
					"dimensions": 512,
				},
			},
			wantErr: errors.New("wrong dimensions setting for text-embedding-3-large model, available dimensions are: [256 1024 3072]"),
		},
		{
			name: "text-embedding-ada-002",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":        "ada",
					"modelVersion": "002",
				},
			},
		},
		{
			name: "text-embedding-ada-002 - dimensions error",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":      "ada",
					"dimensions": 512,
				},
			},
			wantErr: errors.New("dimensions setting can only be used with V3 embedding models: [text-embedding-3-small text-embedding-3-large]"),
		},
		{
			name: "text-embedding-ada-002 - wrong model version",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model":        "ada",
					"modelVersion": "003",
				},
			},
			wantErr: errors.New("unsupported version 003"),
		},
		{
			name: "wrong model name",
			cfg: &fakeClassConfig{
				classConfig: map[string]interface{}{
					"model": "unknown-model",
				},
			},
			wantErr: errors.New("wrong OpenAI model name, available model names are: [ada babbage curie davinci text-embedding-3-small text-embedding-3-large]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewClassSettings(tt.cfg)
			err := cs.Validate(class)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
