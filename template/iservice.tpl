package {{.ParentPackage}}.service;

{{.Import}}
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.http.*;

public interface IService {
    {{range $index,$item := .Routes}}{{$item.Doc}}
	@{{$item.Method}}("{{$item.Path}}")
	Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} {{$item.MethodName}}({{if $item.HaveHeaders}}{{$item.HeadersExpr}}{{end}}{{if $item.HavePath}}{{$item.PathIdExpr}}{{end}}{{if $item.HaveQuery}}{{$item.QueryIdExpr}}{{end}}{{$item.BodyPrefix}}@Body {{if $item.ShowRequestBody}}{{$item.RequestBeanName}}{{else}}RequestBody{{end}} body);
	{{if not $item.ShowRequestBody}}
	@{{$item.Method}}("{{$item.Path}}")
	Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} {{$item.MethodName}}({{if $item.HaveHeaders}}{{$item.HeadersExpr}}{{end}}{{if $item.HavePath}}{{$item.PathIdExpr}}{{end}}{{if $item.HaveQuery}}{{$item.QueryIdExpr}}{{end}});
	{{end}}
	{{end}}
}
