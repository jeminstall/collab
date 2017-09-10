package components

import (
	"collab/common/jsutil"
	"collab/microscope/models"

	"github.com/gopherjs/gopherjs/js"
	vue "github.com/jeminstall/gopherjs-vue"
)

const eventTemplate = `
<div class="card">
	<div class="card-header" role="tab" :id="'eventHeading-'+_uid">
		<h4>
			<a data-toggle="collapse" :href="'#eventCollapse'+_uid" aria-expanded="true" :aria-controls="'eventCollapse'+_uid">
			{{event.title}}
			</a>
		</h4>
	</div>
	<div :id="'eventCollapse'+_uid" class="collapse hide" role="tabpanel" :aria-labelledby="'eventHeading-'+_uid" data-parent="#accordion">
		<div class="card-body">
			<div class="row" v-for="scene in event.scenes">
				<ms-scene v-bind:scene="scene"></ms-scene>
			</div>
		</div>
	</div>
</div>
`

type Event struct {
	*js.Object
	Event *models.Event `js:"event"`
}

func RegisterEvent() {
	creator := func() interface{} {
		return &Event{Object: jsutil.NewJSObject()}
	}
	vue.NewComponentWithProps(creator, eventTemplate, "event").Register("ms-event")
}
