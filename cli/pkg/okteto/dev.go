package okteto

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/okteto/app/cli/pkg/model"
)

// DevModeOn activates a dev environment
func DevModeOn(dev *model.Dev) error {
	c, err := getClient()
	if err != nil {
		return fmt.Errorf("error getting okteto client: %s", err)
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		mutation {
			up(name: "%s", image: "%s", workdir: "%s") {
		  		name
			}
	  	}`, dev.Name, dev.Image, dev.WorkDir))

	oktetoToken, err := getToken()
	if err != nil {
		return fmt.Errorf("please login")
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", oktetoToken))

	ctx := context.Background()

	if err := c.Run(ctx, req, nil); err != nil {
		return fmt.Errorf("error activating dev environment: %s", err)
	}

	return nil
}
