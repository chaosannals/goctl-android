package {{.ParentPackage}}.bean;
{{.Import}}
public class {{.Name.ToCamel}} {
	{{range $index,$item :=  .Members}}{{$item.Doc}}
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
