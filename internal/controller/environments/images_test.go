package environments

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/akuityio/kargo/api/v1alpha1"
	"github.com/akuityio/kargo/internal/credentials"
	"github.com/akuityio/kargo/internal/images"
)

func TestGetLatestImages(t *testing.T) {
	testCases := []struct {
		name           string
		credentialsDB  credentials.Database
		getLatestTagFn func(
			context.Context,
			string,
			images.ImageUpdateStrategy,
			string,
			string,
			[]string,
			string,
			*images.Credentials,
		) (string, error)
		assertions func([]api.Image, error)
	}{
		{
			name: "error getting latest version of an image",
			credentialsDB: &credentials.FakeDB{
				GetFn: func(
					context.Context,
					string,
					credentials.Type,
					string,
				) (credentials.Credentials, bool, error) {
					return credentials.Credentials{}, false, nil
				},
			},
			getLatestTagFn: func(
				ctx context.Context,
				repoURL string,
				updateStrategy images.ImageUpdateStrategy,
				semverConstraint string,
				allowTags string,
				ignoreTags []string,
				platform string,
				creds *images.Credentials,
			) (string, error) {
				return "", errors.New("something went wrong")
			},
			assertions: func(_ []api.Image, err error) {
				require.Error(t, err)
				require.Contains(
					t,
					err.Error(),
					"error getting latest suitable tag for image",
				)
				require.Contains(t, err.Error(), "something went wrong")
			},
		},

		{
			name: "success",
			credentialsDB: &credentials.FakeDB{
				GetFn: func(
					context.Context,
					string,
					credentials.Type,
					string,
				) (credentials.Credentials, bool, error) {
					return credentials.Credentials{}, false, nil
				},
			},
			getLatestTagFn: func(
				ctx context.Context,
				repoURL string,
				updateStrategy images.ImageUpdateStrategy,
				semverConstraint string,
				allowTags string,
				ignoreTags []string,
				platform string,
				creds *images.Credentials,
			) (string, error) {
				return "fake-tag", nil
			},
			assertions: func(images []api.Image, err error) {
				require.NoError(t, err)
				require.Len(t, images, 1)
				require.Equal(
					t,
					api.Image{
						RepoURL: "fake-url",
						Tag:     "fake-tag",
					},
					images[0],
				)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testSubs := []api.ImageSubscription{
				{
					RepoURL: "fake-url",
				},
			}
			reconciler := reconciler{
				credentialsDB:  testCase.credentialsDB,
				getLatestTagFn: testCase.getLatestTagFn,
			}
			testCase.assertions(
				reconciler.getLatestImages(
					context.Background(),
					"fake-namespace",
					testSubs,
				),
			)
		})
	}
}