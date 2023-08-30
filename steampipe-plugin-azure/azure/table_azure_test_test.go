package azure

func tableAzuretesttest(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_test_test",
		Description: "Azure test test",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.Gettesttest,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.Listtesttest,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the test.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.test.Id")},
			{
				Name:        "name",
				Description: "The name of the test.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.test.Name")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.test.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.test.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.test.ID").Transform(idToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func gettesttestTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.testtest).Description.test.Tags
	return ec2V2TagsToMap(tags)
}

func gettesttestArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	test := d.HydrateItem.(kaytu.testtest).Description.test
	metadata := d.HydrateItem.(kaytu.testtest).Metadata

	arn := fmt.Sprintf("") //TODO generate the arn
	return arn, nil
}
