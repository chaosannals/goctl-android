package {{.ParentPackage}}.service;

{{.Import}}
import com.alibaba.fastjson.JSON;
import okhttp3.MediaType;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class Service {
    private static final String MEDIA_TYPE_JSON = "application/json; charset=utf-8";
    private static final String BASE_RUL = "http://localhost:8888/";// TODO replace to your host and delete this comment
    private static Service instance;
    private static IService service;

    private Service() {
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(BASE_RUL)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
        service = retrofit.create(IService.class);
    }

    public static Service getInstance() {
        if (instance == null) {
            instance = new Service();
        }
        return instance;
    }

    private RequestBody buildJSONBody(Object obj) {
        String s = JSON.toJSONString(obj);
        return RequestBody.create(s, MediaType.parse(MEDIA_TYPE_JSON));
    }
	{{range $index,$item := .Routes}}{{$item.Doc}}
    public void {{$item.MethodName}}({{if $item.HasRequest}}{{$item.RequestBeanName}} in, {{end}}Callback{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} callback) {
        Call{{if $item.HasResponse}}<{{$item.ResponseBeanName}}>{{else}}<Void>{{end}} call = service.{{$item.MethodName}}({{if $item.HavePath}}{{$item.PathId}}{{end}}{{if $item.HaveQuery}}{{$item.QueryId}}{{end}}{{if $item.ShowRequestBody}}buildJSONBody(in){{end}});
        call.enqueue(callback);
    }
	{{end}}
}
