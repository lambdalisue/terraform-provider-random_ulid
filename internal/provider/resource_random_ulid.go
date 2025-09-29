package provider

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oklog/ulid/v2"
)

func resourceRandomULID() *schema.Resource {
	return &schema.Resource{
		Description: "The resource `random_ulid` generates a random ULID (Universally Unique Lexicographically Sortable Identifier).",

		CreateContext: resourceRandomULIDCreate,
		ReadContext:   resourceRandomULIDRead,
		DeleteContext: resourceRandomULIDDelete,

		Schema: map[string]*schema.Schema{
			"keepers": {
				Description: "Arbitrary map of values that, when changed, will trigger recreation of resource. " +
					"This can be used to recreate a ULID when certain external dependencies change.",
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"prefix": {
				Description: "Arbitrary string to prefix the ULID with.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"id": {
				Description: "The generated ULID string.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"timestamp": {
				Description: "The timestamp component of the ULID in milliseconds since Unix epoch.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"randomness": {
				Description: "The randomness component of the ULID as a base32 string.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				id := d.Id()

				// Extract ULID without prefix if it exists
				var ulidStr string
				var prefix string
				if len(id) > 26 {
					// Potentially has a prefix
					ulidStr = id[len(id)-26:]
					prefix = id[:len(id)-26]
				} else {
					ulidStr = id
				}

				// Parse the ULID
				parsedULID, err := ulid.Parse(ulidStr)
				if err != nil {
					return nil, fmt.Errorf("failed to parse ULID: %w", err)
				}

				d.SetId(id)
				d.Set("prefix", prefix)
				d.Set("timestamp", int(parsedULID.Time()))

				// Extract randomness component
				entropy := parsedULID.Entropy()
				d.Set("randomness", fmt.Sprintf("%x", entropy))

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceRandomULIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Generate new ULID
	t := time.Now()
	entropy := rand.Reader
	newULID, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to generate ULID: %w", err))
	}

	// Handle prefix if provided
	id := newULID.String()
	if prefix, ok := d.GetOk("prefix"); ok {
		id = prefix.(string) + id
	}

	// Set resource ID and attributes
	d.SetId(id)
	d.Set("timestamp", int(newULID.Time()))

	// Extract and save randomness component
	entropy_bytes := newULID.Entropy()
	d.Set("randomness", fmt.Sprintf("%x", entropy_bytes))

	return diags
}

func resourceRandomULIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Resource is read-only, maintain existing state
	return nil
}

func resourceRandomULIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Simply clear the resource ID on deletion
	d.SetId("")
	return nil
}
