package com.example.genlib.bean;


import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class Request {
	
	@Expose(serialize = false)
	private String name; 
	
	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}
	
}
