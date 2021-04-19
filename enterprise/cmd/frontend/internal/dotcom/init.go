package dotcom

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/enterprise"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/envvar"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/dotcom/billing"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/dotcom/productsubscription"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/dotcom/stripeutil"
	"github.com/sourcegraph/sourcegraph/internal/database/dbutil"
	"github.com/sourcegraph/sourcegraph/internal/oobmigration"
)

// dotcomRootResolver implements the GraphQL types DotcomMutation and DotcomQuery.
type dotcomRootResolver struct {
	productsubscription.ProductSubscriptionLicensingResolver
	billing.BillingResolver
}

func (d dotcomRootResolver) Dotcom() graphqlbackend.DotcomResolver {
	return d
}

var _ graphqlbackend.DotcomRootResolver = dotcomRootResolver{}

func Init(ctx context.Context, db dbutil.DB, outOfBandMigrationRunner *oobmigration.Runner, enterpriseServices *enterprise.Services) error {
	stripeEnabled := stripeutil.ValidateAndPublishConfig()
	// Only enabled on Sourcegraph.com or when Stripe is configured correctly.
	if envvar.SourcegraphDotComMode() || stripeEnabled {
		enterpriseServices.DotcomResolver = dotcomRootResolver{
			ProductSubscriptionLicensingResolver: productsubscription.ProductSubscriptionLicensingResolver{
				DB: db,
			},
			BillingResolver: billing.BillingResolver{DB: db},
		}
	}
	return nil
}
