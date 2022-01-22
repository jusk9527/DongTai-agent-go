package htmlTemplateExecuteTemplate

import (
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"html/template"
	"io"
)

func ExecuteTemplate(template *template.Template, wr io.Writer, name string, data interface{}) error {
	err := ExecuteTemplateT(template, wr, name, data)
	utils.FmtHookPool(request.PoolReq{
		Args:            utils.Collect(wr, name, data),
		Reqs:            utils.Collect(err),
		Source:          false,
		OriginClassName: "template.(*Template)",
		MethodName:      "ExecuteTemplate",
		ClassName:       "template.(*Template)",
	})
	return err
}

func ExecuteTemplateT(template *template.Template, wr io.Writer, name string, data interface{}) error {
	return nil
}
