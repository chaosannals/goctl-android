package {{.ParentPackage}}.bean;

{{.Import}}
import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class {{.Name.ToCamel}} {
	{{range $index,$item :=  .Members}}{{$item.Doc}}
	{{if eq $item.Tag "json"}}@SerializedName("{{$item.Field}}"){{else}}@Expose(serialize = false){{end}}
	private {{$item.TypeName}} {{$item.Name.Untitle}}; {{$item.Comment}}
	{{end}}{{range $index,$item :=  .Members}}
	public {{$item.TypeName}} get{{$item.Name.ToCamel}}() {
		return {{$item.Name.Untitle}};
	}

	public void set{{$item.Name.ToCamel}}({{$item.TypeName}} {{$item.Name.Untitle}}) {
		this.{{$item.Name.Untitle}} = {{$item.Name.Untitle}};
	}
	{{end}}
}
