package gqldeduplicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInflate(t *testing.T) {
	tests := []struct {
		Name     string
		Expected []byte
		Given    []byte
	}{
		{
			Name: "should inflate 1st child",
			Given: []byte(`
			{
				"root": [
					{
						"__typename": "Parent",
						"id": "1",
						"name": "parent 1",
						"child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1"
						},
						"another_child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1",
							"field_2": "field 2"
						}
					},
					{
						"__typename": "Parent",
						"id": "2",
						"name": "parent 2",
						"child": {
							"__typename": "Child",
							"id": "1"
						},
						"another_child": {
							"__typename": "Child",
							"id": "2",
							"field_1": "field 1",
							"field_2": "field 2"
						}
					}
				]
			}`),
			Expected: []byte(`
			{
				"root": [
					{
						"__typename": "Parent",
						"id": "1",
						"name": "parent 1",
						"child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1"
						},
						"another_child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1",
							"field_2": "field 2"
						}
					},
					{
						"__typename": "Parent",
						"id": "2",
						"name": "parent 2",
						"child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1"
						},
						"another_child": {
							"__typename": "Child",
							"id": "2",
							"field_1": "field 1",
							"field_2": "field 2"
						}
					}
				]
			}`),
		},
		{
			Name: "should inflate nth child",
			Given: []byte(`
			{
				"root": [
					{
						"__typename": "Parent",
						"id": "1",
						"name": "parent 1",
						"child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1",
							"another_child": {
								"__typename": "AnotherChild",
								"id": "1",
								"field_1": "field 1",
								"field_2": "field 2"
							}
						}
					},
					{
						"__typename": "Parent",
						"id": "2",
						"name": "parent 2",
						"child": {
							"__typename": "Child",
							"id": "2",
							"field_1": "field 1",
							"another_child": {
								"__typename": "AnotherChild",
								"id": "1"
							}
						}
					}
				]
			}`),
			Expected: []byte(`
			{
				"root": [
					{
						"__typename": "Parent",
						"id": "1",
						"name": "parent 1",
						"child": {
							"__typename": "Child",
							"id": "1",
							"field_1": "field 1",
							"another_child": {
								"__typename": "AnotherChild",
								"id": "1",
								"field_1": "field 1",
								"field_2": "field 2"
							}
						}
					},
					{
						"__typename": "Parent",
						"id": "2",
						"name": "parent 2",
						"child": {
							"__typename": "Child",
							"id": "2",
							"field_1": "field 1",
							"another_child": {
								"__typename": "AnotherChild",
								"id": "1",
								"field_1": "field 1",
								"field_2": "field 2"
							}
						}
					}
				]
			}`),
		},
		{
			Name: "should not inflate single result",
			Given: []byte(`
			{
				"root": true
			}`),
			Expected: []byte(`
			{
				"root": true
			}`),
		},
		{
			Name:     "should not inflate null result",
			Given:    []byte(`null`),
			Expected: []byte(`null`),
		},
		{
			Name: "should not inflate object without typename and id",
			Given: []byte(`
			{
				"id": "1",
				"foo": "bar"
			}`),
			Expected: []byte(`
			{
				"id": "1",
				"foo": "bar"
			}`),
		},
		{
			Name: "should not inflate first object",
			Given: []byte(`
			{
				"root": [
					{
						"__typename": "foo",
						"id": 1,
						"name": "foo"
					},
					{
						"__typename": "foo",
						"id": 1
					}
				]
			}`),
			Expected: []byte(`
			{
				"root": [
					{
						"__typename": "foo",
						"id": 1,
						"name": "foo"
					},
					{
						"__typename": "foo",
						"id": 1,
						"name": "foo"
					}
				]
			}`),
		},
		{
			Name: "should not inflate array of string",
			Given: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"names": [
						"foo",
						"bar"
					]
				}
			}`),
			Expected: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"names": [
						"foo",
						"bar"
					]
				}
			}`),
		},
		{
			Name: "should not inflate array of number",
			Given: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"numbers": [
						1,
						2
					]
				}
			}`),
			Expected: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"numbers": [
						1,
						2
					]
				}
			}`),
		},
		{
			Name: "should not inflate array of boolean",
			Given: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"values": [
						true,
						false
					]
				}
			}`),
			Expected: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"values": [
						true,
						false
					]
				}
			}`),
		},
		{
			Name: "should not inflate nested array",
			Given: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"values": [
						[
							true,
							false
						],
						[
							true,
							false
						]
					]
				}
			}`),
			Expected: []byte(`
			{
				"root": {
					"__typename": "foo",
					"id": 1,
					"values": [
						[
							true,
							false
						],
						[
							true,
							false
						]
					]
				}
			}`),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result, err := Inflate(test.Given)
			assert.NoError(t, err)
			assert.JSONEq(t, string(test.Expected), string(result))
		})
	}

	t.Run("should return error on invalid json", func(t *testing.T) {
		result, err := Inflate([]byte(`{`))
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
