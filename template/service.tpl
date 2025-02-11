package {{.ParentPackage}}.service;

{{.Import}}
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class Service {
    private final IService service;

    public Service(String baseUrl) {
        Gson gson = new GsonBuilder()
                .excludeFieldsWithoutExposeAnnotation()
                .setPrettyPrinting()
                .create();
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(baseUrl)
                .addConverterFactory(GsonConverterFactory.create(gson))
                .build();
        service = retrofit.create(IService.class);
    }

	{{range $index,$item := .Routes}}{{$item.Doc}}
    public void {{$item.MethodName}}({{if $item.HasRequest}}{{$item.RequestBeanName}} req, {{end}}Callback{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} callback{{if not $item.ShowRequestBody}}, RequestBody body{{end}}) {
        Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} call = service.{{$item.MethodName}}({{if $item.HaveHeaders}}{{$item.HeaderIdsExpr}}{{end}}{{if $item.HavePath}}{{$item.PathId}}{{end}}{{if $item.HaveQuery}}{{$item.QueryId}}{{end}}{{if $item.ShowRequestBody}}req{{else}}{{$item.BodyPrefix}}body{{end}});
        call.enqueue(callback);
    }
    {{if not $item.ShowRequestBody}}
    public void {{$item.MethodName}}({{if $item.HasRequest}}{{$item.RequestBeanName}} req, {{end}}Callback{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} callback) {
        Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} call = service.{{$item.MethodName}}({{if $item.HaveHeaders}}{{$item.HeaderIdsExpr}}{{end}}{{if $item.HavePath}}{{$item.PathId}}{{end}}{{if $item.HaveQuery}}{{$item.QueryId}}{{end}});
        call.enqueue(callback);
    }
    {{end}}
	{{end}}
}
