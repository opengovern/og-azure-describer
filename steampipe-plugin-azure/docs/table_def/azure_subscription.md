# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.</td></tr>
	<tr><td>subscription_id</td><td>The subscription ID.</td></tr>
	<tr><td>display_name</td><td>A friendly name that identifies a subscription.</td></tr>
	<tr><td>tenant_id</td><td>The subscription tenant ID.</td></tr>
	<tr><td>state</td><td>The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted. Possible values include: &#39;StateEnabled&#39;, &#39;StateWarned&#39;, &#39;StatePastDue&#39;, &#39;StateDisabled&#39;, &#39;StateDeleted&#39;</td></tr>
	<tr><td>authorization_source</td><td>The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management. For example, &#39;Legacy, RoleBased&#39;.</td></tr>
	<tr><td>managed_by_tenants</td><td>An array containing the tenants managing the subscription.</td></tr>
	<tr><td>subscription_policies</td><td>The subscription policies.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>kaytu_metadata</td><td>Metadata of the AWS resource</td></tr>
</table>