/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SummaryGear struct {
	// The gear's unique identifier.
	Id string `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// Whether this gear's is the owner's default one.
	Primary bool `json:"primary,omitempty"`
	// The gear's name.
	Name string `json:"name,omitempty"`
	// The distance logged with this gear.
	Distance float32 `json:"distance,omitempty"`
}
