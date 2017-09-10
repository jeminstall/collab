package components

import (
	"github.com/jeminstall/collab/common/jsutil"
	"github.com/jeminstall/collab/microscope/models"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jeminstall/gopherjs-vue"
)

const periodTemplate = `
<div class="card">
	<div class="card-body">
		<h4 class="card-title">{{period.title}}</h4>
		<div id="accordion" role="tablist">
			<div v-for="event in period.events">
				<ms-event v-bind:event="event"></ms-event>
			</div>
		</div>
	</div>
</div>
`

type Period struct {
	*js.Object
	Period   *models.Period `js:"period"`
	ParentID string         `js:"parentId"`
}

func RegisterPeriod() {
	creator := func() interface{} {
		return &Period{Object: jsutil.NewJSObject()}
	}
	vue.NewComponentWithProps(creator, periodTemplate, "period").Register("ms-period")
}
