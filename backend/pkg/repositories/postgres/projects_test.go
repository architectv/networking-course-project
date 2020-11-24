package postgres

// func TestProjectPg_Create(t *testing.T) {
// 	db, mock, err := sqlmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	r := NewUserPg(db)

// 	type mockBehavior func(input *models.Project)

// 	tests := []struct {
// 		name    string
// 		mock    mockBehavior
// 		input   *models.Project
// 		want    int
// 		wantErr bool
// 	}{
// 		{
// 			name: "Ok",
// 			mock: func(input *models.Project) {
// 				mock.ExpectBegin()

// 				rows := sqlmock.NewRows([]string{"id"}).AddRow(input.Id)
// 				defPermissions := input.DefaultPermissions
// 				mock.ExpectQuery("INSERT INTO permissions").
// 					WithArgs(defPermissions.Read, defPermissions.Write, defPermissions.Admin).
// 					WillReturnRows(rows)

// 				datetimes := input.Datetimes
// 				mock.ExpectQuery("INSERT INTO permissions").
// 					WithArgs(datetimes.Created, datetimes.Updated, datetimes.Accessed).
// 					WillReturnRows(rows)

// 				mock.ExpectQuery("INSERT INTO projects").
// 					WithArgs(project.OwnerId, defPermissionId, datetimesId, project.Title, project.Description).
// 					WillReturnRows(rows)
// 			},
// 			input: &models.Project{
// 				OwnerId:            1,
// 				DefaultPermissions: &models.Permission{true, true, false},
// 				Datetimes:          &models.Datetimes{100, 100, 100},
// 				Title:              "TestProject",
// 				Description:        "Test description",
// 			},
// 			want: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()

// 			got, err := r.Create(tt.input)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.want, got)
// 			}
// 		})
// 	}
// }
