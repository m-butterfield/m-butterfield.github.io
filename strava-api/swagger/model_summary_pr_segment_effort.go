/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type SummaryPrSegmentEffort struct {
	// The unique identifier of the activity related to the PR effort.
	PrActivityId int64 `json:"pr_activity_id,omitempty"`
	// The elapsed time ot the PR effort.
	PrElapsedTime int32 `json:"pr_elapsed_time,omitempty"`
	// The time at which the PR effort was started.
	PrDate time.Time `json:"pr_date,omitempty"`
	// Number of efforts by the authenticated athlete on this segment.
	EffortCount int32 `json:"effort_count,omitempty"`
}
