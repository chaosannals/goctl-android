package com.example.genlib.service;

import com.example.genlib.bean.*;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.http.*;

public interface IService {
    
	@GET("/from/{name}")
	Call<Response> demoHandler(@Path("name") String name, @Body RequestBody body);
	
	@GET("/from/{name}")
	Call<Response> demoHandler(@Path("name") String name);
	
	
}
