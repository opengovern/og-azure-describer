# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The resource name.</td></tr>
	<tr><td>id</td><td>The resource identifier.</td></tr>
	<tr><td>job_id</td><td>A GUID uniquely identifying the streaming job.</td></tr>
	<tr><td>job_state</td><td>Describes the state of the streaming job.</td></tr>
	<tr><td>provisioning_state</td><td>Describes the provisioning status of the streaming job.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>compatibility_level</td><td>Controls certain runtime behaviors of the streaming job.</td></tr>
	<tr><td>created_date</td><td>Specifies the time when the stream analytics job was created.</td></tr>
	<tr><td>data_locale</td><td>The data locale of the stream analytics job.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>events_late_arrival_max_delay_in_seconds</td><td>The maximum tolerable delay in seconds where events arriving late could be included.</td></tr>
	<tr><td>events_out_of_order_max_delay_in_seconds</td><td>The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.</td></tr>
	<tr><td>events_out_of_order_policy</td><td>Indicates the policy to apply to events that arrive out of order in the input event stream.</td></tr>
	<tr><td>last_output_event_time</td><td>Indicating the last output event time of the streaming job or null indicating that output has not yet been produced.</td></tr>
	<tr><td>output_error_policy</td><td>Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).</td></tr>
	<tr><td>output_start_mode</td><td>This property should only be utilized when it is desired that the job be started immediately upon creation.</td></tr>
	<tr><td>output_start_time</td><td>Indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started.</td></tr>
	<tr><td>sku_name</td><td>Describes the sku name of the streaming job.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the streaming job.</td></tr>
	<tr><td>functions</td><td>A list of one or more functions for the streaming job.</td></tr>
	<tr><td>inputs</td><td>A list of one or more inputs to the streaming job.</td></tr>
	<tr><td>outputs</td><td>A list of one or more outputs for the streaming job.</td></tr>
	<tr><td>transformation</td><td>Indicates the query and the number of streaming units to use for the streaming job.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The Azure region/location in which the resource is located.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>