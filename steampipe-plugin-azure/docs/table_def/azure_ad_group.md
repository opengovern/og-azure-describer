# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>object_id</td><td>The unique ID that identifies a group.</td></tr>
	<tr><td>object_type</td><td>A string that identifies the object type.</td></tr>
	<tr><td>display_name</td><td>A friendly name that identifies a group.</td></tr>
	<tr><td>mail</td><td>The primary email address of the group.</td></tr>
	<tr><td>mail_enabled</td><td>Indicates whether the group is mail-enabled. Must be false. This is because only pure security groups can be created using the Graph API.</td></tr>
	<tr><td>mail_nickname</td><td>The mail alias for the group.</td></tr>
	<tr><td>deletion_timestamp</td><td>The time at which the directory object was deleted.</td></tr>
	<tr><td>security_enabled</td><td>Specifies whether the group is a security group.</td></tr>
	<tr><td>additional_properties</td><td>A list of unmatched properties from the message are deserialized this collection.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
</table>