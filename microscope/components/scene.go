package components

import (
	"github.com/jeminstall/collab/common/jsutil"
	"github.com/jeminstall/collab/microscope/models"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jeminstall/gopherjs-vue"
)

const sceneTemplate = `
<div class="card">
	<div class="card-header" role="tab" :id="'sceneHeading-'+_uid">
		<h4 class="card-title">
			<a data-toggle="collapse" :href="'#sceneCollapse'+_uid" aria-expanded="true" :aria-controls="'sceneCollapse'+_uid">
			{{scene.question}}
			</a>
		</h4>
	</div>
	<div :id="'sceneCollapse'+_uid" class="collapse hide" role="tabpanel" :aria-labelledby="'sceneHeading-'+_uid" data-parent="#accordion">
		<div class="card-body">
			<b>Setting: </b>{{scene.setting}}<br />
			<b>Answer: </b>{{scene.answer}}<br />
		</div>
	</div>
</div>
`

type Scene struct {
	*js.Object
	Scene *models.Scene `js:"scene"`
}

func RegisterScene() {
	creator := func() interface{} {
		return &Scene{Object: jsutil.NewJSObject()}
	}
	vue.NewComponentWithProps(creator, sceneTemplate, "scene").Register("ms-scene")
}
