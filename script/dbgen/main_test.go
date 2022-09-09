package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_readSchema(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []Schema
		wantErr bool
	}{
		{
			args: args{"test/xo.xo.json"},
			want: []Schema{
				{
					Type: "mysql",
					Name: "takos_dev",
					Tables: []Table{
						{
							Name: "examination_questions",
							Columns: []Column{
								{
									Name: "id",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: Char,
										Prec: 36,
									},
									IsPrimary: true,
								},
								{
									Name: "examination_id",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: Char,
										Prec: 36,
									},
								},
								{
									Name: "name",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: VarChar,
										Prec: 300,
									},
								},
								{
									Name: "description",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: VarChar,
										Prec: 300,
									},
								},
								{
									Name: "allocated_score",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: SmallInt,
										Prec: 5,
									},
								},
								{
									Name: "answer_format",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: SmallInt,
										Prec: 5,
									},
								},
								{
									Name: "created_at",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: TimeStamp,
									},
								},
								{
									Name: "updated_at",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: TimeStamp,
									},
								},
							},
							Indexes: []Index{
								{
									Name: "examination_id",
									Fields: []Column{
										{
											Name: "examination_id",
											Datatype: struct {
												Type DataType
												Prec int
											}{
												Type: Char,
												Prec: 36,
											},
										},
									},
								},
								{
									Name: "examination_questions_id_pkey",
									Fields: []Column{
										{
											Name: "id",
											Datatype: struct {
												Type DataType
												Prec int
											}{
												Type: Char,
												Prec: 36,
											},
											IsPrimary: true,
										},
									},
									IsUnique:  true,
									IsPrimary: true,
								},
							},
						},
						{
							Name: "examinations",
							Columns: []Column{
								{
									Name: "id",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: Char,
										Prec: 36,
									},
									IsPrimary: true,
								},
								{
									Name: "name",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: VarChar,
										Prec: 300,
									},
								},
								{
									Name: "description",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: VarChar,
										Prec: 600,
									},
								},
								{
									Name: "created_at",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: TimeStamp,
									},
								},
								{
									Name: "updated_at",
									Datatype: struct {
										Type DataType
										Prec int
									}{
										Type: TimeStamp,
									},
								},
							},
							Indexes: []Index{
								{
									Name: "name",
									Fields: []Column{
										{
											Name: "name",
											Datatype: struct {
												Type DataType
												Prec int
											}{
												Type: VarChar,
												Prec: 300,
											},
										},
									},
									IsUnique: true,
								},
								{
									Name: "examinations_id_pkey",
									Fields: []Column{
										{
											Name: "id",
											Datatype: struct {
												Type DataType
												Prec int
											}{
												Type: Char,
												Prec: 36,
											},
											IsPrimary: true,
										},
									},
									IsUnique:  true,
									IsPrimary: true,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSchema(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got.Schemas); diff != "" {
				t.Error(diff)
			}
		})
	}
}
