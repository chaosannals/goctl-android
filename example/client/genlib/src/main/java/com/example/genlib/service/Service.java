package com.example.genlib.service;

import com.example.genlib.bean.*;
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

	
    public void demoHandler(Request req, Callback<Response> callback, RequestBody body) {
        Call<Response> call = service.demoHandler(req.getName(), body);
        call.enqueue(callback);
    }
    
    public void demoHandler(Request req, Callback<Response> callback) {
        Call<Response> call = service.demoHandler(req.getName());
        call.enqueue(callback);
    }
    
	
}
